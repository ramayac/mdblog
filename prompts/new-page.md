---
description: "Create a static standalone page like About, Now, Uses, or Colophon. Use when: adding an about page, creating a static page, making a standalone page."
agent: "agent"
argument-hint: "What kind of page? (e.g. About, Now, Uses, Projects...)"
---

You are creating a standalone page for an MDBlog blog.

## Project context

- Read [config.toml](../config.toml) for `author_name`, `blog_name`, and `pages_dir` (default `pages`).
- Standalone pages are Markdown files under `pages/` (e.g. `pages/about.md`), served at the clean route `/pages/<slug>`. The legacy route `/page?slug=<slug>` still works and 301-redirects to the clean route.
- Pages use the same front matter as posts (title, date, author, description) and render with `templates/page.html` — no category, tags, or breadcrumbs.
- Pages are **not** indexed by `make build-index`, not in the RSS feed, and not in search. Add nav visibility with a `[[menu_links]]` entry in [config.toml](../config.toml).

## Front matter

```yaml
---
title: About
date: YYYY-MM-DD
author: <author_name from config.toml>
description: A short description of the page.
---
```

## Steps

1. Ask the user what kind of page they want (About, Now, Uses, etc.) if not already specified.
2. Create `pages/<slug>.md` with the front matter above and the Markdown body. Use a lowercase, hyphenated slug.
3. If the page should appear in the nav bar, add a `[[menu_links]]` entry to [config.toml](../config.toml), e.g.:

   ```toml
   [[menu_links]]
   label = "About 💡"
   url   = "/pages/about"
   ```

4. Verify locally with `make serve` at `/pages/<slug>`.
5. No index rebuild is needed — `make build-index` indexes posts only.

## Page type templates

**About page:** Brief intro → what you do → what this blog covers → contact/links.
**Now page** (nownownow.com style): What you're focused on right now — projects, reading, interests.
**Uses page:** Hardware, software, tools, and setup you use daily.
**Colophon:** How the blog is built, what tech stack powers it, credits.

## Landing page blurb

For a short homepage introduction instead of a standalone page, create `content/index.md` (no front matter, pure Markdown) — it renders above the category cards.
