package render

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/server"
)

// Request performs a mock GET request against the HTTP server handler for the
// specified relative URL (e.g. "/" or "/page?slug=privacy") and prints the
// response status, headers, and body directly to stdout.
func Request(cfg *config.Config, url string) error {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	b := blog.New(cfg)
	h := server.New(cfg, b)

	req := httptest.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("request: read response body: %w", err)
	}

	// Print HTTP status line
	fmt.Printf("HTTP/1.1 %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))

	// Print response headers
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
	fmt.Println()

	// Print body content
	fmt.Print(string(body))

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request returned status %d", resp.StatusCode)
	}

	return nil
}
