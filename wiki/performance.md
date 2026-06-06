# Performance and Benchmark Insights

This document outlines the performance characteristics, benchmark results, and architectural optimizations of MDBlog.

---

## 1. Executive Summary

MDBlog is a database-free, flat-file blog engine. To keep serving speeds fast without a database, it uses a pre-built metadata index (`content.index.json`) for lists/search, and direct filesystem lookups for single posts. 

Benchmarks show that MDBlog achieves:
* **Average request latency under realistic traffic:** **~1.27 milliseconds**
* **Metadata index JSON parsing speed:** **~2.86 milliseconds**
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
