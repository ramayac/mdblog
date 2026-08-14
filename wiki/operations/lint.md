# Lint Workflow

## Goal

Keep the wiki coherent, linked, and current.

## Checks

- Pages mentioned in the index still exist.
- Internal wiki links resolve (no dead links to moved pages or folders, e.g. `../prompts/` not `../.github/prompts/`).
- Important repo areas have coverage.
- Stale claims are updated when source files changed.
- Exclusions still match repo reality (`.wikirc` ignores `content/`, `wiki/`, `bin/`, and generated artifacts).
- New recurring topics have a page instead of being trapped in chat history.

## Shell-First Checks

```bash
wiki-engine lint
wiki-engine list
wiki-engine search "TODO:"
```

## Repair Order

1. Fix stale or incorrect topic pages.
2. Fix `wiki/index.md` links or summaries.
3. Append a log entry if the lint changed durable content.

## Log Format

Use this exact heading pattern:

```md
## [YYYY-MM-DD] lint | short summary
```
