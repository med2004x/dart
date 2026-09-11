# Lesson 7: Storage and I/O Performance

## Objective

Explain the latency and throughput characteristics of storage media (HDD, SSD, network), and reason quantitatively about I/O-bound versus CPU-bound workloads.

## Prerequisites

Lesson 4 (memory hierarchy — storage is the next, much slower, level below DRAM).

## Learn

**Extending the hierarchy.** Lesson 4 covered registers through DRAM, roughly 1-200 cycles of latency. Storage sits below DRAM and is orders of magnitude slower again: an SSD random read is roughly 20,000-100,000 nanoseconds (tens of microseconds), and a spinning HDD random read is 5,000,000-10,000,000 nanoseconds (milliseconds) due to physical seek time and rotational latency. A network round trip within a datacenter is roughly 500,000+ nanoseconds; cross-region can be tens of milliseconds. Keeping these orders of magnitude straight (roughly: register ~0.3ns, L1 ~1ns, DRAM ~100ns, SSD ~20-100µs, HDD seek ~5-10ms, network ~0.5-100ms) is the single most useful intuition for reasoning about where time actually goes in a real system — this is the same table Jeff Dean's "Numbers Every Programmer Should Know" popularized, and it's worth memorizing the order of magnitude even if exact values drift with hardware generations.

**Sequential vs random access.** HDDs are catastrophically slower for random access than sequential, because random access incurs a physical seek (moving the read head) and rotational delay for every access, while sequential access amortizes that cost over a long contiguous read. SSDs have no moving parts, so the gap between sequential and random is much smaller, though still present (controller and flash-block effects). This is why database engines historically optimized heavily for sequential I/O patterns (write-ahead logs, log-structured merge trees) even on SSD-era hardware — old habits that turned out to still matter, just for different underlying reasons (flash erase-block granularity, wear leveling).

**Latency vs throughput vs IOPS.** Latency is time for one operation to complete. Throughput is bytes moved per second, sustained. IOPS (I/O operations per second) matters most for workloads dominated by many small operations (e.g. a database doing many small random reads) rather than a few large ones (e.g. streaming a large file) — a device can have excellent throughput and mediocre IOPS or vice versa, and which one matters depends entirely on your access pattern.

**Queueing and concurrency.** A single storage device has limited parallelism. Issuing many I/O requests concurrently (instead of one at a time, waiting for each to complete before issuing the next) can dramatically increase achieved throughput on SSDs, because the device can service multiple requests in flight — this is why async I/O and connection pooling to disk-backed services matter, and it connects directly to the queueing theory in the math-for-engineering track (Lesson 08 step-06).

**Why this matters for you directly:** TARDOC and Mahall both do database I/O and network I/O as their dominant cost centers, not CPU. Understanding that a database round trip costs roughly 1000x a memory access, and a network call to another service costs another order of magnitude beyond that, is what justifies patterns like batching queries, connection pooling, caching frequently-read rows in Redis, and avoiding N+1 query patterns — all things you've already had to fix in real projects, now with the underlying "why" made explicit.

## Attempt

1. Write a benchmark (Go, using `os.File`, or C) that writes and then reads back 1GB of data two ways: (a) sequentially in large chunks (e.g. 1MB writes), and (b) randomly at 4KB granularity to random offsets within a pre-allocated 1GB file. Time both the write and read phases separately. Run this on whatever disk you actually have (note whether it's SSD or HDD, since the gap size depends on this).

2. Predict before running which will be slower and by roughly what order of magnitude, then compare to your actual results.

3. Query your own TARDOC or Mahall database for one endpoint you know is used often (e.g. clinic lookup, or a product listing query) and use `EXPLAIN ANALYZE` (PostgreSQL) to see the actual execution time and whether it's using an index (fast, roughly logarithmic random access) or a sequential scan (linear, potentially slow on a large table). Record the actual numbers.

4. Using the latency table from Learn, estimate (order of magnitude, not exact) how many DRAM accesses "fit" in the time of one SSD random read, and how many SSD random reads "fit" in the time of one cross-region network round trip. This is meant to build intuition, not precision.

## Verify

Report actual measured numbers for step 1 (sequential vs random throughput/latency on your real disk) and step 3 (real query time and whether an index was used), not estimates.

## Failure drill

In step 3, if the query is already using an index, deliberately write a query on the same table that forces a sequential scan (e.g. a condition on a non-indexed column, or wrapping the indexed column in a function that defeats index usage) and compare `EXPLAIN ANALYZE` timing. This makes the cost of losing index usage concrete rather than theoretical.

## Transfer

Identify one place in TARDOC, Mahall, or Lead Sourcer where you already added caching, batching, or connection pooling (or where you should). State explicitly, using the latency numbers from this lesson, what order-of-magnitude improvement that change is actually buying you, and whether the workload is closer to latency-bound (few large ops) or IOPS-bound (many small ops).

## Done when

You can recite the rough order-of-magnitude latency numbers for register, DRAM, SSD, HDD seek, and network round trip without needing to look them up, you've measured sequential vs random I/O performance on real hardware, and you can point to one real query in your own systems and explain its cost in terms of this hierarchy.
