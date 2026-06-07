# Dry Run Plan: Renaming `posts` to `content`

This document details the exact files and lines that would be modified to change the primary content directory from `posts/` to `content/` across the repository.

---

## 1. Directory Structure Changes

```diff
- posts/
+ content/
    guides/
    projects/
    writings/
```

---

## 2. Configuration & Go Defaults

### [config.toml](file:///home/ramayac/git/MDBlog/config.toml)
Updates the runtime directory configurations to point to `content`:
```diff
- posts_dir              = "posts"
+ posts_dir              = "content"
  pages_dir              = "pages"
- post_index_file        = "posts/posts.index.json"
+ post_index_file        = "content/posts.index.json"
```

### [internal/config/config.go](file:///home/ramayac/git/MDBlog/internal/config/config.go)
Updates the fallback Go configuration values if they are omitted in `config.toml`:
```diff
  if cfg.PostsDir == "" {
- 	cfg.PostsDir = "posts"
+ 	cfg.PostsDir = "content"
  }
...
  if cfg.PostIndexFile == "" {
- 	cfg.PostIndexFile = "posts/posts.index.json"
+ 	cfg.PostIndexFile = "content/posts.index.json"
  }
```

---

## 3. Developer Workflows & CLI

### [Makefile](file:///home/ramayac/git/MDBlog/Makefile)
Updates Makefile targets for index compilation and CLI scaffolding:
```diff
- build-index: ## Generate post metadata index (writes posts/posts.index.json)
+ build-index: ## Generate post metadata index (writes content/posts.index.json)
  	@echo "Building post metadata index..."
  	go run ./cmd/mdblog build-index
- 	@echo "Post index built successfully. (posts/posts.index.json)"
+ 	@echo "Post index built successfully. (content/posts.index.json)"
...
  new-post: ## Scaffold a new Markdown post (TITLE="Title" [CATEGORY=slug] [TAGS="t1, t2"])
- 	$(eval DIR     := $(if $(CATEGORY),posts/$(CATEGORY),posts))
+ 	$(eval DIR     := $(if $(CATEGORY),content/$(CATEGORY),content))
  	$(eval DATE    := $(shell date +%Y-%m-%d))
  	$(eval FILENAME:= $(DATE)-$(shell echo "$(TITLE)" | tr '[:upper:]' '[:lower:]' | tr -cd 'a-z0-9 ' | tr ' ' '-').md)
- 	@if [ -n "$(CATEGORY)" ] && [ ! -d "posts/$(CATEGORY)" ]; then \
- 		echo "Category folder not found: posts/$(CATEGORY)"; exit 1; \
+ 	@if [ -n "$(CATEGORY)" ] && [ ! -d "content/$(CATEGORY)" ]; then \
+ 		echo "Category folder not found: content/$(CATEGORY)"; exit 1; \
  	fi
```

### [internal/blog/blog_test.go](file:///home/ramayac/git/MDBlog/internal/blog/blog_test.go)
Updates the relative lookup path in structural tests:
```diff
  func TestPostTraversal_RealDir(t *testing.T) {
- 	postsDir := filepath.Join("..", "..", "posts")
+ 	postsDir := filepath.Join("..", "..", "content")
```

---

## 4. Docker & CI/CD Deployment

### [.github/workflows/ghcr-release.yml](file:///home/ramayac/git/MDBlog/.github/workflows/ghcr-release.yml)
Updates the automated git triggers:
```diff
  on:
    push:
      branches:
        - master
      paths:
-       - 'posts/**/*.md'
+       - 'content/**/*.md'
```

### [Dockerfile](file:///home/ramayac/git/MDBlog/Dockerfile)
Updates the paths copied into the production container image stage:
```diff
  # Copy compiled binary and content assets
  COPY --from=build /out/lambda      /lambda
  COPY --from=build /out/mdblog      /mdblog
- COPY --from=build /src/posts/      /posts/
+ COPY --from=build /src/content/    /content/
```

### [Dockerfile.debug](file:///home/ramayac/git/MDBlog/Dockerfile.debug)
Updates debug container copy stages:
```diff
- COPY --from=build /src/posts/       /posts/
+ COPY --from=build /src/content/     /content/
```

---

## 5. Git & Tool Configurations

### [.gitignore](file:///home/ramayac/git/MDBlog/.gitignore)
Updates git ignores:
```diff
- posts/posts.index.json
+ content/posts.index.json
```

### [.wikirc](file:///home/ramayac/git/MDBlog/.wikirc)
Instructs `wiki-engine` to ignore the new content directory from wiki ingests:
```json
  "ignorePatterns": [
-   "posts/",
+   "content/",
```

---

## 6. Wiki & Documentation

### [wiki/agents.md](file:///home/ramayac/git/MDBlog/wiki/agents.md)
Update directory map overview:
```diff
- posts/              # All blog post content lives here
+ content/            # All blog post content lives here
```

### [wiki/repo-map.md](file:///home/ramayac/git/MDBlog/wiki/repo-map.md)
Update architectural references to files and paths from `posts/` to `content/`.

---

## Assessment: How Clean is the Change?
This change is **highly clean and self-contained**. 
Because MDBlog relies on the standard Go library and config variables, there are **no hardcoded references** to `"posts"` in the core handler routing or template execution functions—they all dynamically check `cfg.PostsDir` and `cfg.PostIndexFile`.

Performing this refactoring sweep will not impact the blog's runtime logic, template structures, or styling.
