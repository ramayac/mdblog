# MDBlog Agent Shim

This repository uses a persistent wiki for architecture documentation, agent instructions, and operational logs.

**CRITICAL: Do not look for detailed project instructions in this file.**

Instead, always start by reading the repository wiki:

1. **[wiki/index.md](wiki/index.md)** — The entry point for all documentation.
2. **[wiki/agents.md](wiki/agents.md)** — Detailed instructions, architecture, coding conventions, and workflows for AI agents.
3. **[wiki/log.md](wiki/log.md)** — Recent changes, audit records, and maintenance logs.
4. **[wiki/schema.md](wiki/schema.md)** — Required wiki structure and maintenance contract.

## Quick Start for Agents

- Read [wiki/repo-map.md](wiki/repo-map.md) for the current codebase architecture (`internal/config` URL helpers, cached `content.index.json`, `http.ServeMux` + `http.FileServerFS`, injected `server.Options`, and `mdblog build`).
- Use the global `wiki-engine` CLI (`wiki-engine list`, `search`, `candidates`, `lint`, `refresh`) for wiki tasks.
- Validate changes with `make lint` (`go vet`, config check, internal link linter) and `make test` (`go test ./...`).
- Follow role-specific and on-demand workflows documented in [prompts/](prompts/).
