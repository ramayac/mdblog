package render

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/server"
)

func setupRenderTest(t *testing.T) *config.Config {
	t.Helper()

	absTemplates, err := filepath.Abs("../../templates")
	if err != nil {
		t.Fatal(err)
	}
	server.TemplateFS = os.DirFS(absTemplates)
	if _, err := fs.Stat(server.TemplateFS, "layout.html"); err != nil {
		t.Skipf("skipping: templates not found: %v", err)
	}

	// Run inside a temp working directory so render/ output doesn't pollute the repo.
	workDir := t.TempDir()
	origWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workDir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(origWD) })

	postsDir := filepath.Join(workDir, "content")
	techDir := filepath.Join(postsDir, "tech")
	if err := os.MkdirAll(techDir, 0755); err != nil {
		t.Fatal(err)
	}

	post := "---\ntitle: Render Test Post\ndate: 2026-01-15\nauthor: Tester\n---\n\nHello rendered world."
	if err := os.WriteFile(filepath.Join(techDir, "2026-01-15-render-test.md"), []byte(post), 0644); err != nil {
		t.Fatal(err)
	}

	return &config.Config{
		BlogName:      "Render Blog",
		Lang:          "en",
		PostsDir:      postsDir,
		PagesDir:      filepath.Join(workDir, "pages"),
		PostIndexFile: filepath.Join(postsDir, "content.index.json"),
		DateFormat:    "2006-01-02",
		PostsPerPage:  10,
		ExcerptLength: 200,
		CSSTheme:      "/assets/css/default.style.css",
		Categories: map[string]config.Category{
			"tech": {BlogName: "Tech", Folder: "tech", Index: true},
		},
	}
}

func TestRun_RenderByFilename(t *testing.T) {
	cfg := setupRenderTest(t)

	if err := Run(cfg, []string{"2026-01-15-render-test.md"}); err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	outPath := filepath.Join("render", "2026-01-15-render-test.html")
	raw, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("expected output file %s: %v", outPath, err)
	}
	html := string(raw)
	if !strings.Contains(html, `<base href="../">`) {
		t.Error("expected <base href=\"../\"> tag in rendered HTML")
	}
	if !strings.Contains(html, "Render Test Post") || !strings.Contains(html, "Hello rendered world.") {
		t.Error("expected full rendered post HTML, not a redirect or empty body")
	}
}

func TestRun_RandomAndCategoryRandom(t *testing.T) {
	cfg := setupRenderTest(t)

	if err := Run(cfg, []string{"random"}); err != nil {
		t.Fatalf("Run(random) failed: %v", err)
	}
	if err := Run(cfg, []string{"tech", "random"}); err != nil {
		t.Fatalf("Run(tech, random) failed: %v", err)
	}
	if err := Run(cfg, []string{"missing-cat", "random"}); err == nil {
		t.Error("expected error for unknown category")
	}
}

func TestRun_EmptyArgsAndNotFound(t *testing.T) {
	cfg := setupRenderTest(t)

	if err := Run(cfg, nil); err != nil {
		t.Fatalf("Run(nil) should print usage without error, got: %v", err)
	}
	if err := Run(cfg, []string{"does-not-exist.md"}); err == nil {
		t.Error("expected error for missing post")
	}
	if got := orUncategorized(""); got != "(uncategorized)" {
		t.Errorf("orUncategorized('') = %q, want (uncategorized)", got)
	}
	if got := orUncategorized("tech"); got != "tech" {
		t.Errorf("orUncategorized('tech') = %q, want tech", got)
	}
}
