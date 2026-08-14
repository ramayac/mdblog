# Clean URL Slugs — Implementation Record

**Status:** shipped. This page records the implementation of [clean-slugs-plan.md](clean-slugs-plan.md).

## Canonical URL Schema

| Content Type | Clean URL |
|---|---|
| Category listing | `/content/<folder>/` |
| Categorized post | `/content/<folder>/<slug>` |
| Uncategorized post | `/content/<slug>` |
| Standalone page | `/pages/<slug>` |

Category folders come from the `folder` key of each `[categories.<slug>]` block in `config.toml` (e.g. `writings/srbyte`).

## Routing (`internal/server/handler.go`)

- `serveCleanContent` resolves `/content/*` paths: an exact match with a configured category folder renders the paginated category listing; otherwise the last path segment is treated as a post slug under the folder prefix.
- `serveCleanPage` resolves `/pages/<slug>` via `blog.GetPage`.
- Legacy routes are preserved via `301 Moved Permanently` redirects to their clean equivalents:
  - `/?category=<slug>` → `/content/<folder>/`
  - `/post?slug=<slug>&category=<slug>` → `/content/<folder>/<slug>` (or `/content/<slug>` when uncategorized)
  - `/page?slug=<slug>` → `/pages/<slug>`
  - Blogger year/month/slug paths (e.g. `/2008/07/my-post.html`) → `/content/...`
- `/search/label/<tag>` continues to redirect to `/?q=<tag>&search=true`.

## URL Generation

- Navigation and category links are emitted as `/content/<folder>/` (see `internal/blog/blog.go`).
- `postPreviewData` builds `PostURL` as `/content/<folder>/<slug>` (or `/content/<slug>` when uncategorized).
- `internal/buildfeed/buildfeed.go` (`buildPostURL`) and `internal/buildsitemap/buildsitemap.go` generate absolute clean URLs from `feed.base_url`, including `/pages/<slug>` for standalone pages.
- The markdown link linter (`make lint-links`) validates `/content/...`, `/pages/...`, legacy patterns, and asset paths.

## Backward Compatibility

- Old query-string and Blogger URLs keep working through 301 redirects, protecting SEO and external backlinks.
- `ResolveOldURL` performs fuzzy matching (prefix-limited Levenshtein, diacritics folding) against the prebuilt index before issuing the redirect.

## Content Layout

`content/` mirrors the clean route space:

- `guides/` — guide category
- `projects/` — parent folder (`android/`, `tools/`)
- `writings/` — parent folder (`personal/`, `srbyte/`, `substack/`, `unalistaaparte/`)
- root-level `.md` files — uncategorized posts (shown when `show_uncategorized = true`)
