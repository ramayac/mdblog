package server

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/buildfeed"
	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
)

// testSetup creates a minimal temp dir with posts and returns a ready Handler.
func testSetup(t *testing.T) *Handler {
	t.Helper()

	dir := t.TempDir()

	// Create srbyte category
	srbyteDir := dir + "/srbyte"
	if err := os.MkdirAll(srbyteDir, 0755); err != nil {
		t.Fatal(err)
	}

	for _, folder := range []string{"personal", "projects", "projects/android", "projects/tools"} {
		dirPath := filepath.Join(dir, folder)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			t.Fatal(err)
		}
		postContent := "---\ntitle: Post\ndate: 2026-01-01\nauthor: Rodrigo Amaya\n---\nPost content."
		if err := os.WriteFile(filepath.Join(dirPath, "post.md"), []byte(postContent), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create pages directory with an about page
	pagesDir := dir + "/pages"
	if err := os.MkdirAll(pagesDir, 0755); err != nil {
		t.Fatal(err)
	}
	aboutContent := `---
title: About Me
description: All about me.
---

# About Me

I am a test page.
`
	if err := os.WriteFile(pagesDir+"/about.md", []byte(aboutContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Write a known post
	postContent := `---
title: 12:34:56 7 8 9 y el tiempo
date: 2008-01-01
author: Rodrigo Amaya
tags: tiempo, linux
description: A post about time
---

# 12:34:56 7 8 9 y el tiempo

El tiempo es relativo.
`
	if err := os.WriteFile(srbyteDir+"/srbyte-12-34-56-7-8-9-y-el-tiempo.md", []byte(postContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Write a second post
	post2 := `---
title: Linux Commands
date: 2009-01-01
author: Rodrigo Amaya
tags: linux
description: Useful linux commands
---

# Linux Commands

Many useful commands.
`
	if err := os.WriteFile(srbyteDir+"/linux-commands.md", []byte(post2), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		BlogName:               "Rodrigo A.",
		AuthorName:             "Rodrigo Amaya",
		Lang:                   "en",
		BlogDescription:        "Wholesome Software Development.",
		PostsPerPage:           25,
		ExcerptLength:          200,
		ShowUncategorized:      false,
		PostsDir:               dir,
		PagesDir:               dir + "/pages",
		PostIndexFile:          dir + "/posts.index.json",
		DateFormat:             "2006-01-02",
		DefaultMetaDescription: "Test blog.",
		CSSTheme:               "assets/css/default.style.css",
		CSP: config.CSPConfig{
			Enabled: true,
			Header:  "Content-Security-Policy: default-src 'self';",
		},
		Cache: config.CacheConfig{
			Enabled:      true,
			MaxAgePages:  3600,
			MaxAgeAssets: 86400,
		},
		Feed: config.FeedConfig{
			Enabled:    true,
			BaseURL:    "https://example.com",
			MaxItems:   50,
			OutputFile: dir + "/feed.xml",
		},
		MenuLinks: []config.MenuLink{{Label: "Home", URL: "/"}},
		Categories: map[string]config.Category{
			"srbyte":   {BlogName: "Sr. Byte 👨‍💻", Folder: "srbyte", Index: false},
			"personal": {BlogName: "Personal", Folder: "personal", Index: true},
			"projects": {BlogName: "Projects", Folder: "projects", Index: true},
			"android":  {BlogName: "Android", Folder: "projects/android", Index: false},
			"tools":    {BlogName: "Tools", Folder: "projects/tools", Index: false},
		},
		Menu: config.MenuConfig{
			Dropdowns: []config.MenuDropdown{
				{
					Label: "Writings",
					Item: []config.MenuCategoryRef{
						{Category: "srbyte", Order: 1},
						{Category: "personal", Order: 2},
					},
				},
				{
					Label: "Projects",
					Item: []config.MenuCategoryRef{
						{Category: "android", Order: 1},
						{Category: "tools", Order: 2},
					},
				},
			},
		},
		Labels: config.Labels{
			ReadMore:             "Read more →",
			BackToAll:            "← Back to all posts",
			BackToCategory:       "← Back to %s",
			NotFoundTitle:        "404 — Post Not Found",
			NotFoundMessage:      "The post you're looking for doesn't exist.",
			NoPostsInCategory:    "No posts found in this category.",
			PaginationPrev:       "← Newer Posts",
			PaginationNext:       "Older Posts →",
			PageIndicator:        "Page %d of %d",
			AuthorBy:             "By %s",
			SearchTitle:          "Search",
			SearchPlaceholder:    "What are you looking for?",
			SearchButton:         "🔍 Search",
			SearchShowingResults: `Showing results for "%s"`,
			SearchEmptyQuery:     "Enter a keyword above to search through posts.",
			SearchNoResults:      "No posts found matching your query.",
			SearchResultsTitle:   `Search Results for "%s"`,
		},
	}

	// Build the post index
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex: %v", err)
	}

	// Build the RSS feed
	if err := buildfeed.Build(cfg); err != nil {
		t.Fatalf("buildfeed: %v", err)
	}

	b := blog.New(cfg)

	// Point template FS at our real templates directory
	TemplateFS = os.DirFS("../../templates")
	// Verify it's readable
	if _, err := fs.Stat(TemplateFS, "layout.html"); err != nil {
		t.Skipf("templates not found (running outside repo root): %v", err)
	}

	return New(cfg, b)
}

func get(h *Handler, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func TestHome(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Rodrigo A.") {
		t.Error("should contain blog name")
	}
	if !strings.Contains(body, `href="?category=personal"`) {
		t.Error("should contain category card for Personal")
	}
	if !strings.Contains(body, `href="?category=projects"`) {
		t.Error("should contain category card for Projects")
	}
	if strings.Contains(body, `href="?category=srbyte"`) {
		t.Error("should NOT contain category card link for srbyte (index=false)")
	}
	if strings.Contains(body, `href="?category=android"`) {
		t.Error("should NOT contain category card link for android (index=false)")
	}
	if strings.Contains(body, `href="?category=tools"`) {
		t.Error("should NOT contain category card link for tools (index=false)")
	}
	if strings.Contains(body, "What are you looking for?") {
		t.Error("should NOT contain search form on home page")
	}
}

func TestSubCategoryRendering(t *testing.T) {
	h := testSetup(t)

	// Verify projects category page renders sub-categories android and tools
	w := get(h, "/?category=projects")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()

	// Should contain the parent category title
	if !strings.Contains(body, "Projects") {
		t.Error("should contain parent category title Projects")
	}

	// Should contain sub-category cards for Android and Tools
	if !strings.Contains(body, `href="?category=android"`) {
		t.Error("should contain sub-category card link for Android")
	}
	if !strings.Contains(body, `href="?category=tools"`) {
		t.Error("should contain sub-category card link for Tools")
	}

	// Should NOT contain personal as it's not a subcategory of projects
	if strings.Contains(body, `href="?category=personal"`) {
		t.Error("should NOT contain category card link for Personal inside Projects page")
	}

	// Verify that Android category page does NOT show sub-categories but shows its posts
	w2 := get(h, "/?category=android")
	if w2.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w2.Code)
	}
	body2 := w2.Body.String()
	if !strings.Contains(body2, "Android") {
		t.Error("should contain category title Android")
	}
	if !strings.Contains(body2, "post-preview") {
		t.Error("should show post-preview for android category posts")
	}
}

func TestCategory(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?category=srbyte")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "post-preview") {
		t.Error("should show post list with post-preview cards")
	}
}

func TestSearchLayout(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?q=linux")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Search Results") {
		t.Errorf("should contain Search Results heading, got excerpt: %s", body[:500])
	}
	if !strings.Contains(body, "What are you looking for?") {
		t.Error("should contain search form")
	}
}

func TestSearchWorking(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?q=tiempo")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "post-preview") {
		t.Error("should find matching post")
	}
	if !strings.Contains(body, "12:34:56") {
		t.Errorf("should find the specific post, body excerpt: %s", body[:500])
	}
}

func TestPost(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/post?slug=srbyte-12-34-56-7-8-9-y-el-tiempo&category=srbyte")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "12:34:56") {
		t.Error("post title should be displayed")
	}
	if !strings.Contains(body, "Rodrigo Amaya") {
		t.Error("author should be displayed")
	}
}

func Test404(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/post?slug=does-not-exist")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Errorf("Content-Type = %q, want text/html", ct)
	}
	body := w.Body.String()
	if !strings.Contains(body, "404") {
		t.Error("should contain 404 in page")
	}
	if !strings.Contains(body, "Not Found") {
		t.Error("should contain 'Not Found' in page")
	}
	if !strings.Contains(body, "site-footer") {
		t.Error("should render the standard layout footer")
	}
}

func TestAssetPathTraversal(t *testing.T) {
	h := testSetup(t)
	for _, path := range []string{
		"/assets/../config.toml",
		"/assets/%2e%2e/config.toml",
	} {
		w := get(h, path)
		if w.Code == http.StatusOK {
			t.Errorf("path %q should be rejected, got 200", path)
		}
	}
}

func TestCSPHeader(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/")
	if csp := w.Header().Get("Content-Security-Policy"); csp == "" {
		t.Error("CSP header should be set")
	}
}

func TestFeedXML(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/feed.xml")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	ct := w.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/rss+xml") {
		t.Errorf("Content-Type = %q, want application/rss+xml", ct)
	}
	body := w.Body.String()
	if !strings.Contains(body, "<?xml") {
		t.Error("feed.xml should contain XML declaration")
	}
	if !strings.Contains(body, "<rss") {
		t.Error("feed.xml should contain <rss> element")
	}
	if !strings.Contains(body, "Rodrigo A.") {
		t.Error("feed.xml should contain the blog title")
	}
}

func TestFeedPage(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/feed")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "feed-table") {
		t.Error("feed page should contain feed-table")
	}
	if !strings.Contains(body, "/feed.xml") {
		t.Error("feed page should link to /feed.xml")
	}
	if !strings.Contains(body, "Linux Commands") {
		t.Error("feed page should list posts")
	}
}

func TestCacheControl_PostPage_Enabled(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/post?slug=srbyte-12-34-56-7-8-9-y-el-tiempo&category=srbyte")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	cc := w.Header().Get("Cache-Control")
	if cc != "public, max-age=3600" {
		t.Errorf("Cache-Control = %q, want %q", cc, "public, max-age=3600")
	}
}

func TestCacheControl_PostPage_Disabled(t *testing.T) {
	h := testSetup(t)
	h.cfg.Cache.Enabled = false
	w := get(h, "/post?slug=srbyte-12-34-56-7-8-9-y-el-tiempo&category=srbyte")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	cc := w.Header().Get("Cache-Control")
	if cc != "no-store" {
		t.Errorf("Cache-Control = %q, want %q", cc, "no-store")
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// /page route tests
// ─────────────────────────────────────────────────────────────────────────────

func TestStaticPage_Found(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/page?slug=about")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "About Me") {
		t.Error("page should contain the title")
	}
	if !strings.Contains(body, "I am a test page") {
		t.Error("page should contain the rendered body")
	}
}

func TestStaticPage_NotFound(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/page?slug=does-not-exist")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Errorf("Content-Type = %q, want text/html", ct)
	}
	if !strings.Contains(w.Body.String(), "site-footer") {
		t.Error("should render the standard layout footer")
	}
}

func TestStaticPage_EmptySlugRedirects(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/page")
	if w.Code != http.StatusFound {
		t.Errorf("status = %d, want 302 redirect", w.Code)
	}
	if loc := w.Header().Get("Location"); loc != "/" {
		t.Errorf("Location = %q, want /", loc)
	}
}

func TestStaticPage_PathTraversalSlug(t *testing.T) {
	h := testSetup(t)
	for _, slug := range []string{"../etc/passwd", "foo/bar"} {
		w := get(h, "/page?slug="+slug)
		if w.Code == http.StatusOK {
			t.Errorf("slug %q should not return 200", slug)
		}
	}
}

func TestStaticPage_NoMetadata(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/page?slug=about")
	body := w.Body.String()
	// Pages should not show post metadata like dates or category breadcrumbs
	if strings.Contains(body, "post-meta") {
		t.Error("static page should not contain post-meta section")
	}
	if strings.Contains(body, "post-tags") {
		t.Error("static page should not contain post-tags section")
	}
}

func TestStaticPage_NavDropdownRendered(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/page?slug=about")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "nav-dropdown") {
		t.Error("nav dropdown should be rendered on static page")
	}
	if !strings.Contains(body, "Writings") {
		t.Error("dropdown label 'Writings' should appear in nav")
	}
}

func TestLegacyAlternativeResolution(t *testing.T) {
	h := testSetup(t)
	
	// Test with .html suffix
	w1 := get(h, "/2008/01/12-34-56-7-8-9-y-el-tiempo.html")
	if w1.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w1.Code)
	}
	body1 := w1.Body.String()
	if !strings.Contains(body1, "12:34:56 7 8 9 y el tiempo") {
		t.Error("rendered page should contain the post title")
	}

	// Test without .html suffix
	w2 := get(h, "/2008/01/12-34-56-7-8-9-y-el-tiempo")
	if w2.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w2.Code)
	}
	body2 := w2.Body.String()
	if !strings.Contains(body2, "12:34:56 7 8 9 y el tiempo") {
		t.Error("rendered page should contain the post title")
	}
}

