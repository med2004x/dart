# Lesson 4: Caches

## Objective

Explain the memory hierarchy, why caches exist, and predict when code will be cache-friendly or cache-hostile. Measure the effect directly.

## Prerequisites

Lesson 3 (datapath — the memory stage is where caches sit).

## Learn

**The gap.** A CPU core can execute an instruction in under 1 nanosecond. Main memory (DRAM) access latency is 50-100+ nanoseconds. If every instruction touching memory had to wait for DRAM, the CPU would be idle almost all the time. The memory hierarchy exists to hide this gap: small, fast memory (registers, then L1/L2/L3 cache, in that order of size and speed) sits between the CPU and DRAM, holding recently/frequently used data.

**Locality is why caching works at all.** Two empirical properties of real programs make caching effective:
- *Temporal locality*: if you accessed an address recently, you'll likely access it again soon (a loop variable, a hot function's local state).
- *Spatial locality*: if you accessed an address, you'll likely access nearby addresses soon (the next element in an array, the next field in a struct).

Caches exploit spatial locality directly by fetching a whole **cache line** (typically 64 bytes) on a miss, not just the single byte or word requested. This is why iterating an array sequentially is fast and iterating with a large stride, or chasing pointers scattered across memory (a linked list, a hash map with poor locality), is slow — you pay a full miss penalty per element instead of amortizing one fetch over many nearby elements.

**Cache hierarchy.** L1 (per-core, ~32-64KB, ~4 cycles), L2 (per-core or shared, ~256KB-1MB, ~12 cycles), L3 (shared across cores, several MB, ~40 cycles), DRAM (~200+ cycles). Each level is larger and slower than the one before it.

**Hit, miss, eviction.** A cache hit means the data is already in cache — fast. A miss means it isn't, and must be fetched from the next level down — slow, and it also evicts something else from the cache (following a replacement policy like LRU) to make room. A working set larger than the cache causes **thrashing**: constant eviction and refetching, which can make an algorithm with better asymptotic complexity actually run slower in practice than a "worse" one that fits in cache.

**Why this matters for you directly:** this is the concrete reason "row-major vs column-major matrix traversal" or "array of structs vs struct of arrays" changes real-world performance by 2-10x with zero change to the Big-O complexity. It's also why a hash map (pointer-chasing, poor locality) is often slower in practice than a linear scan over a small sorted array, despite worse asymptotic lookup complexity.

## Attempt

1. Write a benchmark (Go or C) that sums a large 2D array (e.g. 4096x4096 of int32) twice: once iterating row-major (`for i { for j { sum += a[i][j] } }`) and once column-major (`for j { for i { sum += a[i][j] } }`). Time both. In Go, use `testing.B` or manual `time.Now()` deltas with enough iterations to be stable; in C, use `clock()` or a high-resolution timer.

2. Predict before running: which will be faster, and by roughly how much? Then run it and record actual numbers.

3. Write a second benchmark: sum an array by iterating sequentially (stride 1) versus iterating with a stride that's a multiple of the cache line size divided by element size (e.g. stride 16 for int32 with 64-byte lines) over the same total element count, so both touch the same number of elements but the strided version touches far more distinct cache lines relative to elements used.

4. If your platform has `perf` (Linux) available, run `perf stat -e cache-misses,cache-references ./your_binary` on both versions of step 1 and record actual cache miss counts, not just timing.

## Verify

Produce a table: benchmark variant, wall-clock time, (if available) cache-miss count, and a one-sentence explanation tying the timing difference to spatial locality. The row-major version should be measurably faster (often 2x or more depending on array size and cache size) — if it isn't, your array is small enough to fit entirely in cache regardless of order, and you should increase the size until the effect appears.

## Failure drill

Shrink the array size until row-major and column-major run at the same speed. Find (approximately, by binary search on size) the array size where the effect disappears, and explain why in terms of the array now fitting inside L2 or L3 cache regardless of access order — locality only matters when the working set exceeds cache capacity.

## Transfer

Look at one real data structure you use in TARDOC, Mahall, or Lead Sourcer that involves either a slice of structs or a map with non-trivial value types. State, without necessarily rewriting it, whether its access pattern in the hot path is closer to sequential (cache-friendly) or pointer-chasing (cache-hostile), and what you'd change if a profiler showed this location as a bottleneck.

## Done when

You can explain why cache line size makes spatial locality matter, you have measured (not just cited) a real timing difference between a cache-friendly and cache-hostile access pattern on your own machine, and you can identify which access pattern a piece of your own code uses.
