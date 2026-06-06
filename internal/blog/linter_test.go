package blog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/config"
)

func TestLintLinks(t *testing.T) {
	dir := t.TempDir()

	// Create subdirectories
	techDir := filepath.Join(dir, "tech")
	pagesDir := filepath.Join(dir, "pages")

	// Create post 1 (valid internal links)
	writePost(t, techDir, "post1.md", `---
title: Post One
date: 2026-01-01
---
Check out [Home](/)
Link to [post2](/post?slug=post2&category=tech)
Link to [about](/page?slug=about)
Link to [legacy](/2026/01/post2)
Link to [relative](post2.md)
Link to [external](https://google.com)
Link to [mailto](mailto:test@example.com)
Link to [anchor](#section)
`)

	// Create post 2 (referenced by post 1)
	writePost(t, techDir, "post2.md", `---
title: Post Two
date: 2026-01-15
---
Post Two body.
`)

	// Create page about
	writePost(t, pagesDir, "about.md", `---
title: About
---
About page.
`)

	// Create post 3 with broken links
	writePost(t, techDir, "post3.md", `---
title: Post Three
date: 2026-01-20
---
Broken [post](/post?slug=nonexistent&category=tech)
Broken [page](/page?slug=nowhere)
Broken [legacy](/2020/01/not-found)
Broken [relative](missing.md)
`)

	cfg := makeTestConfig(dir)
	cfg.PagesDir = pagesDir
	cfg.Categories["tech"] = config.Category{BlogName: "Tech", Folder: "tech", Index: true}

	// Write index JSON
	idx := `[
		{"slug":"post1","title":"Post One","date":"2026-01-01","category_slug":"tech","filename":"post1.md"},
		{"slug":"post2","title":"Post Two","date":"2026-01-15","category_slug":"tech","filename":"post2.md"},
		{"slug":"post3","title":"Post Three","date":"2026-01-20","category_slug":"tech","filename":"post3.md"}
	]`
	_ = os.WriteFile(cfg.PostIndexFile, []byte(idx), 0644)

	b := New(cfg)

	// Temporarily change working directory so relative file checks resolve relative to TempDir
	oldCwd, err := os.Getwd()
	if err == nil {
		_ = os.Chdir(dir)
		defer func() { _ = os.Chdir(oldCwd) }()
	}

	filesChecked, errors := b.LintLinks()

	if filesChecked != 4 {
		t.Errorf("expected 4 files checked, got %d", filesChecked)
	}

	expectedErrors := 4
	if len(errors) != expectedErrors {
		t.Errorf("expected %d errors, got %d:\n%s", expectedErrors, len(errors), strings.Join(errors, "\n"))
	}

	expectedSubstring := []string{
		"post3.md:5: link \"/post?slug=nonexistent&category=tech\" is broken",
		"post3.md:6: link \"/page?slug=nowhere\" is broken",
		"post3.md:7: link \"/2020/01/not-found\" is broken",
		"post3.md:8: link \"missing.md\" is broken",
	}

	for _, sub := range expectedSubstring {
		found := false
		for _, errStr := range errors {
			if strings.Contains(errStr, sub) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected error containing %q, but got none", sub)
		}
	}
}
