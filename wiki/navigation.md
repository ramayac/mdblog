# Parent/Sub-category Navigation & Structure

This document records the architectural decisions, design choices, and implementation details for the parent and sub-category navigation structure implemented on the `feat/new-menu` branch.

## Requirements & Scope
The objective was to support nesting categories visually while keeping the flat-file blog engine lightweight.
Specifically:
- Homepage index page (`/`) should only render parent-level category cards (e.g. `Personal` and `Projects`).
- Sub-categories (like `Android` or `Tools`) should be hidden from the main index.
- Clicking on a parent category card (like `Projects`) should render its sub-category cards.
- Sub-categories are dynamically identified by matching folder prefixes (e.g., folder `projects/tools` is a child of parent category `projects` with folder `projects`).
- Android category should be hidden completely from dropdowns and cards for now.
- Open Source category was renamed to `Tools`, and `mdblog` post was moved into it.

## Implementation Details

### Configuration (`config.toml`)
- The `index` flag under `[categories.<slug>]` determines whether a category is shown on the homepage card grid. Setting `index = false` hides it.
- Dropdown menus (`[[menu.dropdowns]]`) manage navigation bar dropdown references.
- `android` was removed from category lists and dropdown item arrays to hide it completely.
- `opensource` was replaced by `tools` category.

### Core Logic (`internal/blog/blog.go`)
- **Category Retrieval & Counting**: `GetCategories()` recursively walks each category's folder using `filepath.WalkDir` to count `.md` posts. This ensures that parent categories (e.g. `projects`) that do not contain posts directly but contain posts in nested sub-category folders are still identified as containing posts (count > 0).
- **Sorted Filter**: `GetCategoriesSorted()` filters out categories where `index = false` in `config.toml`.
- **Sub-category Lookup**: `GetSubCategories(parentSlug)` checks all active categories and matches folders starting with the parent folder path followed by a slash (e.g., `projects/tools` starts with `projects/`).

### Server Handlers & Templates (`internal/server/` & `templates/`)
- **Handler Data**: Added `SubCategories` field to `templateData` struct in `handler.go`. Populated using `blog.GetSubCategories(categorySlug)` when serving category lists.
- **Template Rendering (`category.html`)**: If `SubCategories` is present, it renders them as interactive cards at the top of the category view, maintaining identical styling to the homepage cards.

## Conversation History & Milestones
1. **Dropdown Refactor**: Migrated configuration from single category dropdown to multiple dropdown mappings under `[[menu.dropdowns]]`.
2. **Sub-category Filtering**: Updated landing page listing to exclude sub-categories, and added sub-category cards inside the parent category page template.
3. **Android Hiding & Tools Renaming**: Cleaned up active categories to hide `android` completely and rename `opensource` to `tools` (storing `mdblog.md` inside it).
4. **Writings Relocation**: Grouped the `personal`, `srbyte`, and `substack` directories on disk under a new `posts/writings/` parent directory to match the structural layout of `posts/projects/` and updated category configuration folders in `config.toml`. Left the guide category `posts/mdblog/` untouched.

## Dev Note: Compiled Go & Reloading Behavior during Category Moves

Unlike raw Markdown post content (which is loaded and parsed dynamically from the filesystem on every request), **Go package files (`*.go`) and category caches are compiled/stored inside the running process**. 

When restructuring categories or changing structural listing code (e.g. migrating `os.ReadDir` to `filepath.WalkDir` recursive scanning):
- Individual post edits will reload automatically on page refresh.
- Go code updates and the internal category cache (`categoriesCache`) **will not reload** until the dev server is restarted.
- When moving posts to sub-folders (like moving `mdblog.md` to `projects/tools/mdblog.md`), if the server isn't restarted, the old binary (relying on `os.ReadDir`) will report 0 posts directly in `projects/` and hide the card. Always restart the server using `make serve` after changing structural category configs or Go source files.
