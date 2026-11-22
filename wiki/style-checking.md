# Deterministic CSS & Layout Checking

MDBlog validates stylesheets, templates, and rendered page layout deterministically across **two complementary layers**, keeping the Go binary and `go.mod` completely free of frontend build tools.

## Layer 1: Pure-Go Static Style & Template Linter (`make lint-styles`)

Implemented in [internal/server/style_lint_test.go](../internal/server/style_lint_test.go). Runs automatically as part of `make lint` and `make test` (`<20ms`, zero external dependencies).

### Checks enforced
- **CSS hygiene (`TestCSSHygiene`)**:
  - Rejects `@import` in `assets/css/*.css` (base stylesheet is loaded in parallel via `<link>` in `templates/layout.html`).
  - Rejects `!important` overrides.
  - Detects duplicate rule blocks within the same stylesheet.
  - Detects unused/orphan `.class` selectors that do not appear in `templates/*.html` (or the Goldmark allowlist).
- **WCAG AA contrast (`TestCSSWCAGContrast`)**:
  - Parses `light-dark(#light, #dark)` tokens from `base.style.css` and overlays each theme stylesheet.
  - Computes WCAG 2.1 relative luminance and verifies `>= 4.5:1` contrast in **both light and dark modes** for:
    - `--color-text` on `--color-bg`
    - `--color-heading` on `--color-bg`
    - `--color-primary` on `--color-bg`
    - `--color-text-muted` on `--color-bg`
    - `--color-text` on `--color-surface`
- **Template CSP & security hygiene (`TestTemplateCSPAndSecurityHygiene`)**:
  - Rejects inline `style=` attributes, inline `<script>` blocks, and inline `on*=` event handlers (enforcing strict `script-src 'self'; style-src 'self';` CSP).
  - Rejects piping `.Query` to `safeHTML` (preventing reflected XSS).

## Layer 2: Headless Chromium Computed-Geometry Checker (`make lint-layout`)

Implemented in [scripts/lint-layout.mjs](../scripts/lint-layout.mjs) and wrapped by [scripts/lint-layout.sh](../scripts/lint-layout.sh) (excluded from Docker images via `.dockerignore`). Runs in GitHub Actions CI (`layout-lint` job in `.github/workflows/ci.yml`) and on demand locally.

### Matrix
- **Routes**: `/`, `/pages/about`, `/content/writings/personal/`, `/content/writings/substack/well-well-well`, `/content/writings/srbyte/srbyte-cambios-en-los-terminos-de-uso-de-google-shared-en`, `/content/writings/unalistaaparte/post-semana-3-ajedrez-blake-y-briboner-a`, `/?q=linux`, `/feed`, `/content/does-not-exist`
- **Viewports**: Desktop (`1280×800`) and Mobile (`390×844`)
- **Color modes**: `light` and `dark` (via `localStorage` `theme-mode`)

### Deterministic invariants
1. **Zero horizontal overflow**: `scrollWidth <= innerWidth + 1` on every route and viewport.
2. **Column alignment**: left and right edges of `.post-header`/`.blog-title`/`.category-title`, `.post-content`, `.post-navigation`, and `.site-footer` align within `2px` of `.site-menu`.
3. **Vertical rhythm**: gap from `.site-menu` bottom to primary page header is within `12px..56px`, and `.post-content > :first-child` has `margin-top <= 2px`.
4. **Typography bounds**: `.post-content` computed font size is `15px..19px`, line-height ratio is `1.5..1.9`, and `h1` is strictly larger than `.post-content`.
5. **Dark mode switch**: `light` and `dark` modes render distinct body background colors.
6. **Runtime cleanliness**: zero console errors, zero CSP violations (`securitypolicyviolation`), and zero `>= 400` responses for `/assets/*`.

### Running locally
```bash
# Layer 1 (always available, zero deps)
make lint-styles

# Layer 2 (requires Node 22+ and Playwright Chromium)
npm install -g playwright && npx playwright install chromium
make lint-layout
```
