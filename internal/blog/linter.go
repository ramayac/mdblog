package blog

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var mdLinkRegex = regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)

// LintLinks scans all posts and pages, validating internal links.
// Returns the count of checked files and a slice of detailed error messages.
func (b *Blog) LintLinks() (int, []string) {
	var errors []string
	filesChecked := 0

	// 1. Gather all files to check
	markdownFilesMap := make(map[string]bool)
	var markdownFiles []string

	addFile := func(path string) {
		abs, err := filepath.Abs(path)
		if err == nil {
			path = abs
		}
		if !markdownFilesMap[path] {
			markdownFilesMap[path] = true
			markdownFiles = append(markdownFiles, path)
		}
	}

	// Scan posts directory
	if fileExists(b.cfg.PostsDir) {
		_ = filepath.Walk(b.cfg.PostsDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
				addFile(path)
			}
			return nil
		})
	}

	// Scan pages directory
	if fileExists(b.cfg.PagesDir) {
		_ = filepath.Walk(b.cfg.PagesDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
				addFile(path)
			}
			return nil
		})
	}

	// 2. Validate links in each file
	for _, path := range markdownFiles {
		filesChecked++
		content, err := os.ReadFile(path)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: failed to read file: %v", path, err))
			continue
		}

		lines := strings.Split(string(content), "\n")
		for lineNum, line := range lines {
			matches := mdLinkRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				link := match[1]
				if b.isInternalLink(link) {
					if err := b.validateLink(link, path); err != nil {
						errors = append(errors, fmt.Sprintf("%s:%d: link %q is broken: %v", path, lineNum+1, link, err))
					}
				}
			}
		}
	}

	return filesChecked, errors
}

// isInternalLink checks if the link points internally
func (b *Blog) isInternalLink(link string) bool {
	link = strings.TrimSpace(link)
	if link == "" {
		return false
	}
	if strings.HasPrefix(link, "//") {
		return false
	}
	if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
		// If it's absolute, check if it points to localhost or the base URL
		if strings.HasPrefix(link, "http://localhost") {
			return true
		}
		baseURL := b.cfg.Feed.BaseURL
		if baseURL != "" && strings.HasPrefix(link, baseURL) {
			return true
		}
		return false
	}
	// Ignore mailto, tel, javascript, and anchor links
	if strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "tel:") || strings.HasPrefix(link, "javascript:") {
		return false
	}
	if strings.HasPrefix(link, "#") {
		return false
	}
	return true
}

// validateLink checks if the internal link correctly resolves
func (b *Blog) validateLink(link string, currentFile string) error {
	// If it has absolute URL prefix, strip it for internal resolution
	if strings.HasPrefix(link, "http://localhost") {
		u, err := url.Parse(link)
		if err != nil {
			return err
		}
		link = u.RequestURI()
	} else if baseURL := b.cfg.Feed.BaseURL; baseURL != "" && strings.HasPrefix(link, baseURL) {
		link = strings.TrimPrefix(link, baseURL)
		if !strings.HasPrefix(link, "/") {
			link = "/" + link
		}
	}

	u, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	path := u.Path
	query := u.Query()

	// 1. Home / Root routes
	if path == "/" || path == "" || path == "/feed" || path == "/feed.xml" || path == "/sitemap.xml" || path == "/robots.txt" {
		return nil
	}

	// 1b. Search Label routes (legacy blogger tag lists)
	if strings.HasPrefix(path, "/search/label/") {
		return nil
	}

	// 2. Post route
	if path == "/post" || path == "/post/" {
		slug := query.Get("slug")
		category := query.Get("category")
		if slug == "" {
			return fmt.Errorf("missing slug parameter")
		}
		if b.GetPostBySlug(slug, category) == nil {
			return fmt.Errorf("post not found")
		}
		return nil
	}

	// 3. Page route
	if path == "/page" || path == "/page/" {
		slug := query.Get("slug")
		if slug == "" {
			return fmt.Errorf("missing slug parameter")
		}
		if b.GetPage(slug) == nil {
			return fmt.Errorf("page not found")
		}
		return nil
	}

	// 3a. Clean pages route
	if strings.HasPrefix(path, "/pages/") {
		slug := strings.TrimPrefix(path, "/pages/")
		slug = strings.TrimSuffix(slug, "/")
		if slug == "" {
			return fmt.Errorf("missing page slug")
		}
		if b.GetPage(slug) == nil {
			return fmt.Errorf("page not found: %s", slug)
		}
		return nil
	}

	// 3b. Clean content route
	if strings.HasPrefix(path, "/content/") {
		rel := strings.TrimPrefix(path, "/content/")
		rel = strings.TrimSuffix(rel, "/")
		if rel == "" {
			return nil
		}
		var categorySlug string
		for slug, cat := range b.cfg.Categories {
			folder := cat.Folder
			if folder == "" {
				folder = slug
			}
			if folder == rel {
				categorySlug = slug
				break
			}
		}
		if categorySlug != "" {
			return nil
		}

		var folderPath, postSlug string
		lastSlash := strings.LastIndex(rel, "/")
		if lastSlash != -1 {
			folderPath = rel[:lastSlash]
			postSlug = rel[lastSlash+1:]
		} else {
			postSlug = rel
		}

		var postCategorySlug string
		if folderPath != "" {
			for slug, cat := range b.cfg.Categories {
				folder := cat.Folder
				if folder == "" {
					folder = slug
				}
				if folder == folderPath {
					postCategorySlug = slug
					break
				}
			}
			if postCategorySlug == "" {
				return fmt.Errorf("invalid category folder: %s", folderPath)
			}
		}

		if b.GetPostBySlug(postSlug, postCategorySlug) == nil {
			return fmt.Errorf("post not found: %s (category: %s)", postSlug, postCategorySlug)
		}
		return nil
	}

	// 4. Legacy/Alternative URL route
	if _, resolved := b.ResolveOldURL(path); resolved {
		return nil
	}

	// 5. Asset Links (e.g. /assets/...)
	if strings.HasPrefix(path, "/assets/") {
		targetPath := strings.TrimPrefix(path, "/")
		if fileExists(targetPath) {
			return nil
		}
		return fmt.Errorf("asset file not found")
	}

	// 6. File-relative links (e.g., another markdown file or directory resource)
	// Check relative to current file
	dir := filepath.Dir(currentFile)
	targetPath := filepath.Join(dir, path)
	if fileExists(targetPath) {
		return nil
	}
	// Check relative to repo root
	if fileExists(path) {
		return nil
	}

	return fmt.Errorf("route or file not found")
}
