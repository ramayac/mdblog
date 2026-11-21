package buildindex

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/markdown"
)

// IndexPost is the structure written into content.index.json.
type IndexPost = blog.IndexPost

var (
	reFencedCode     = regexp.MustCompile("(?s)```[\\s\\S]*?```")
	reInlineCode     = regexp.MustCompile("`[^`]+`")
	reHeading        = regexp.MustCompile(`(?m)^#{1,6}\s+`)
	reImage          = regexp.MustCompile(`!\[[^\]]*\]\([^\)]+\)`)
	reLink           = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`)
	reBoldItalicStar = regexp.MustCompile(`\*{1,2}([^*]+)\*{1,2}`)
	reBoldItalicUnd  = regexp.MustCompile(`_{1,2}([^_]+)_{1,2}`)
	reWhitespace     = regexp.MustCompile(`\s+`)
)

// Build scans all post directories, extracts metadata, and writes
// posts.index.json to the path specified in cfg.PostIndexFile.
func Build(cfg *config.Config) error {
	var entries []IndexPost

	// Uncategorized posts (root posts dir)
	if cfg.ShowUncategorized {
		root, err := scanFolder(cfg.PostsDir, "", cfg.DateFormat)
		if err != nil {
			log.Printf("buildindex: scan root: %v", err)
		}
		entries = append(entries, root...)
	}

	// Categorized posts
	for slug := range cfg.Categories {
		dir := filepath.Join(cfg.PostsDir, cfg.CategoryFolder(slug))
		catEntries, err := scanFolder(dir, slug, cfg.DateFormat)
		if err != nil {
			log.Printf("buildindex: scan %s: %v", dir, err)
			continue
		}
		entries = append(entries, catEntries...)
	}

	// Sort by date descending (newest first) — pre-sorted so runtime skips sorting
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date > entries[j].Date
	})

	// Atomic write: tmp → rename
	indexPath := cfg.PostIndexFile
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("buildindex: marshal: %w", err)
	}
	if err := blog.AtomicWrite(indexPath, data); err != nil {
		return fmt.Errorf("buildindex: write %s: %w", indexPath, err)
	}

	fmt.Printf("buildindex: wrote %d posts to %s\n", len(entries), indexPath)
	return nil
}

// scanFolder reads all .md files in dir and extracts post metadata.
func scanFolder(dir, categorySlug, dateFormat string) ([]IndexPost, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}

	var posts []IndexPost
	for _, e := range entries {
		if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), ".md") {
			continue
		}
		// Skip the optional landing-page blurb
		if e.Name() == "index.md" && categorySlug == "" {
			continue
		}

		fullPath := filepath.Join(dir, e.Name())
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Printf("buildindex: cannot read %s, skipping: %v", fullPath, err)
			continue
		}

		// Repair invalid UTF-8 before parsing
		if !utf8.ValidString(string(content)) {
			content = []byte(strings.ToValidUTF8(string(content), "\uFFFD"))
		}

		meta := markdown.ParseMetaOnly(string(content))
		fm := meta.FrontMatter

		slug := blog.SlugFromFilename(e.Name())
		title := fm.Title
		if title == "" {
			title = blog.TitleFromFilename(e.Name())
		}

		date := fm.Date
		if date == "" {
			info, err := e.Info()
			if err == nil {
				date = info.ModTime().Format(dateFormat)
			} else {
				date = time.Now().Format(dateFormat)
			}
		}

		excerpt := fm.Description
		if excerpt == "" {
			excerpt = rawMarkdownExcerpt(meta.Body, 200)
		}

		// Source path relative to cwd
		cwd, _ := os.Getwd()
		sourcePath := fullPath
		if cwd != "" {
			if rel, err := filepath.Rel(cwd, fullPath); err == nil {
				sourcePath = rel
			}
		}

		posts = append(posts, IndexPost{
			Slug:         slug,
			Title:        title,
			Date:         date,
			Author:       fm.Author,
			Tags:         fm.Tags,
			Excerpt:      excerpt,
			CategorySlug: categorySlug,
			SourcePath:   sourcePath,
			Filename:     e.Name(),
		})
	}
	return posts, nil
}

// rawMarkdownExcerpt strips Markdown syntax to produce a plain-text excerpt.
func rawMarkdownExcerpt(raw string, length int) string {
	text := reFencedCode.ReplaceAllString(raw, "")
	text = reInlineCode.ReplaceAllString(text, "")
	text = reHeading.ReplaceAllString(text, "")
	text = reImage.ReplaceAllString(text, "")
	text = reLink.ReplaceAllString(text, "$1")
	text = reBoldItalicStar.ReplaceAllString(text, "$1")
	text = reBoldItalicUnd.ReplaceAllString(text, "$1")
	text = reWhitespace.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	runes := []rune(text)
	if len(runes) <= length {
		return text
	}
	cut := string(runes[:length])
	if idx := strings.LastIndex(cut, " "); idx > 0 {
		cut = cut[:idx]
	}
	return cut + "..."
}
