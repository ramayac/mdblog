package buildsitemap

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
)

// ─────────────────────────────────────────────────────────────────────────────
// Sitemap XML structures
// ─────────────────────────────────────────────────────────────────────────────

type sitemapEntry struct {
	Loc        string `xml:"loc"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
	LastMod    string `xml:"lastmod,omitempty"`
}

type sitemapURLSet struct {
	XMLName xml.Name       `xml:"urlset"`
	XMLNS   string         `xml:"xmlns,attr"`
	URLs    []sitemapEntry `xml:"url"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Build
// ─────────────────────────────────────────────────────────────────────────────

// Build reads posts.index.json, writes sitemap.xml and robots.txt atomically.
func Build(cfg *config.Config) error {
	if !cfg.Sitemap.Enabled {
		fmt.Println("buildsitemap: sitemap disabled, skipping")
		return nil
	}

	base := strings.TrimRight(cfg.Feed.BaseURL, "/")
	if base == "" {
		return fmt.Errorf("buildsitemap: feed.base_url must be set when sitemap.enabled is true")
	}

	// Load post index
	posts, err := blog.LoadIndex(cfg.PostIndexFile)
	if err != nil {
		return fmt.Errorf("buildsitemap: load index %s: %w", cfg.PostIndexFile, err)
	}

	// Build URL list
	urls := []sitemapEntry{
		{
			Loc:        base + "/",
			ChangeFreq: cfg.Sitemap.ChangeFreqHome,
			Priority:   cfg.Sitemap.PriorityHome,
		},
	}

	// Category pages — collect unique slugs in index order
	seen := map[string]bool{}
	for _, p := range posts {
		if p.CategorySlug != "" && !seen[p.CategorySlug] {
			seen[p.CategorySlug] = true
			urls = append(urls, sitemapEntry{
				Loc:        base + cfg.CategoryPath(p.CategorySlug),
				ChangeFreq: cfg.Sitemap.ChangeFreqCategory,
				Priority:   cfg.Sitemap.PriorityCategory,
			})
		}
	}

	// Individual post pages
	for _, p := range posts {
		entry := sitemapEntry{
			Loc:        base + cfg.PostPath(p.CategorySlug, p.Slug),
			ChangeFreq: cfg.Sitemap.ChangeFreqPost,
			Priority:   cfg.Sitemap.PriorityPost,
		}
		if p.Date != "" {
			entry.LastMod = p.Date
		}
		urls = append(urls, entry)
	}
	// Standalone pages (from pages_dir)
	if cfg.PagesDir != "" {
		if entries, err := os.ReadDir(cfg.PagesDir); err == nil {
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(strings.ToLower(e.Name()), ".md") {
					slug := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
					u := base + "/pages/" + slug
					entry := sitemapEntry{
						Loc:        u,
						ChangeFreq: cfg.Sitemap.ChangeFreqPost, // fallback to post defaults
						Priority:   cfg.Sitemap.PriorityPost,
					}
					info, err := e.Info()
					if err == nil {
						entry.LastMod = info.ModTime().Format(cfg.DateFormat)
					}
					urls = append(urls, entry)
				}
			}
		}
	}


	urlset := sitemapURLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	xmlBytes, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		return fmt.Errorf("buildsitemap: marshal xml: %w", err)
	}
	output := []byte(xml.Header + string(xmlBytes) + "\n")

	if err := blog.AtomicWrite(cfg.Sitemap.OutputFile, output); err != nil {
		return fmt.Errorf("buildsitemap: write sitemap: %w", err)
	}
	fmt.Printf("buildsitemap: wrote %d URLs to %s\n", len(urls), cfg.Sitemap.OutputFile)

	// robots.txt
	robots := "User-agent: *\nAllow: /\nSitemap: " + base + "/" + cfg.Sitemap.OutputFile + "\n"
	if err := blog.AtomicWrite(cfg.Sitemap.RobotsFile, []byte(robots)); err != nil {
		return fmt.Errorf("buildsitemap: write robots: %w", err)
	}
	fmt.Printf("buildsitemap: wrote %s\n", cfg.Sitemap.RobotsFile)

	return nil
}
