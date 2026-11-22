# Performance and Benchmark Insights

This document outlines the performance characteristics, benchmark results, and architectural optimizations of MDBlog.

---

## 1. Executive Summary

MDBlog is a database-free, flat-file blog engine. To keep serving speeds fast without a database, it uses a pre-built metadata index (`content.index.json`) for lists/search, and direct filesystem lookups for single posts. 

Benchmarks under Go 1.26 show that MDBlog achieves:
* **HTTP Throughput (Home Page /):** **~10,389 requests/sec** (average latency **~0.96 ms**)
* **HTTP Throughput (Post Page):** **~5,418 requests/sec** (average latency **~1.84 ms**)
* **Average request latency under realistic traffic:** **~1.27 milliseconds**
* **Metadata index JSON parsing speed:** **~2.86 milliseconds**
* **Legacy URL redirect resolution latency:** **~4.02 milliseconds**
* **Asset-serving latency (in-memory embed):** **~10 to 100 nanoseconds**

---

## 2. Benchmarks

Benchmarks are implemented in [internal/blog/performance_test.go](../internal/blog/performance_test.go).

### Test A: Raw Index JSON Parsing
Measures the duration of reading the `content.index.json` from disk and parsing it into Go structs using `json.Unmarshal` over 500 sequential runs.
* **Index Size**: ~400 KB (659 posts)
* **Average Latency**: **2.86 milliseconds**

### Test B: Realistic User Traffic Simulation
Simulates production requests over 500 iterations using a randomized distribution mix:
* **70% Direct Post Lookups** (`GetPostBySlug` — file read + Markdown parsing)
* **20% Category Page Listings** (`GetPosts` — index load + pagination)
* **10% Search Queries** (`SearchPosts` — index load + case-insensitive substring matching)

#### Results:
* **Operation mix**: Lookups: 345 (69.0%), Lists: 92 (18.4%), Searches: 63 (12.6%)
* **Total duration**: 637.8 milliseconds
* **Average request latency**: **1.27 milliseconds**

Because 70% of requests bypass the metadata index using direct slug resolution, the overall average latency of the blog service is cut in half to **1.27ms**, confirming that the routing design is highly optimized for realistic usage patterns.

### Test C: Legacy URL Redirect Resolution (ResolveOldURL)
Measures the latency of parsing and resolving old Blogger paths (e.g. `/2008/07/musica-legalmente-gratuita.html`) to modern slugs via fuzzy matching over 500 iterations.
* **Average Latency**: **4.02 milliseconds**

#### Impact Analysis:
1. **CPU/Memory cost**: Parsing the 400 KB index, doing regex parsing, diacritics cleaning, and Levenshtein distance computations over 659 posts takes ~4.02ms.
2. **Parity**: While this is slightly higher than direct lookups (~1.27ms), it is still extremely fast and has negligible impact on standard user requests.
3. **Usage Pattern**: Legacy URLs are typically hit by search crawlers indexing old backlinks or visitors from old bookmarks, rather than primary navigators. Therefore, the overall runtime impact on the blog's server resources is virtually unnoticeable.

### Test D: Live HTTP Load Testing (ApacheBench)
Measures raw HTTP server performance under Go 1.26 (1000 requests, 10 concurrency) run locally via `make benchmark`:

#### Home Page (`/`)
* **Throughput**: **10,389.61 requests/second**
* **Average Latency**: **0.96 milliseconds**
* **99% Latency**: **4.0 milliseconds**

#### Post Page (`/content/writings/srbyte/srbyte-la-aquitectura-von-neumann`)
*(Involves reading markdown from disk and parsing/rendering via Goldmark on the fly)*
* **Throughput**: **5,418.67 requests/second**
* **Average Latency**: **1.84 milliseconds**
* **99% Latency**: **5.0 milliseconds**

---

## 3. Post-Optimization Results (2026-11-21)

After caching `content.index.json` in memory (guarded by `sync.RWMutex` with `ModTime`/`Size` invalidation for local dev), indexing slugs in a `map[string]IndexPost`, precomputing static server chrome (`menu`, `footerHTML`, `cssVersion`, `versionInfo`, `indexBlurb`) in `server.NewWithOptions()`, and hoisting regex compilation out of hot paths:

| Benchmark | Before | After | Speedup |
|-----------|--------|-------|---------|
| Realistic User Traffic Simulation (500 reqs) | 1.27 ms avg (637.8 ms total) | **0.35 ms (348 µs) avg** (174.0 ms total) | **3.6x faster** |
| Legacy URL Resolution (`ResolveOldURL`, 500 reqs) | 4.02 ms avg | **0.81 ms (810 µs) avg** (405.2 ms total) | **5.0x faster** |
| `go test -race ./internal/blog` | 37.8 s | **19.7 s** | **1.9x faster** |

### AWS Lambda Cold vs. Warm Reality & Frontend Impact

- **Cold start (1st isolated request):** In-memory index caching does **not** speed up request #1 on a fresh Lambda sandbox (the first listing/search still parses `content.index.json` once, ~3 ms, and single post reads don't touch the index at all). On isolated cold hits, Lambda container init (~50–150 ms) and network RTT dominate.
- **Warm invocations (5–15 min Lambda keep-alive):** Any second page click, search query, or crawler burst within the warm window hits the cached index and precomputed chrome (`menu`, `footerHTML`, `indexBlurb`, `cssVersion`), realizing the 3.6×–5× backend reduction and lower GC/CPU usage.
- **First-paint win even on a single cold visit:** Replacing CSS `@import "base.style.css"` with parallel `<link>` tags removes a full render-blocking network round-trip (50–150 ms) on cold browser loads, and moving inline JS into cacheable `assets/js/site.js` shrinks every HTML response.
- **Benchmark scope note:** `TestRealisticUserTrafficSimulation` benchmarks `internal/blog` directly rather than full HTTP round-trips; single post pages still render Markdown on demand via Goldmark to keep memory well under the `64M` container budget.

