# Binary Size Investigation

**Date:** 2026-07-18
**Branch:** `feature/smaller-go-binary`
**Decision:** Use gc compiler flags only; reject TinyGo.

## Goal

Reduce the `mdblog` binary size as much as possible without code changes.

## Baseline

| Binary | Size (stripped) |
|---|---|
| `mdblog` | 9.6 MB |
| `lambda` | 9.8 MB |

## On-disk section breakdown (gc, stripped)

| Section | Size | What it is |
|---|---|---|
| `.text` | 4.2 MB | Compiled machine code |
| `.gopclntab` | 3.5 MB | Go runtime PC-line table (stack unwinding, GC) |
| `.rodata` | 1.2 MB | String literals, type descriptors, constant data |
| `.noptrdata` + `.data` | 0.4 MB | Global variables |
| **Total on disk** | **~9.4 MB** | |

Runtime-only (BSS, not on disk): `crypto/internal/fips140/drbg.memory` at 32 MB. This is Go 1.26's FIPS 140-3 DRBG pool, allocated at startup, cannot be disabled via build flags.

## Why 9.4 MB (vs ~1.6 MB for minimal Go binary)

The dependency chain from `net/http` dominates the binary:

```
net/http → crypto/tls → crypto/x509 → crypto/{ecdsa,rsa,ed25519,...}
                                    → encoding/asn1, math/big
         → mime, mime/multipart
         → net, net/url
         → vendor/golang.org/x/net/http2/hpack
         → vendor/golang.org/x/net/idna
         → vendor/golang.org/x/text/unicode/{bidi,norm}
```

Plus: `encoding/json`, `html/template`, `goldmark` (Markdown + GFM + footnotes), `go-toml/v2`.

These are **imported because the code uses them** — compiler flags alone cannot eliminate them.

## Compiler/linker flags evaluated

| Flag | Saving | Notes |
|---|---|---|
| `-ldflags="-s -w"` | 4.4 MB (14→9.6) | Strips DWARF + symbol table. **Already in use.** |
| `-ldflags="-funcalign 1"` | ~200 KB (2.1%) | Reduces function alignment from 16→1 byte. Negligible perf impact for a blog server. **Adopted.** |
| `-trimpath` | ~16 KB | Removes absolute filesystem paths from binary. Reproducible builds. **Adopted.** |
| `-buildvcs=false` | ~16 KB | Skips embedding VCS metadata. **Adopted.** |
| `-ldflags="-B none"` | 19 bytes | Omits GNU build-id note. Not worth it. |
| `-checklinkname=0` | 0 | No measurable difference. |
| `GOFIPS140=off` | 0 | Controls FIPS module version (frozen vs latest), not inclusion. |

## Result

| Binary | Before | After |
|---|---|---|
| `mdblog` | 9.6 MB | **9.4 MB** |
| `lambda` | 9.8 MB | **9.6 MB** |

**Total gain: ~200 KB (2.1%).** This is the ceiling for pure compiler/linker flags in Go 1.26.

## TinyGo evaluation

### Size

TinyGo 0.41.1 compiles the project to a **1.6 MB** stripped binary — an 83% reduction (5.9× smaller).

### Blockers

Three stdlib packages required by MDBlog are **missing from TinyGo entirely**:

| Package | Used by | Status |
|---|---|---|
| `html/template` | Server rendering, all page templates | ❌ Not in TinyGo stdlib |
| `encoding/json` | Blog index, feed, sitemap, SEO JSON-LD, build tools | ❌ Not in TinyGo stdlib |
| `compress/gzip` | HTTP response compression | ❌ Not in TinyGo stdlib |

The entire `encoding/` tree is absent. TinyGo's `reflect` package has 11+ methods that panic at runtime (`NumOut`, `NumIn`, `NumMethod`, `Method`, `MethodByName`, `ConvertibleTo`, `ChanOf`, etc.), which blocks any package that introspects types (templates, JSON marshalers, TOML decoder).

The `content-count` subcommand works correctly under TinyGo. The `serve` subcommand panics immediately:

```
panic: unimplemented: (reflect.Type).NumOut()
```

### What a TinyGo port would require

1. Replace `html/template` with `text/template` or a static renderer (high effort, safety risk)
2. Replace `encoding/json` with hand-rolled marshalers or a TinyGo-compatible library across 7 files (medium-high effort)
3. Drop or replace `compress/gzip` (low effort, acceptable behind a reverse proxy)
4. Verify `goldmark` and `go-toml/v2` compile and run under TinyGo's limited `reflect` (unknown risk)

## Decision

**Reject TinyGo.** The port would be a partial rewrite of the server layer for a 7.8 MB savings — not justified given the project's deployment model (Docker image, Lambda). The gc compiler flags (`-s -w -funcalign 1 -trimpath -buildvcs=false`) are applied to the Makefile and will carry forward.

## Makefile changes

Added to `LDFLAGS`: `-funcalign 1`
Added `GOBUILDFLAGS := -trimpath -buildvcs=false`, applied to all `go build` and `go run` targets.
