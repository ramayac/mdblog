package server

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	reCSSComment      = regexp.MustCompile(`(?s)/\*.*?\*/`)
	reCSSRuleBlock    = regexp.MustCompile(`([^{}]+)\{([^{}]+)\}`)
	reCSSClassSel     = regexp.MustCompile(`\.([a-zA-Z_-][a-zA-Z0-9_-]*)`)
	reLightDarkVar    = regexp.MustCompile(`(--[a-z0-9-]+)\s*:\s*light-dark\(\s*(#[0-9a-fA-F]{3,6})\s*,\s*(#[0-9a-fA-F]{3,6})\s*\)`)
	reHTMLInlineStyle = regexp.MustCompile(`\bstyle\s*=`)
	reHTMLInlineEvent = regexp.MustCompile(`\bon[a-z]+\s*=`)
	reHTMLInlineJS    = regexp.MustCompile(`(?i)<script(?:\s[^>]*)?>`)
	reQuerySafeHTML   = regexp.MustCompile(`\.Query[^}]*safeHTML`)
)

// TestCSSHygiene enforces structural CSS rules across all stylesheets:
// no @import waterfalls, no !important overrides, no duplicate consecutive blocks,
// and no orphan class selectors missing from templates/JS.
func TestCSSHygiene(t *testing.T) {
	cssFiles, err := filepath.Glob("../../assets/css/*.css")
	if err != nil || len(cssFiles) == 0 {
		t.Skip("assets/css/*.css not found")
	}

	knownClasses := collectKnownClasses(t)

	for _, path := range cssFiles {
		name := filepath.Base(path)
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		clean := reCSSComment.ReplaceAllString(string(raw), "")

		if strings.Contains(clean, "@import") {
			t.Errorf("%s: contains @import (load base.style.css via <link> in layout.html instead)", name)
		}
		if strings.Contains(clean, "!important") {
			t.Errorf("%s: contains !important override", name)
		}

		// Detect duplicate consecutive rule blocks (e.g. accidental copy-paste)
		blocks := reCSSRuleBlock.FindAllStringSubmatch(clean, -1)
		seenSig := make(map[string]int)
		for _, b := range blocks {
			sel := normalizeWS(b[1])
			body := normalizeWS(b[2])
			if strings.HasPrefix(sel, "@") {
				continue
			}
			sig := sel + "{" + body + "}"
			seenSig[sig]++
			if seenSig[sig] > 1 {
				t.Errorf("%s: duplicate rule block for selector %q", name, sel)
			}

			// Verify class selectors exist in templates, JS, or allowlist
			for _, m := range reCSSClassSel.FindAllStringSubmatch(sel, -1) {
				cls := m[1]
				if !knownClasses[cls] {
					t.Errorf("%s: unused/orphan CSS class selector .%s", name, cls)
				}
			}
		}
	}
}

// TestCSSWCAGContrast verifies WCAG AA contrast ratios (4.5:1 for normal text)
// for every theme in both light and dark modes by parsing light-dark() tokens
// layered on top of base.style.css.
func TestCSSWCAGContrast(t *testing.T) {
	baseTokens := parseLightDarkTokens(t, "../../assets/css/base.style.css")

	themes, err := filepath.Glob("../../assets/css/*.style.css")
	if err != nil || len(themes) == 0 {
		t.Skip("assets/css/*.style.css not found")
	}

	type pair struct {
		fg, bg string
		min    float64
	}
	checks := []pair{
		{"--color-text", "--color-bg", 4.5},
		{"--color-heading", "--color-bg", 4.5},
		{"--color-primary", "--color-bg", 4.5},
		{"--color-text-muted", "--color-bg", 4.5},
		{"--color-text", "--color-surface", 4.5},
	}

	for _, path := range themes {
		themeName := filepath.Base(path)
		merged := make(map[string][2]string, len(baseTokens))
		for k, v := range baseTokens {
			merged[k] = v
		}
		if themeName != "base.style.css" {
			for k, v := range parseLightDarkTokens(t, path) {
				merged[k] = v
			}
		}

		for _, c := range checks {
			fgPair, okFg := merged[c.fg]
			bgPair, okBg := merged[c.bg]
			if !okFg || !okBg {
				t.Errorf("%s: missing required light-dark() token %s or %s", themeName, c.fg, c.bg)
				continue
			}
			for modeIdx, modeName := range []string{"light", "dark"} {
				ratio, err := contrastRatioHex(fgPair[modeIdx], bgPair[modeIdx])
				if err != nil {
					t.Errorf("%s (%s): invalid hex for %s/%s: %v", themeName, modeName, c.fg, c.bg, err)
					continue
				}
				if ratio < c.min {
					t.Errorf("%s (%s): contrast %s (%s) on %s (%s) = %.2f:1, want >= %.1f:1 (WCAG AA)",
						themeName, modeName, c.fg, fgPair[modeIdx], c.bg, bgPair[modeIdx], ratio, c.min)
				}
			}
		}
	}
}

// TestTemplateCSPAndSecurityHygiene ensures templates contain no inline styles,
// no inline scripts, no inline event handlers, and never pipe .Query to safeHTML.
func TestTemplateCSPAndSecurityHygiene(t *testing.T) {
	tmpls, err := filepath.Glob("../../templates/*.html")
	if err != nil || len(tmpls) == 0 {
		t.Skip("templates/*.html not found")
	}

	for _, path := range tmpls {
		name := filepath.Base(path)
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		content := string(raw)

		if reHTMLInlineStyle.MatchString(content) {
			t.Errorf("%s: contains inline style= attribute (violates strict style-src 'self' CSP)", name)
		}
		if reHTMLInlineEvent.MatchString(content) {
			t.Errorf("%s: contains inline event handler attribute (violates strict script-src 'self' CSP)", name)
		}
		for _, tag := range reHTMLInlineJS.FindAllString(content, -1) {
			if !strings.Contains(strings.ToLower(tag), "src=") {
				t.Errorf("%s: contains inline <script> tag %q (violates strict script-src 'self' CSP)", name, tag)
			}
		}
		if reQuerySafeHTML.MatchString(content) {
			t.Errorf("%s: pipes .Query to safeHTML (reflected XSS risk)", name)
		}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func collectKnownClasses(t *testing.T) map[string]bool {
	t.Helper()
	known := map[string]bool{
		// Goldmark / runtime classes not written literally in templates
		"footnotes": true,
	}
	reClassAttr := regexp.MustCompile(`class="([^"]+)"`)
	files, _ := filepath.Glob("../../templates/*.html")
	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		for _, m := range reClassAttr.FindAllStringSubmatch(string(raw), -1) {
			for _, cls := range strings.Fields(m[1]) {
				known[cls] = true
			}
		}
	}
	return known
}

func parseLightDarkTokens(t *testing.T, path string) map[string][2]string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	clean := reCSSComment.ReplaceAllString(string(raw), "")
	tokens := make(map[string][2]string)
	for _, m := range reLightDarkVar.FindAllStringSubmatch(clean, -1) {
		tokens[m[1]] = [2]string{m[2], m[3]}
	}
	return tokens
}

func normalizeWS(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func contrastRatioHex(fgHex, bgHex string) (float64, error) {
	l1, err := hexLuminance(fgHex)
	if err != nil {
		return 0, err
	}
	l2, err := hexLuminance(bgHex)
	if err != nil {
		return 0, err
	}
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05), nil
}

func hexLuminance(h string) (float64, error) {
	h = strings.TrimPrefix(strings.TrimSpace(h), "#")
	if len(h) == 3 {
		h = fmt.Sprintf("%c%c%c%c%c%c", h[0], h[0], h[1], h[1], h[2], h[2])
	}
	if len(h) != 6 {
		return 0, fmt.Errorf("invalid hex color %q", h)
	}
	r, err1 := strconv.ParseUint(h[0:2], 16, 8)
	g, err2 := strconv.ParseUint(h[2:4], 16, 8)
	b, err3 := strconv.ParseUint(h[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, fmt.Errorf("invalid hex color %q", h)
	}
	return 0.2126*srgbLinear(float64(r)/255.0) +
		0.7152*srgbLinear(float64(g)/255.0) +
		0.0722*srgbLinear(float64(b)/255.0), nil
}

func srgbLinear(c float64) float64 {
	if c <= 0.03928 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}
