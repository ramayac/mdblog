# Technical Plan: Clean Category and Page Slugs

This document outlines the design and implementation details to transition the site's URL structure to clean, SEO-friendly paths, matching the folder structure in `posts/` (which will also be renamed to `content/`).

Before each step, create a unit test, and after all steps, run integration tests to verify the entire flow from URL resolution to content rendering works as expected. Youre development toop is: test -> code -> check -> pass -> next step.
At the end of all the tasks, run a full suite of tests and do a local manual check of the site to verify all links and pages render correctly with the new URL schema.
And document the changes in the wiki for future reference and onboarding, use a new page called `slugs.md` to track the implementation details and rationale for this change.

---

## 1. Target URL Schema

| Content Type | Legacy URL | Clean URL | File System Path (after rename) |
|---|---|---|---|
| **Category List** | `/?category=srbyte` | `/content/writings/srbyte/` | `content/writings/srbyte/` |
| **Category Post** | `/post?slug=my-post&category=srbyte` | `/content/writings/srbyte/my-post` | `content/writings/srbyte/my-post.md` |
| **Uncategorized Post** | `/post?slug=my-post` | `/content/my-post` | `content/my-post.md` |
| **Standalone Page** | `/page?slug=about` | `/pages/about` | `pages/about.md` |

---

## 2. Phase 1: Rename `posts/` to `content/`

The primary content folder will be renamed from `posts/` to `content/`. This matches the dry-run plan already documented in [wiki/posts-to-content-dryrun.md](posts-to-content-dryrun.md).

### File Modifications
- **[config.toml](../config.toml)**: Point `posts_dir` to `"content"` and `post_index_file` to `"content/content.index.json"`.
- **[internal/config/config.go](../internal/config/config.go)**: Update fallback Go values if defaults are omitted.
- **[Makefile](../Makefile)**: Update paths in `build-index` and `new-post` targets.
- **[internal/blog/blog_test.go](../internal/blog/blog_test.go)**: Update relative test path references.
- **Dockerfiles ([Dockerfile](../Dockerfile), [Dockerfile.embed](../Dockerfile.embed))**: Copy `content/` instead of `posts/` into the image.
- **[.github/workflows/ghcr-release.yml](../.github/workflows/ghcr-release.yml)**: Update push triggers path to `content/**/*.md`.
- **[.gitignore](../.gitignore)** & **[.wikirc](../.wikirc)**: Ignore `content/content.index.json` and ignore `content/` in wiki searches.

---

## 3. Phase 2: Route Resolution & Backward Compatibility

Routing modifications will reside entirely in `internal/server/handler.go`.

### Clean Route Resolution in `ServeHTTP`
```go
// In internal/server/handler.go: ServeHTTP()

// 1. Clean Standalone Pages (/pages/<slug>)
if strings.HasPrefix(path, "/pages/") {
    h.serveCleanPage(w, r)
    return
}

// 2. Clean Content Routes (/content/*)
if strings.HasPrefix(path, "/content/") {
    h.serveCleanContent(w, r)
    return
}
```

### Handler Implementations

#### 1. `serveCleanPage`
- Trim `/pages/` prefix and trailing slash from `r.URL.Path` to obtain the page `slug`.
- Lookup the page with `h.b.GetPage(slug)`.
- If page exists, render `page.html`. Otherwise, serve 404.

#### 2. `serveCleanContent`
- Trim `/content/` prefix and trailing slash from the request path.
- Check if the remaining path matches the configured `Folder` of any category in `config.toml` (e.g. `writings/srbyte`):
  - **Category Index**: If it matches a category folder exactly, serve that category's paginated list (same logic as legacy `/?category=<slug>`).
- If it does not match a category folder, split the path at the last `/`:
  - **Categorized Post**: If there is a `/` separating the path:
    - The prefix is the category `folder` (e.g. `writings/srbyte`), and the suffix is the post `slug` (e.g. `my-post`).
    - Resolve the category slug from the folder.
    - Fetch the post via `h.b.GetPostBySlug(slug, categorySlug)`.
    - If found, render `post.html`. Otherwise, serve 404.
  - **Uncategorized Post**: If there is no `/` separating the path:
    - The path segment is treated as a root-level post `slug`.
    - Fetch the post via `h.b.GetPostBySlug(slug, "")`.
    - If found, render `post.html`. Otherwise, serve 404.

---

## 4. Phase 3: Backward Compatibility (301 Permanent Redirects)

To maintain SEO rankings and prevent breaking external backlinks, the old query parameters will permanently redirect (HTTP 301) to the new clean paths.

### Redirection Rules in `ServeHTTP`
- **Legacy Category URLs (`/?category=srbyte`)**:
  - Redirect to `/content/writings/srbyte/` (lookup category folder from config).
- **Legacy Post URLs (`/post?slug=my-post&category=srbyte`)**:
  - Redirect to `/content/writings/srbyte/my-post`.
  - If uncategorized: redirect to `/content/my-post`.
- **Legacy Page URLs (`/page?slug=about`)**:
  - Redirect to `/pages/about`.

---

## 5. Phase 4: URL Generation Updates

All links produced by the blog must use the new clean schema.

### Modifications
1. **Nav Menu generation ([internal/blog/blog.go](../internal/blog/blog.go))**:
   - Modify `GetNavPinned` and `GetDropdownCategories` to output `/content/<category_folder>/` links instead of `/?category=<category_slug>`.
2. **Template Helpers ([internal/server/handler.go](../internal/server/handler.go))**:
   - Update `postPreviewData` to populate `PostURL` with the clean route: `/content/<folder>/<slug>` (or `/content/<slug>` if uncategorized).
3. **Templates**:
   - **[templates/home.html](../templates/home.html)** & **[templates/category.html](../templates/category.html)**:
     - Update category card links to point directly to `/content/<folder>/`.
     - Update pagination links to follow the clean format `/content/<folder>/?page=X`.
   - **[templates/post.html](../templates/post.html)** & **[templates/feed.html](../templates/feed.html)**:
     - Replace category query links with their clean category URL equivalents.
4. **Feeds & SEO Files**:
   - **[internal/buildfeed/buildfeed.go](../internal/buildfeed/buildfeed.go)**:
     - Update `buildPostURL` to format absolute URLs as `<baseURL>/content/<folder>/<slug>`.
   - **[internal/buildsitemap/buildsitemap.go](../internal/buildsitemap/buildsitemap.go)**:
     - Format sitemap URLs for categories as `<baseURL>/content/<folder>/`.
     - Format sitemap URLs for posts as `<baseURL>/content/<folder>/<slug>`.
5. **Markdown Link Linter ([internal/blog/linter.go](../internal/blog/linter.go))**:
   - Update `validateLink` to correctly parse and confirm the existence of targets for:
     - `/content/<folder>/`
     - `/content/<folder>/<post_slug>`
     - `/pages/<page_slug>`

---

## 6. Verification and Testing

1. **Unit & Integration Tests**:
   - Add test coverage in `internal/server/handler_test.go` verifying:
     - Clean page rendering (`/pages/about` returns `200 OK`).
     - Clean category rendering (`/content/writings/srbyte/` returns `200 OK`).
     - Clean post rendering (`/content/writings/srbyte/my-post` returns `200 OK`).
     - Legacy redirects (verify `301 Moved Permanently` headers and correct `Location` targets).
2. **Local Execution**:
   - Run `make test` to verify all tests pass.
   - Run `make build-index` to build the new content-based index.
   - Execute `make lint` to verify go styling and type safety.
   - Run `make serve` and navigate to local preview pages to visually check menu links, category cards, and post preview links.
3. **Markdown Link Linting**:
   - Execute `make lint-links` (or `go run ./cmd/mdblog lint-links`) to ensure all in-file markdown links are still valid or check if updates to existing markdown files are necessary.
