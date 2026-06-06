package render

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/server"
)

// captureStdout executes f while redirecting os.Stdout to capture its output.
func captureStdout(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String(), err
}

func TestRequest_Success(t *testing.T) {
	// Point templates at real repo directory
	server.TemplateFS = os.DirFS("../../templates")
	if _, err := fs.Stat(server.TemplateFS, "layout.html"); err != nil {
		t.Skipf("skipping: templates not found: %v", err)
	}

	dir := t.TempDir()

	// Setup basic config
	cfg := &config.Config{
		BlogName:      "Test Blog",
		Lang:          "en",
		PostsDir:      dir,
		PagesDir:      dir + "/pages",
		PostIndexFile: dir + "/posts.index.json",
		DateFormat:    "2006-01-02",
		CSSTheme:      "assets/css/default.style.css",
		MenuLinks:     []config.MenuLink{{Label: "Home", URL: "/"}},
		Labels: config.Labels{
			NotFoundTitle:   "404 — Post Not Found",
			NotFoundMessage: "The post you're looking for doesn't exist.",
		},
	}

	// Create directories
	if err := os.MkdirAll(cfg.PagesDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Write static page
	aboutContent := `---
title: Test About Page
description: About description
---
This is the test about page content.
`
	if err := os.WriteFile(cfg.PagesDir+"/about.md", []byte(aboutContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Build empty index so server runs OK
	if err := buildindex.Build(cfg); err != nil {
		t.Fatal(err)
	}

	// Capture output for Request("/page?slug=about")
	output, err := captureStdout(func() error {
		return Request(cfg, "/pages/about")
	})

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	// Verify headers and body
	if !strings.Contains(output, "HTTP/1.1 200 OK") {
		t.Errorf("Expected 200 OK status, got output:\n%s", output)
	}
	if !strings.Contains(output, "Content-Type: text/html; charset=utf-8") {
		t.Errorf("Expected Content-Type header, got output:\n%s", output)
	}
	if !strings.Contains(output, "Test About Page") {
		t.Errorf("Expected page title, got output:\n%s", output)
	}
	if !strings.Contains(output, "This is the test about page content.") {
		t.Errorf("Expected content in response body, got output:\n%s", output)
	}
}

func TestRequest_NotFound(t *testing.T) {
	server.TemplateFS = os.DirFS("../../templates")
	if _, err := fs.Stat(server.TemplateFS, "layout.html"); err != nil {
		t.Skipf("skipping: templates not found: %v", err)
	}

	dir := t.TempDir()

	cfg := &config.Config{
		BlogName:      "Test Blog",
		PostsDir:      dir,
		PagesDir:      dir + "/pages",
		PostIndexFile: dir + "/posts.index.json",
		DateFormat:    "2006-01-02",
		CSSTheme:      "assets/css/default.style.css",
		Labels: config.Labels{
			NotFoundTitle:   "404 — Post Not Found",
			NotFoundMessage: "The post you're looking for doesn't exist.",
		},
	}

	// Build index
	if err := buildindex.Build(cfg); err != nil {
		t.Fatal(err)
	}

	// Capture output for a non-existent page
	output, err := captureStdout(func() error {
		return Request(cfg, "/pages/does-not-exist")
	})

	if err == nil {
		t.Error("Expected error from Request for a 404 URL, but got nil")
	}

	if !strings.Contains(output, "HTTP/1.1 404 Not Found") {
		t.Errorf("Expected 404 Not Found in headers, got output:\n%s", output)
	}
}
