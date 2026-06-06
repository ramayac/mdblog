package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/buildfeed"
	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/buildsitemap"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/render"
	"github.com/ramayac/mdblog/internal/server"
)

// Injected via -ldflags at build time:
//
//	-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	blog.BuildVersion = version
	blog.BuildCommit = commit
	blog.BuildDate = date

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg := config.MustLoad("config.toml")

	switch os.Args[1] {
	case "serve":
		runServe(cfg)
	case "build-index":
		runBuildIndex(cfg)
	case "build-feed":
		runBuildFeed(cfg)
	case "build-sitemap":
		runBuildSitemap(cfg)
	case "render":
		runRender(cfg, os.Args[2:])
	case "request":
		runRequest(cfg, os.Args[2:])
	case "lint-links":
		runLintLinks(cfg)
	case "version":
		fmt.Printf("mdblog %s (%s) built %s\n", version, commit, date)
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %q\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: mdblog <subcommand> [args]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("  serve          Start HTTP server (default :8080, set PORT env to override)")
	fmt.Println("  build-index    Generate posts/posts.index.json")
	fmt.Println("  build-feed     Generate feed.xml (requires build-index to run first)")
	fmt.Println("  build-sitemap  Generate sitemap.xml and robots.txt (requires build-index)")
	fmt.Println("  render         Render a post to a standalone HTML file")
	fmt.Println("  request        Simulate a GET request to a relative URL and print to stdout")
	fmt.Println("  lint-links     Verify all internal post/page markdown links resolve successfully")
	fmt.Println("  version        Print version information")
}

func runServe(cfg *config.Config) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	b := blog.New(cfg)
	h := server.New(cfg, b)

	fmt.Printf("Listening on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, h); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func runBuildIndex(cfg *config.Config) {
	if err := buildindex.Build(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "build-index: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Post index built successfully.")
}

func runBuildFeed(cfg *config.Config) {
	if err := buildfeed.Build(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "build-feed: %v\n", err)
		os.Exit(1)
	}
}

func runBuildSitemap(cfg *config.Config) {
	if err := buildsitemap.Build(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "build-sitemap: %v\n", err)
		os.Exit(1)
	}
}

func runRender(cfg *config.Config, args []string) {
	if err := render.Run(cfg, args); err != nil {
		fmt.Fprintf(os.Stderr, "render: %v\n", err)
		os.Exit(1)
	}
}

func runRequest(cfg *config.Config, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: mdblog request <url>\n")
		os.Exit(1)
	}
	if err := render.Request(cfg, args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "request: %v\n", err)
		os.Exit(1)
	}
}

func runLintLinks(cfg *config.Config) {
	b := blog.New(cfg)
	filesChecked, errors := b.LintLinks()
	fmt.Printf("Linting completed. Checked %d markdown files.\n", filesChecked)
	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "Found %d broken links:\n", len(errors))
		for _, errStr := range errors {
			fmt.Fprintln(os.Stderr, "  - "+errStr)
		}
		os.Exit(1)
	}
	fmt.Println("No broken links found!")
}

