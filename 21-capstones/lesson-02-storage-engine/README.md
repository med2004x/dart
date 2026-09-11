# Capstone 2: Embedded Storage Engine

## Objective

Build a genuinely complete, persistent embedded storage engine — pages, indexing, and crash recovery integrated into one coherent system — the second terminal capstone, focused specifically on systems and database-internals reasoning at a scale beyond that track's own Lesson 7 capstone.

## Prerequisites

Database-internals track (all 7 lessons), computer architecture (caching, storage performance), C or Rust track (for the actual implementation — this capstone is deliberately suited to either language, given its systems-level nature).

## Learn

There is no new material. This capstone differs from database-internals Lesson 7 (which explicitly offered BusTub as an equally-valid alternative path) by requiring the full system built and reasoned about end to end by you, specifically to surface the gaps between "I understand pages, buffer pools, B+ trees, and WAL as separate, individually-tested concepts" and "I can build one coherent, correct system from them" — the same integration-testing philosophy as Capstone 1, applied here to a systems-programming artifact rather than a network service.

**What "embedded" specifically means, and why it's the right scope for this capstone.** An embedded storage engine (like SQLite, or the storage layer inside a larger database) runs in-process, with no network layer, no client/server separation, no concurrent multi-process access to reason about — this deliberately narrows the scope compared to Capstone 1 or 3, letting this capstone go *deeper* on storage-specific correctness (crash recovery, index correctness, exact durability guarantees) without also having to solve networking or distributed-systems problems simultaneously.

## Attempt

Build a complete embedded storage engine with this minimum integrated scope, in C or Rust:

1. **Page-based storage** (database-internals Lesson 1): a real, on-disk slotted-page format, with measured storage overhead reported honestly.
2. **Buffer pool management** (database-internals Lesson 2): correct pin/unpin discipline, LRU eviction, dirty-page flush-before-evict — verified, not assumed, by actually reading modified data back from disk after a forced eviction.
3. **B+ tree indexing** (database-internals Lesson 3): real search, insert-with-splitting, and range queries via leaf-chain traversal, with measured node-visit counts confirming logarithmic-ish lookup behavior at real data scale.
4. **A query execution layer** (database-internals Lesson 4): the iterator pattern (scan, filter, project), with both sequential and index-scan paths, and a real, measured comparison of when each wins.
5. **Write-ahead logging and crash recovery** (database-internals Lesson 6): every mutation logged before applied, a working `recover()` that reconstructs correct state after a simulated crash of the *entire integrated system* — not the isolated WAL exercise, the whole storage engine.
6. **A minimal public API**: enough of a client interface (even a simple `put(key, value)`/`get(key)`/`scan(range)` surface) that the engine could plausibly be embedded in an application, with that application's perspective driving your API design (API engineering Lesson 1's contract-first discipline, applied to a library API instead of an HTTP one).

If working in Rust: apply Rust track Lessons 1-3 deliberately — model your page/buffer-pool ownership using Rust's type system rather than manual discipline, and use enums for any state machine (buffer slot state, transaction state) rather than loose integer flags.

## Verify

1. **Crash-recovery correctness under real load**: run a substantial, realistic sequence of operations (thousands, not dozens), simulate a crash at a genuinely random point (not one you chose for convenience), and confirm `recover()` reconstructs state that exactly matches a crash-free reference execution.
2. **Benchmark reads and writes** (performance track Lesson 1's rigor): report real numbers with variance, for both sequential and random access patterns, and explain the results using computer architecture Lesson 4/7's cache and storage-latency principles — your numbers should be *explicable*, not just reported.
3. **Index versus scan comparison**: real measured numbers, at both a selective and non-selective query, per database-internals Lesson 4's own capstone requirement, now integrated into the full system rather than tested in isolation.

## Failure drill

Take your crash-recovery mechanism and specifically attack the *interaction* between buffer pool eviction and WAL ordering — construct a scenario where a page is evicted (and flushed) at nearly the same moment its corresponding WAL record is being written, and confirm your system correctly maintains the write-ahead ordering constraint (database-internals Lesson 6) even under this timing pressure, not just in the lesson's simpler isolated test. If you find a genuine ordering bug at this integration point, document it, fix it, and explain specifically why it wasn't visible when buffer pool and WAL were each tested independently.

## Transfer

Compare your engine's actual measured read/write performance and crash-recovery guarantees against SQLite's documented behavior (a real, production embedded engine solving the same core problem) — identify at least 2 specific things SQLite does that yours doesn't, and reason about what real-world requirement (multi-process concurrent access, a broader SQL surface, decades of edge-case hardening) drove that additional complexity beyond your capstone's necessarily narrower scope.

## Done when

Your storage engine correctly survives a crash-recovery test under real, substantial load with a randomly-timed simulated crash, you have real, explicable benchmark numbers for both access patterns, you've found and fixed (or confirmed the absence of) a buffer-pool/WAL-ordering integration bug specifically, and you've honestly compared your engine's guarantees against a real production system's documented behavior.
