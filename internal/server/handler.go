package server

import (
	"compress/gzip"
	"fmt"
	"hash/fnv"
	"html/template"
	"io"
	"io/fs"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
)

// TemplateFS is the filesystem used to load templates.
// In production this is the real FS; tests may override it.
var TemplateFS fs.FS = os.DirFS("templates")

// AssetsFS is the filesystem used to serve /assets/* files.
// Override in embed builds to use an embedded FS.
var AssetsFS fs.FS = os.DirFS("assets")

// templateData is the top-level data passed to every page template.
type templateData struct {
	Config          *config.Config
	PageTitle       string
	PageDescription string
	OGType          string
	Canonical       string
	CSSVersion      string
	Menu            []blog.MenuLink
	FooterHTML      template.HTML
	Version         blog.VersionInfo
	RenderTime      string
	JSFiles         []string
	JSONLD          template.HTML // structured data JSON-LD block (injected into <head>)
	Content         template.HTML // rendered inner content
	// page-specific fields
	Posts           []blog.Post
	Post            *blog.Post
	Page            *blog.Page
	Tags            []string
	CategorySlug    string
	CurrentCategory *blog.CategoryInfo
	Pagination      blog.Pagination
	Categories      []blog.CategoryInfo
	SubCategories   []blog.CategoryInfo
	IndexBlurb      template.HTML
	Query           string
}

// postPreviewData is passed to the post_preview template.
type postPreviewData struct {
	Post    blog.Post
	PostURL string
	Config  *config.Config
}

// Options configures optional filesystem overrides for Handler.
// Nil fields default to TemplateFS and AssetsFS.
type Options struct {
	TemplateFS fs.FS
	AssetsFS   fs.FS
}

// Handler is the main HTTP handler for the blog.
type Handler struct {
	cfg         *config.Config
	b           *blog.Blog
	tmpl        *template.Template
	assetsFS    fs.FS
	mux         *http.ServeMux
	menu        []blog.MenuLink
	footerHTML  template.HTML
	cssVersion  string
	versionInfo blog.VersionInfo
	indexBlurb  template.HTML
}

// New creates and returns the blog HTTP handler using the default TemplateFS
// and AssetsFS. It panics if templates cannot be loaded.
func New(cfg *config.Config, b *blog.Blog) *Handler {
	return NewWithOptions(cfg, b, Options{})
}

// NewWithOptions creates and returns the blog HTTP handler with injected filesystems.
// It panics if templates cannot be loaded.
func NewWithOptions(cfg *config.Config, b *blog.Blog, opts Options) *Handler {
	if cfg.CSSTheme != "" && !strings.HasPrefix(cfg.CSSTheme, "/") && !strings.HasPrefix(cfg.CSSTheme, "http://") && !strings.HasPrefix(cfg.CSSTheme, "https://") {
		cfg.CSSTheme = "/" + cfg.CSSTheme
	}
	tmplFS := opts.TemplateFS
	if tmplFS == nil {
		tmplFS = TemplateFS
	}
	assetsFS := opts.AssetsFS
	if assetsFS == nil {
		assetsFS = AssetsFS
	}

	tmpl := mustLoadTemplates(tmplFS)
	versionInfo := b.GetVersionInfo()

	var indexBlurb template.HTML
	blurbFile := filepath.Join(cfg.PostsDir, "index.md")
	if raw, err := os.ReadFile(blurbFile); err == nil {
		indexBlurb = template.HTML(b.ParseMarkdown(string(raw)))
	}

	h := &Handler{
		cfg:         cfg,
		b:           b,
		tmpl:        tmpl,
		assetsFS:    assetsFS,
		menu:        b.GetMenu(),
		footerHTML:  template.HTML(b.ParseMarkdown(cfg.FooterContent)),
		cssVersion:  cssVersion(assetsFS, cfg.CSSTheme, versionInfo.Commit),
		versionInfo: versionInfo,
		indexBlurb:  indexBlurb,
	}
	h.mux = h.buildMux()
	return h
}

func (h *Handler) buildMux() *http.ServeMux {
	mux := http.NewServeMux()
	assetServer := http.StripPrefix("/assets/", http.FileServerFS(h.assetsFS))

	mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		rel := strings.TrimPrefix(r.URL.Path, "/assets/")
		if strings.Contains(rel, "..") || strings.Contains(rel, "\x00") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", h.cacheControlAssets())
		assetServer.ServeHTTP(w, r)
	})
	mux.HandleFunc("/favicon.ico", h.serveFavicon)
	mux.HandleFunc("/search/label/", func(w http.ResponseWriter, r *http.Request) {
		tag := strings.TrimPrefix(r.URL.Path, "/search/label/")
		http.Redirect(w, r, "/?q="+tag+"&search=true", http.StatusMovedPermanently)
	})
	mux.HandleFunc("/feed.xml", h.serveFeedXML)
	mux.HandleFunc("/feed", h.serveFeedPage)
	mux.HandleFunc("/sitemap.xml", h.serveSitemap)
	mux.HandleFunc("/robots.txt", h.serveRobots)
	mux.HandleFunc("/pages/", h.serveCleanPage)
	mux.HandleFunc("/content/", h.serveCleanContent)
	mux.HandleFunc("/", h.serveFallback)
	return mux
}

func (h *Handler) baseTemplateData(r *http.Request, ogType string) templateData {
	return templateData{
		Config:     h.cfg,
		OGType:     ogType,
		Canonical:  buildCanonical(r),
		CSSVersion: h.cssVersion,
		Menu:       h.menu,
		FooterHTML: h.footerHTML,
		Version:    h.versionInfo,
	}
}

// ServeHTTP routes all requests.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if h.cfg.CSP.Enabled && !strings.HasPrefix(path, "/assets/") && path != "/favicon.ico" && !strings.HasPrefix(path, "/search/label/") {
		w.Header().Set(splitCSPHeader(h.cfg.CSP.Header))
	}
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) serveFallback(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Legacy category query redirect (/?category=...)
	if path == "/" {
		q := r.URL.Query()
		if q.Has("category") {
			catSlug := q.Get("category")
			if h.b.GetCategoryBySlug(catSlug) != nil {
				dest := h.cfg.CategoryPath(catSlug)
				if q.Has("page") {
					dest += "?page=" + q.Get("page")
				}
				http.Redirect(w, r, dest, http.StatusMovedPermanently)
				return
			}
		}
	}

	// Legacy /page?slug=... redirect
	if path == "/page" || strings.HasSuffix(path, "/page") {
		if slug := r.URL.Query().Get("slug"); slug != "" {
			http.Redirect(w, r, "/pages/"+slug, http.StatusMovedPermanently)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Legacy /post?slug=... redirect
	if path == "/post" || strings.HasSuffix(path, "/post") {
		q := r.URL.Query()
		if slug := q.Get("slug"); slug != "" {
			catSlug := q.Get("category")
			if catSlug == "" || h.b.GetCategoryBySlug(catSlug) == nil {
				_, catSlug = h.b.ResolveSlugViaIndex(slug)
			}
			http.Redirect(w, r, h.cfg.PostPath(catSlug, slug), http.StatusMovedPermanently)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Legacy Blogger URL resolution (/YYYY/MM/slug[.html])
	if post, resolved := h.b.ResolveOldURL(path); resolved {
		http.Redirect(w, r, h.cfg.PostPath(post.CategorySlug, post.Slug), http.StatusMovedPermanently)
		return
	}

	// Home / search
	h.serveIndex(w, r)
}

// ─────────────────────────────────────────────────────────────────────────────
// Route handlers
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) serveIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	q := r.URL.Query()
	categorySlug := q.Get("category")
	searchQuery := q.Get("q")
	isSearch := q.Has("search") || searchQuery != ""
	page := max1(intParam(q.Get("page"), 1))

	base := h.baseTemplateData(r, "website")

	if isSearch {
		title := h.cfg.Labels.SearchTitle + " - " + h.cfg.BlogName
		if searchQuery != "" {
			title = fmt.Sprintf(h.cfg.Labels.SearchResultsTitle, searchQuery) + " - " + h.cfg.BlogName
		}
		list := h.b.SearchPosts(searchQuery, page)
		data := base
		data.PageTitle = title
		data.Posts = list.Posts
		data.Pagination = list.Pagination
		data.Query = searchQuery
		h.renderPage(w, r, start, "search.html", &data)
		return
	}

	if categorySlug != "" {
		// Valid categories are 301-redirected in ServeHTTP; unknown ones return 404.
		h.serve404(w, r, start)
		return
	}

	// Landing page — no post scanning
	data := base
	data.PageTitle = h.cfg.BlogName
	data.PageDescription = h.cfg.DefaultMetaDescription
	data.Categories = h.b.GetCategoriesSorted()
	data.IndexBlurb = h.indexBlurb
	data.JSONLD = buildWebSiteJSONLD(h.cfg)
	h.renderPage(w, r, start, "home.html", &data)
}

func (h *Handler) renderSinglePost(w http.ResponseWriter, r *http.Request, start time.Time, post *blog.Post, categorySlug string, base templateData) {
	// ETag / 304
	postMtime := int64(0)
	if t, err := time.Parse("2006-01-02", post.Date); err == nil {
		postMtime = t.Unix()
	}
	etag := fmt.Sprintf(`"%s"`, etagHash(post.Slug+post.Date))
	lastMod := time.Unix(postMtime, 0).UTC().Format(http.TimeFormat)
	w.Header().Set("Last-Modified", lastMod)
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", h.cacheControlPages())

	ifNoneMatch := strings.TrimSpace(r.Header.Get("If-None-Match"))
	ifModSince := r.Header.Get("If-Modified-Since")
	if ifNoneMatch == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	if ifModSince != "" {
		if t, err := http.ParseTime(ifModSince); err == nil && t.Unix() >= postMtime {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// Parse JS files from front matter
	var jsFiles []string
	if post.FrontMatter.JS != "" {
		jsFile := strings.TrimSpace(post.FrontMatter.JS)
		if jsFile != "" && assetExists(h.assetsFS, "js/"+jsFile) {
			jsFiles = []string{jsFile}
		}
	}

	// Parse tags
	var tags []string
	if post.FrontMatter.Tags != "" {
		for _, t := range strings.Split(post.FrontMatter.Tags, ",") {
			if trimmed := strings.TrimSpace(t); trimmed != "" {
				tags = append(tags, trimmed)
			}
		}
	}

	data := base
	data.PageTitle = post.Title + " - " + h.cfg.BlogName
	data.PageDescription = post.FrontMatter.Description
	data.Post = post
	data.CategorySlug = categorySlug
	data.JSFiles = jsFiles
	data.Tags = tags
	data.JSONLD = buildArticleJSONLD(h.cfg, post, base.Canonical, categorySlug)
	h.renderPage(w, r, start, "post.html", &data)
}

func (h *Handler) cacheControlPages() string {
	if !h.cfg.Cache.Enabled {
		return "no-store"
	}
	return fmt.Sprintf("public, max-age=%d", h.cfg.Cache.MaxAgePages)
}

func (h *Handler) serveFeedXML(w http.ResponseWriter, r *http.Request) {
	if !h.cfg.Feed.Enabled {
		http.NotFound(w, r)
		return
	}
	data, err := os.ReadFile(h.cfg.Feed.OutputFile)
	if err != nil {
		http.Error(w, "feed not available", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.Header().Set("Cache-Control", h.cacheControlPages())

	out, cleanup := responseWriter(w, r, http.StatusOK)
	defer cleanup()
	_, _ = out.Write(data)
}

func (h *Handler) serveFeedPage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	data := h.baseTemplateData(r, "website")
	data.PageTitle = "Feed — " + h.cfg.BlogName
	data.Posts = h.b.GetFeedPosts(h.cfg.Feed.MaxItems)
	h.renderPage(w, r, start, "feed.html", &data)
}

func (h *Handler) cacheControlAssets() string {
	if !h.cfg.Cache.Enabled {
		return "no-store"
	}
	return fmt.Sprintf("public, max-age=%d", h.cfg.Cache.MaxAgeAssets)
}

func (h *Handler) serve404(w http.ResponseWriter, r *http.Request, start time.Time) {
	data := h.baseTemplateData(r, "website")
	data.PageTitle = h.cfg.Labels.NotFoundTitle + " — " + h.cfg.BlogName
	h.renderPageStatus(w, r, start, http.StatusNotFound, "404.html", &data)
}

func (h *Handler) serveFavicon(w http.ResponseWriter, r *http.Request) {
	f, err := h.assetsFS.Open("favicon.ico")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil || st.IsDir() {
		http.NotFound(w, r)
		return
	}
	rs, ok := f.(io.ReadSeeker)
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", h.cacheControlAssets())
	http.ServeContent(w, r, "favicon.ico", st.ModTime(), rs)
}

// ─────────────────────────────────────────────────────────────────────────────
// Template rendering
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) renderPage(w http.ResponseWriter, r *http.Request, start time.Time, tmplName string, data *templateData) {
	h.renderPageStatus(w, r, start, http.StatusOK, tmplName, data)
}

func (h *Handler) renderPageStatus(w http.ResponseWriter, r *http.Request, start time.Time, status int, tmplName string, data *templateData) {
	// Render the inner page template into a buffer, then wrap with layout.
	var innerBuf strings.Builder
	if err := h.tmpl.ExecuteTemplate(&innerBuf, tmplName, data); err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data.Content = template.HTML(innerBuf.String())

	if h.cfg.ShowRenderTime {
		ms := math.Round(float64(time.Since(start).Microseconds()) / 1000.0)
		data.RenderTime = fmt.Sprintf("%.2f", ms)
		log.Printf("render %s %s %.2fms", r.Method, r.URL.RequestURI(), ms)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	out, cleanup := responseWriter(w, r, status)
	defer cleanup()

	if err := h.tmpl.ExecuteTemplate(out, "layout.html", data); err != nil {
		// Headers already sent — can't send a proper error response
		fmt.Fprintf(out, "\n<!-- render error: %v -->", err)
	}
}

// responseWriter sets Content-Encoding: gzip when supported (outside Lambda),
// writes any non-200 status code after headers are finalized, and returns the
// target writer plus a cleanup function that closes the gzip stream if used.
func responseWriter(w http.ResponseWriter, r *http.Request, status int) (io.Writer, func()) {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		if status != http.StatusOK {
			w.WriteHeader(status)
		}
		gz := gzip.NewWriter(w)
		return gz, func() { _ = gz.Close() }
	}
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	return w, func() {}
}

// ─────────────────────────────────────────────────────────────────────────────
// Template loading
// ─────────────────────────────────────────────────────────────────────────────

func mustLoadTemplates(tmplFS fs.FS) *template.Template {
	funcMap := template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		"formatDate": func(dateStr string) string {
			t, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return dateStr
			}
			return t.Format("January 2, 2006")
		},
		"printf": fmt.Sprintf,
		"defaultStr": func(dflt, s string) string {
			if s == "" {
				return dflt
			}
			return s
		},
		"postPreviewData": func(p blog.Post, catSlug string, cfg *config.Config) postPreviewData {
			slug := p.CategorySlug
			if slug == "" {
				slug = catSlug
			}
			return postPreviewData{Post: p, PostURL: cfg.PostPath(slug, p.Slug), Config: cfg}
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	entries, err := fs.ReadDir(tmplFS, ".")
	if err != nil {
		panic("server: cannot read templates dir: " + err.Error())
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") {
			continue
		}
		data, err := fs.ReadFile(tmplFS, e.Name())
		if err != nil {
			panic("server: cannot read template " + e.Name() + ": " + err.Error())
		}
		name := e.Name()
		t := tmpl.New(name)
		if _, err := t.Parse(string(data)); err != nil {
			panic("server: cannot parse template " + name + ": " + err.Error())
		}
	}
	return tmpl
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

// splitCSPHeader splits "Header-Name: value" into name, value for w.Header().Set.
func splitCSPHeader(h string) (string, string) {
	if idx := strings.Index(h, ":"); idx > 0 {
		return strings.TrimSpace(h[:idx]), strings.TrimSpace(h[idx+1:])
	}
	return h, ""
}

func buildCanonical(r *http.Request) string {
	scheme := "https"
	if r.TLS == nil && r.Header.Get("X-Forwarded-Proto") == "" {
		scheme = "http"
	}
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}
	host := r.Host
	if host == "" {
		return ""
	}
	uri := r.URL.RequestURI()
	// Remove ?page= to avoid duplicate canonicals on paginated pages
	uri = removePaginationParam(uri)
	return scheme + "://" + host + uri
}

func removePaginationParam(uri string) string {
	// Very lightweight: strip &page=N or ?page=N
	idx := strings.Index(uri, "?")
	if idx < 0 {
		return uri
	}
	base := uri[:idx]
	query := uri[idx+1:]
	var parts []string
	for _, kv := range strings.Split(query, "&") {
		if !strings.HasPrefix(kv, "page=") {
			parts = append(parts, kv)
		}
	}
	if len(parts) == 0 {
		return base
	}
	return base + "?" + strings.Join(parts, "&")
}

func cssVersion(assetsFS fs.FS, cssPath, buildCommit string) string {
	rel := strings.TrimPrefix(cssPath, "/assets/")
	rel = strings.TrimPrefix(rel, "assets/")
	if info, err := fs.Stat(assetsFS, rel); err == nil && !info.ModTime().IsZero() {
		return strconv.FormatInt(info.ModTime().Unix(), 10)
	}
	if buildCommit != "" && buildCommit != "unknown" {
		return buildCommit
	}
	return "0"
}

func assetExists(fsys fs.FS, rel string) bool {
	if strings.Contains(rel, "..") || strings.Contains(rel, "\x00") {
		return false
	}
	info, err := fs.Stat(fsys, rel)
	return err == nil && !info.IsDir()
}

func validSlug(s string) bool {
	if len(s) > 200 {
		return false
	}
	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

func intParam(s string, def int) int {
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func max1(n int) int {
	if n < 1 {
		return 1
	}
	return n
}

// etagHash returns a fast non-cryptographic FNV-1a hex digest for ETags.
func etagHash(s string) string {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum64())
}

func (h *Handler) serveCleanPage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	slug := strings.TrimPrefix(r.URL.Path, "/pages/")
	slug = strings.TrimSuffix(slug, "/")

	if slug == "" || !validSlug(slug) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	page := h.b.GetPage(slug)
	if page == nil {
		h.serve404(w, r, start)
		return
	}

	data := h.baseTemplateData(r, "website")
	data.PageTitle = page.Title + " — " + h.cfg.BlogName
	data.PageDescription = page.FrontMatter.Description
	data.Page = page
	h.renderPage(w, r, start, "page.html", &data)
}

func (h *Handler) serveCleanContent(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	relPath := strings.TrimPrefix(r.URL.Path, "/content/")
	relPath = strings.TrimSuffix(relPath, "/")

	if relPath == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// 1. Check if the entire relPath corresponds to a category folder
	if categorySlug, ok := h.cfg.CategorySlugByFolder(relPath); ok {
		cat := h.b.GetCategoryBySlug(categorySlug)
		if cat == nil {
			h.serve404(w, r, start)
			return
		}
		q := r.URL.Query()
		page := max1(intParam(q.Get("page"), 1))

		list := h.b.GetPosts(page, categorySlug)
		data := h.baseTemplateData(r, "website")
		data.PageTitle = h.cfg.BlogName + " - " + cat.BlogName
		data.PageDescription = cat.HeaderContent
		data.CurrentCategory = cat
		data.CategorySlug = categorySlug
		data.Posts = list.Posts
		data.Pagination = list.Pagination
		data.SubCategories = h.b.GetSubCategories(categorySlug)
		data.JSONLD = buildWebPageJSONLD(h.cfg, data.Canonical, data.PageTitle, cat.HeaderContent)
		h.renderPage(w, r, start, "category.html", &data)
		return
	}

	// 2. Check if it's a post. Split the last segment as the slug, and everything before it as the category folder.
	var folderPath, postSlug string
	lastSlash := strings.LastIndex(relPath, "/")
	if lastSlash != -1 {
		folderPath = relPath[:lastSlash]
		postSlug = relPath[lastSlash+1:]
	} else {
		postSlug = relPath
	}

	if postSlug == "" || !validSlug(postSlug) {
		h.serve404(w, r, start)
		return
	}

	var postCategorySlug string
	if folderPath != "" {
		var ok bool
		postCategorySlug, ok = h.cfg.CategorySlugByFolder(folderPath)
		if !ok {
			h.serve404(w, r, start)
			return
		}
	}

	post := h.b.GetPostBySlug(postSlug, postCategorySlug)
	if post == nil {
		h.serve404(w, r, start)
		return
	}

	base := h.baseTemplateData(r, "article")
	h.renderSinglePost(w, r, start, post, postCategorySlug, base)
}
