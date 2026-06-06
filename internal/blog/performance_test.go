package blog

import (
	"encoding/json"
	"os"
	"testing"
	"time"
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
