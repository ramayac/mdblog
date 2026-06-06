# Wiki Log

## [2026-06-06] ingest | added clean-urls script and Makefile target

- Created `scripts/clean-urls.py` containing domain cleaning logic to replace absolute `srbyte.com` URLs with relative root paths.
- Added `clean-urls` target to the `Makefile` and registered it in `.PHONY` and documented it in `README.md` and `wiki/agents.md`.

## [2026-06-06] ingest | implemented content-count subcommand and Makefile target

- Created `internal/blog/contentcount.go` containing logic to recursively scan content and page directories and format a text-based treeview of subdirectory file counts.
- Wired `content-count` into the `mdblog` CLI in `cmd/mdblog/main.go` and defined the `content-count` target in the `Makefile`.
- Verified execution of `make content-count` returning correct directory tree details and counts.

## [2026-06-06] ingest | added standalone pages to sitemap generation

- Updated `internal/buildsitemap/buildsitemap.go` to scan the pages directory and include standalone page routes (e.g. `/pages/about`, `/pages/privacy`, `/pages/support`) in the sitemap output.
- Re-run sitemap generation and verified that sitemap counts correctly increased from 665 to 668 URLs, matching both posts and standalone pages.

## [2026-06-06] ingest | created performance insights wiki page

- Created `wiki/performance.md` documenting key optimization strategies (redundancy elimination, index bypassing via clean slugs, embedded memory vs disk reads) and raw benchmark/simulation test results.
- Added legacy URL redirect resolution (`ResolveOldURL`) benchmark results, yielding an average latency of **4.02ms** using a shuffled dataset of 50+ dynamically loaded posts and failed paths.
- Linked the performance page in `wiki/index.md`.

## [2026-06-06] ingest | created default embedded templates and assets plan

- Created `wiki/embedded-path-plan.md` outlining the technical strategy to enable in-memory assets and templates by default in both local and production environments, while preserving dynamic disk-reloading during development.
- Linked the plan in `wiki/index.md`.

## [2026-06-06] ingest | renamed posts.index.json to content.index.json and verified parsing performance

- Renamed `posts.index.json` to `content.index.json` to reflect that it lives in the `content/` folder and indexes clean content slugs.
- Updated default path configurations in `config.toml`, `internal/config/config.go`, `Makefile`, and updated references across all Go unit tests.
- Created a parsing speed performance test and a randomized traffic simulation in `internal/blog/performance_test.go` running 500 iterations.
- Verified that parsing the ~400 KB `content.index.json` containing 659 posts averages 2.86ms.
- Verified that a realistic user traffic simulation (69% post reads, 18.4% category list views, 12.6% searches) averages **1.27ms** per request, proving that direct path lookups successfully bypass the index and minimize latency.

## [2026-06-06] ingest | merged description and excerpt in post index

- Removed the redundant `description` field from `posts.index.json` schema (`IndexPost` struct in `buildindex.go`, `indexPost` struct in `blog.go`, and `indexPost` struct in `buildfeed.go`).
- Updated `indexPostToPost` in `internal/blog/blog.go` to fallback/hydrate `FrontMatter.Description` from the post `Excerpt` when building post list responses.
- Cleaned up mocked post index JSON strings in unit tests (`internal/blog/blog_test.go` and `internal/buildindex/buildindex_test.go`) to conform to the new structure without the `description` key.
- Verified that all unit and integration tests pass successfully and regenerated `content/posts.index.json` without the `description` key.

## [2026-06-06] ingest | created clean category and page slugs technical plan

- Created `wiki/clean-slugs-plan.md` outlining the technical strategy to rename the `posts/` folder to `content/` and route clean URL slugs (e.g. `/content/category/post` and `/pages/about`) while preserving backward compatibility.
- Created git branch `feat/new-slugs-2` for the implementation phase.
- Registered the clean slugs plan in the wiki index.

## [2026-06-05] ingest | implemented URL mapping, link linter, CSS theme normalization, and root favicon serving

- **Legacy URL Native Resolution**:
  - Added `ResolveOldURL` method to `internal/blog/blog.go` that parses year/month/slug patterns (supporting both optional `.html` extension and extensionless formats) and resolves them to active posts.
  - Implemented prefix-limited Levenshtein matching, letter collapsing, and Spanish character normalization (`cleanSlug`) to handle Blogger date discrepancies, accents, and slug truncations.
  - Refactored post rendering in `internal/server/handler.go` to extract a reusable `renderSinglePost` helper, and updated `ServeHTTP` to serve alternative post URLs inline (returning `200 OK`).
- **Internal Link Linter**:
  - Created `internal/blog/linter.go` to scan all posts and pages, validating internal links (routes, relative files, and assets).
  - Wired `lint-links` subcommand into `cmd/mdblog/main.go`, `Makefile` (`lint` and `lint-links`), and `.github/workflows/ci.yml` to prevent future broken links in CI.
  - Ignored protocol-relative external links (`//...`) and bypassed/redirected legacy Blogger label paths (`/search/label/...` redirecting to `/?q=...`).
  - Swapped absolute `srbyte.com` links for clean, portable, and relative root paths in 121 markdown files and cleaned up remaining dead links.
- **Asset Rendering Fix**:
  - Normalized `css_theme` path inside config loading (`internal/config/config.go`) and server constructor (`internal/server/handler.go`) to guarantee a leading slash `/` if relative. This prevents browsers from requesting stylesheets relative to nested directories (like `/2010/05/assets/...`) on deep legacy paths, restoring correct page rendering.
- **Root Favicon serving**:
  - Added a specific route handler for `/favicon.ico` at the domain root, serving `/assets/favicon.ico` directly. This avoids returning the HTML homepage catch-all for automatic browser queries, fixing the browser-reported "corrupt/broken" favicon error.
- **Unit and Integration Testing**:
  - Added comprehensive unit tests in `internal/blog/blog_test.go`, `internal/config/config_test.go`, and `internal/server/handler_test.go` verifying legacy URL parsing, linter behavior, CSS path normalization, and the root favicon endpoint.

## [2026-06-03] ingest | created posts-to-content directory renaming dry run plan

- Created `wiki/posts-to-content-dryrun.md` detailing the step-by-step changes required to rename the primary post directory to `content/` across configurations, Makefiles, Dockerfiles, and CI/CD pipelines.
- Linked the dry run plan in the main `wiki/index.md` list.

## [2026-06-03] ingest | implemented parent and sub-category navigation structure

- Updated `internal/blog/blog.go` to filter `GetCategoriesSorted` by the `Index` configuration flag, so sub-categories are hidden from the homepage cards by default.
- Added `GetSubCategories` in `internal/blog/blog.go` to dynamically identify nested categories under a parent category by folder prefix (e.g. `projects/android` under parent `projects`).
- Added `SubCategories` field to `templateData` in `internal/server/handler.go` and populated it for category request rendering.
- Modified `templates/category.html` to render sub-category cards styled identically to homepage category cards.
- Configured sub-categories (`android`, `opensource`, `sketches`) to have `index = false` in `config.toml` so they are hidden from the landing page.
- Added unit tests covering sorted category filtering and sub-category retrieval in `internal/blog/blog_test.go` and `internal/server/handler_test.go`.
- Documented parent/sub-category navigation feature in `wiki/agents.md`.

## [2026-06-03] ingest | documented multiple navigation dropdowns support

- Updated `internal/config` and `internal/blog` to support a list of `[[menu.dropdowns]]` in `config.toml` instead of the old single `[menu.categories]`.
- Configured "Writings" and "Projects" dropdowns in `config.toml`, and registered `android`, `opensource`, and `sketches` categories.
- Updated all unit tests in `internal/blog` and `internal/server` to match the new configuration schema.
- Added sample/placeholder posts for the new categories (`android`, `opensource`, `sketches`).
- Updated `wiki/agents.md` documentation to describe the multiple dropdowns configuration schema.

## [2026-06-03] ingest | documented custom slugs and added CLI request subcommand

- Added `request` subcommand to simulate GET requests to relative URLs without starting the server, and wired it to `make request URL="/..."`.
- Added unit tests in `internal/render/request_test.go` to validate status headers and body responses.
- Documented slug extraction logic and the `request` command in `README.md`, `wiki/agents.md`, and `wiki/repo-map.md`.
- Registered the `projects` category mapping in `config.toml`.

## [2026-04-22] ingest | migrated AGENTS.md to wiki and linked prompts

- Moved root `AGENTS.md` to `wiki/agents.md` to consolidate agent instructions into the persistent knowledge layer.
- Created a shim `AGENTS.md` at the root that redirects agents to the wiki.
- Updated `wiki/index.md` to include `agents.md` and added a link to the `.github/prompts/` directory.
- Updated `wiki/repo-map.md` and self-references in `agents.md` to reflect the new location.
- Verified that `AGENTS.md` remains a primary instruction surface, now versioned and maintained as part of the wiki.

## [2026-04-16] ingest | migrated wiki maintenance from scripts to wiki-engine CLI

- Deleted all 8 `scripts/wiki-*.sh` helpers; wiki maintenance now runs through the global `wiki-engine` binary (`github.com/ramayac/go-wiki-engine`).
- Added `.wikirc` at the repo root to configure wiki dir, diff base (`master...HEAD`), log line count, and ignore patterns for `wiki-engine`.
- Makefile wiki targets unchanged externally but now delegate to `wiki-engine` subcommands; `wiki-ingest-candidates` renamed to `wiki-candidates`.
- Added `wiki/` and `scripts/` to `.dockerignore` so neither the wiki nor the now-deleted scripts directory affects Docker image builds.
- Updated prompts, operation pages, and all docs to reference `wiki-engine <subcommand>` instead of `make wiki-*` and `scripts/`.
- Removed `theme-preview.html` (standalone dev preview file, no longer needed).

## [2026-04-13] bootstrap | established repo-local wiki scaffold

- Added a dedicated `wiki/` directory with a stable index, log, schema, operations, and rollout plan.
- Defined MDBlog-specific exclusions so routine wiki maintenance ignores `posts/` unless explicitly requested.
- Added repo instructions so the agent reads the wiki before broad analysis and files durable findings back into it.

## [2026-04-13] ingest | added shell-first wiki helper targets

- Added `wiki-*` Make targets for listing, searching, diff-driven ingest, linting, and a combined refresh snapshot.
- Added plain `sh` helper scripts under `scripts/` so the workflow remains unix-friendly and portable.
- Documented the new manual entrypoints in the repo docs and wiki operations pages.

## [2026-04-13] ingest | added wiki slash prompts

- Added workspace prompt files for `wiki-refresh`, `wiki-ingest`, and `wiki-query` under `.github/prompts/`.
- Matched the prompt workflows to the shell-first wiki commands instead of introducing a second maintenance path.
- Documented the new on-demand prompt entrypoints in the repo docs and wiki overview.

## [2026-04-13] ingest | gated wiki refresh on actual branch changes

- Updated `make wiki-refresh` to exit early when `make wiki-ingest-candidates` finds no ingestable changes for the current diff range.
- Added the same short-circuit rule to the `wiki-refresh` prompt so chat-driven refreshes do not run a no-op maintenance cycle.

## [2026-04-13] query | documented template rendering architecture

- Expanded `wiki/repo-map.md` with the server-side template lifecycle used by MDBlog.
- Captured the split between route-specific inner templates, the shared `layout.html` wrapper, and the reusable `_post_preview.html` partial.
- Recorded that embed builds can swap template and asset loading from disk to embedded filesystems.

## [2026-04-13] query | documented embed.go purpose and wiring

- Expanded `wiki/repo-map.md` with the `embed.go` mechanism used by the Lambda embed variant.
- Recorded why `embed.go` lives at the repo root, how it exposes subtree-scoped filesystems, and where those filesystems are injected into the server.
- Noted that the embed build removes runtime `templates/` and `assets/` dependencies, but still keeps posts, pages, config, and generated SEO/feed files on disk.

## [2026-04-13] query | documented AWS runtime and no-database model

- Expanded `wiki/repo-map.md` with the AWS Lambda container runtime path and the role of API Gateway plus `algnhsa`.
- Recorded which content and generated artifacts are copied into the production image and how the embed variant differs.
- Documented why MDBlog does not need a database: Markdown files and the pre-built JSON index are the runtime data layer.

## [2026-04-13] query | documented testing model

- Expanded `wiki/repo-map.md` with the repo's Go testing approach and the `make test` workflow.
- Recorded that tests live under `internal/` and cover config, parsing, domain logic, build artifacts, handler behavior, and SEO behavior.
- Noted that the test flow regenerates derived artifacts before running `go test ./...`.

## [2026-04-13] query | documented Makefile target inventory

- Expanded `wiki/repo-map.md` with the repo's Makefile target groups for development, wiki maintenance, and Docker workflows.
- Recorded that `help` is the default goal and that `render` plus `new-post` are variable-driven command entrypoints.

## [2026-04-13] query | documented CSS theme system

- Expanded `wiki/repo-map.md` with how CSS theme selection works through `config.toml`, the layout template, and asset serving.
- Recorded the distinction between the shared `base.style.css` layer, the standalone default theme, and the anthropic theme that imports the base layer.
- Noted that light and dark mode switching is handled inside the selected stylesheet via CSS variables plus the `data-theme` attribute set by `layout.html`.

## [2026-04-13] query | documented compression and security model

- Expanded `wiki/repo-map.md` with the runtime compression path and the Lambda-specific handoff to API Gateway or CloudFront.
- Recorded the main in-repo security controls: CSP headers, safe Markdown rendering, `html/template` escaping, and path-traversal guards for posts, pages, and assets.
- Noted the deployment hardening angle of the minimal container image and baked-in content artifacts.

## [2026-04-13] query | documented feed, sitemap, and robots lifecycle

- Expanded `wiki/repo-map.md` with how `feed.xml`, `sitemap.xml`, and `robots.txt` are generated from the prebuilt post index.
- Recorded the difference between the machine-readable feed endpoint and the human-readable feed page, plus the lack of a dynamic fallback for `feed.xml`.
- Noted that sitemap and robots have dynamic dev fallbacks while production prefers the prebuilt files on disk.

## [2026-04-13] query | documented post structure

- Expanded `wiki/repo-map.md` with the Markdown post file format, filename convention, and recognized front matter keys.
- Recorded which fields are optional in practice because the runtime has title and date fallback behavior.
- Noted that per-post JavaScript comes from the `js` front matter key and `assets/js/`.

## [2026-04-13] query | documented exact Docker image contents

- Expanded `wiki/repo-map.md` with the exact files produced during the Docker build stage and copied into the standard and embed final images.
- Recorded the distinction between generated artifacts such as `feed.xml` and `sitemap.xml` versus copied runtime content such as `posts/` and `pages/`.

## [2026-04-13] query | documented search flow

- Expanded `wiki/repo-map.md` with the request flow for standalone search, including the `q` and `search=1` query parameters.
- Recorded that search is a case-insensitive substring match over indexed title, excerpt, and tags fields from `posts/posts.index.json`.
- Noted that search depends on the built index and does not fall back to a live filesystem scan when the index is missing.

## [2026-04-13] query | documented footer flow

- Expanded `wiki/repo-map.md` with how the shared footer is assembled from `footer_content`, build version metadata, and optional render timing.
- Recorded that footer content is Markdown-rendered from config and only appears on HTML pages that pass through `layout.html`.

## [2026-04-13] query | documented Markdown publication flow

- Expanded `wiki/repo-map.md` with the GitHub Actions publication path from `posts/**/*.md` changes to GHCR, ECR, and the Lambda function update.
- Recorded that the automatic push trigger is specific to `posts/**/*.md`, which is narrower than all Markdown files in the repository.

## [2026-04-13] query | documented docker compose local preview flow

- Expanded `wiki/repo-map.md` with how `docker-compose.yml` is used for local preview rather than the production Lambda entry path.
- Recorded that the compose service builds from the standard `Dockerfile`, runs `/mdblog serve`, exposes port `8080`, and carries local runtime hardening settings.
- Noted that `make docker-run` rebuilds through Compose while `make docker-run-release` reuses the same compose file without rebuilding.
