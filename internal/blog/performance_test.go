package blog

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ramayac/mdblog/internal/config"
)

func TestIndexParsingPerformance(t *testing.T) {
	// Try loading either content.index.json or posts.index.json
	path := "../../content/content.index.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "../../content/posts.index.json"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Skipf("skipping benchmark: index file not found at %s", path)
	}

	iterations := 500
	var totalDuration time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now()
		var posts []indexPost
		err := json.Unmarshal(data, &posts)
		if err != nil {
			t.Fatalf("failed to unmarshal JSON at iteration %d: %v", i, err)
		}
		totalDuration += time.Since(start)
	}

	avg := totalDuration / time.Duration(iterations)
	t.Logf("Performance test: parsed %s %d times", path, iterations)
	t.Logf("Total duration: %v", totalDuration)
	t.Logf("Average parsing duration per run: %v", avg)
}

func TestRealisticUserTrafficSimulation(t *testing.T) {
	// Load production config
	cfgPath := "../../config.toml"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		t.Skip("skipping simulation: config.toml not found")
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Adjust relative paths for testing directory context (../../)
	cfg.PostsDir = "../../" + cfg.PostsDir
	cfg.PagesDir = "../../" + cfg.PagesDir
	cfg.PostIndexFile = "../../" + cfg.PostIndexFile

	b := New(cfg)

	// Load posts from the index to get valid slugs and category slugs
	indexData, err := os.ReadFile(cfg.PostIndexFile)
	if err != nil {
		t.Fatalf("failed to read index file: %v", err)
	}
	var indexedPosts []indexPost
	if err := json.Unmarshal(indexData, &indexedPosts); err != nil {
		t.Fatalf("failed to unmarshal index: %v", err)
	}

	if len(indexedPosts) == 0 {
		t.Skip("no posts found in index for traffic simulation")
	}

	// Common search keywords based on existing post content
	searchKeywords := []string{"go", "google", "android", "blogger", "web", "markdown", "software", "development"}

	// Seed random generator to ensure reproducibility
	rng := rand.New(rand.NewSource(12345))

	iterations := 500
	var totalDuration time.Duration

	lookupCount := 0
	listCount := 0
	searchCount := 0

	for i := 0; i < iterations; i++ {
		op := rng.Float64()
		start := time.Now()

		if op < 0.70 {
			// 70% chance: Read a random post by slug (Simulates direct page load / reading)
			postIdx := rng.Intn(len(indexedPosts))
			p := indexedPosts[postIdx]
			_ = b.GetPostBySlug(p.Slug, p.CategorySlug)
			lookupCount++
		} else if op < 0.90 {
			// 20% chance: Load a listing page of a random category or home page
			catIdx := rng.Intn(len(cfg.Categories) + 1)
			catSlug := ""
			j := 0
			for slug := range cfg.Categories {
				if j == catIdx {
					catSlug = slug
					break
				}
				j++
			}
			_ = b.GetPosts(1, catSlug)
			listCount++
		} else {
			// 10% chance: Perform a search query
			query := searchKeywords[rng.Intn(len(searchKeywords))]
			_ = b.SearchPosts(query, 1)
			searchCount++
		}

		totalDuration += time.Since(start)
	}

	avg := totalDuration / time.Duration(iterations)
	t.Logf("Realistic traffic simulation results (over %d iterations):", iterations)
	t.Logf("  Total duration: %v", totalDuration)
	t.Logf("  Average request latency: %v", avg)
	t.Logf("  Operation mix: Lookups: %d (%.1f%%), Lists: %d (%.1f%%), Searches: %d (%.1f%%)",
		lookupCount, float64(lookupCount)/float64(iterations)*100,
		listCount, float64(listCount)/float64(iterations)*100,
		searchCount, float64(searchCount)/float64(iterations)*100)
}

func TestResolveOldURLPerformance(t *testing.T) {
	// Load production config
	cfgPath := "../../config.toml"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		t.Skip("skipping benchmark: config.toml not found")
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Adjust relative paths for testing directory context (../../)
	cfg.PostsDir = "../../" + cfg.PostsDir
	cfg.PagesDir = "../../" + cfg.PagesDir
	cfg.PostIndexFile = "../../" + cfg.PostIndexFile

	b := New(cfg)

	// Load posts from the index to generate 50 valid legacy URLs plus failed ones
	indexData, err := os.ReadFile(cfg.PostIndexFile)
	if err != nil {
		t.Fatalf("failed to read index: %v", err)
	}
	var posts []indexPost
	if err := json.Unmarshal(indexData, &posts); err != nil {
		t.Fatalf("failed to parse index: %v", err)
	}

	var urls []string
	for i, p := range posts {
		if i >= 50 {
			break
		}
		if len(p.Date) >= 7 {
			year := p.Date[0:4]
			month := p.Date[5:7]
			urls = append(urls, fmt.Sprintf("/%s/%s/%s.html", year, month, p.Slug))
		}
	}
	// Add failed resolution paths to test indexing/matching fail latency
	for i := 0; i < 10; i++ {
		urls = append(urls, fmt.Sprintf("/2009/10/non-existent-legacy-slug-%d.html", i))
	}

	// Shuffle the URLs list to simulate random user traffic
	rng := rand.New(rand.NewSource(54321))
	rng.Shuffle(len(urls), func(i, j int) {
		urls[i], urls[j] = urls[j], urls[i]
	})

	iterations := 500
	var totalDuration time.Duration

	for i := 0; i < iterations; i++ {
		url := urls[i%len(urls)]
		start := time.Now()
		_, _ = b.ResolveOldURL(url)
		totalDuration += time.Since(start)
	}

	avg := totalDuration / time.Duration(iterations)
	t.Logf("Legacy URL resolution performance test (ResolveOldURL):")
	t.Logf("  Total duration for %d requests: %v", iterations, totalDuration)
	t.Logf("  Average latency per legacy redirect resolution: %v", avg)
}

