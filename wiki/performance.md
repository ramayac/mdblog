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

Benchmarks are implemented in [internal/blog/performance_test.go](file:///home/ramayac/git/MDBlog/internal/blog/performance_test.go).

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

