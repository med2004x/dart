# Lesson 2: Buffer Pool

## Objective

Implement a buffer pool — the database's own in-memory cache of disk pages — with correct pin/unpin reference counting and eviction, connecting directly to computer architecture Lesson 4's cache concepts applied at the page-management layer instead of the CPU-cache layer.

## Prerequisites

Lesson 1 (pages — the buffer pool caches exactly the page format designed there), computer architecture Lesson 4 (caching principles — a buffer pool is conceptually a software-managed cache, applying the same locality/eviction ideas at a much larger granularity and with the database, not hardware, in control of policy).

## Learn

**Why a database manages its own cache rather than relying purely on the OS page cache (OS Lesson 4).** The OS page cache is real and does help (a page read from disk that's already OS-cached avoids the actual disk I/O, per OS Lesson 4's warm-cache experiment) — but the database has application-level knowledge the OS doesn't: which pages are part of an active transaction, which pages must never be evicted while dirty and unflushed (Lesson 6's WAL/durability requirements), and which access patterns are predictable in advance (e.g. a sequential table scan can benefit from explicit prefetching the OS wouldn't know to do). A buffer pool is the database's own, purpose-built cache layer sitting above the OS's general-purpose one, trading some redundancy (data can genuinely be cached in both layers at once) for control the OS cache can't provide.

**Pin/unpin: reference counting to prevent evicting pages in active use.** Before a transaction reads or modifies a page, it "pins" it (increments a reference count on that buffer pool entry); when done, it "unpins" (decrements). A page with a nonzero pin count must never be evicted, no matter what the eviction policy would otherwise choose — evicting a page currently being read or modified by in-flight code would be a correctness bug, not just a performance one (the code holding a pointer/reference to that buffer pool slot would suddenly be reading whatever new page got loaded into it). This is the buffer pool's own version of the ownership discipline from C track Lesson 4 — every pin needs a matching unpin, and a "pin leak" (forgetting to unpin) permanently prevents that buffer slot from ever being reclaimed, degrading available cache capacity over time.

**Eviction policy.** When the buffer pool is full and a new page needs to be loaded, some currently-unpinned page must be evicted to make room. LRU (Least Recently Used — evict the page that hasn't been accessed in the longest time) is the conceptually simplest policy and a reasonable default, though real systems (including PostgreSQL) often use variants like clock-sweep (an approximation of LRU that's cheaper to maintain exactly, avoiding the overhead of updating a precise recency ordering on every single access) — worth knowing the tradeoff exists even if you implement plain LRU for this lesson.

**Dirty pages.** A page modified in memory but not yet written back to disk is "dirty." Evicting a dirty page requires writing it back to disk first (a "flush") — evicting a clean page requires no such write, since the disk copy is still accurate. This distinction connects directly to Lesson 6: a dirty page's flush timing interacts with write-ahead logging guarantees, since the WAL record describing a change must reach durable storage *before* the corresponding dirty page is allowed to be flushed (the "write-ahead" in WAL is precisely this ordering constraint), which Lesson 6 covers in full.

## Attempt

1. Implement a buffer pool with a fixed capacity (e.g. 10 page slots) over your Lesson 1 page format: `fetchPage(pageID)` (returns the page, loading it from disk into an available buffer slot if not already cached, and pinning it), `unpinPage(pageID, isDirty)` (decrements the pin count, marking the page dirty if the caller modified it), and internal LRU eviction logic used when `fetchPage` needs a free slot but none is available.

2. Test basic correctness: fetch a page, modify it, unpin it as dirty, fetch a *different* set of pages until eviction is forced, and confirm your buffer pool correctly flushed the dirty page to disk before evicting it (verify by reading the page fresh from disk afterward and confirming the modification persisted).

3. Test pin/unpin correctness under a scenario that forces contention: fill the buffer pool to capacity with pinned pages (never unpinning them), then attempt to fetch one more, new page. Confirm your implementation correctly reports failure (no evictable page available) rather than either crashing or, worse, silently evicting a still-pinned page — the latter would be a serious correctness bug, exactly the kind pin/unpin discipline exists to prevent.

4. Measure eviction behavior directly: run a workload that repeatedly accesses a working set of pages larger than your buffer pool's capacity (e.g. 15 distinct pages against a 10-slot pool, accessed in a repeating cyclical pattern), count the total number of actual disk reads your buffer pool performs, and compare against the theoretical minimum if the buffer pool had unlimited capacity (which would need to read each of the 15 pages from disk only once). Report the actual overhead this "working set exceeds cache capacity" scenario causes — directly echoing computer architecture Lesson 4's cache-thrashing discussion, now at the page-management layer.

## Verify

For step 2, show the actual bytes/values you wrote into the page, confirm the buffer pool correctly evicted and flushed it, and confirm reading the page fresh from disk after eviction shows the modified value, not the stale original — proving the flush-before-evict logic is genuinely correct, not just assumed.

## Failure drill

Deliberately introduce a pin leak: modify your test code to call `fetchPage` without ever calling the matching `unpinPage`, repeated for several distinct pages until you've pinned more pages than your buffer pool has slots for. Confirm the buffer pool now cannot service a fetch for yet another new page — every slot is permanently occupied by a leaked pin, exactly mirroring C track Lesson 4's memory-leak failure drill, except the leaked resource here is buffer pool capacity instead of heap memory. Explain in your own words why this failure mode is structurally identical to a memory leak — a resource acquired but never released, silently degrading available capacity — even though buffer pool pins and `malloc`/`free` are different mechanisms entirely.

## Transfer

If TARDOC's PostgreSQL database (covered practically in the postgresql-engineering track) has a `shared_buffers` configuration setting, describe, using this lesson's buffer pool implementation as the concrete model, what that setting actually controls (PostgreSQL's own buffer pool capacity, directly analogous to your `fixed capacity` parameter in step 1), and reason about what would happen — using your step 4 working-set-exceeds-capacity measurement as the basis — if TARDOC's actual working set of frequently-accessed pages (e.g. active clinic records, current billing period data) grew to exceed whatever `shared_buffers` is currently configured to hold.

## Done when

Your buffer pool correctly implements pin/unpin reference counting, LRU eviction, and dirty-page flush-before-evict, verified by actually reading modified data back from disk after eviction rather than just trusting the logic, and you've deliberately caused and explained a pin-leak failure, connecting it explicitly to the same resource-leak pattern covered for heap memory in the C track.
