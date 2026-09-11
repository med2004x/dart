# Lesson 3: Allocation and Memory Behavior

## Objective

Understand Go's escape analysis (what determines whether a value is stack- or heap-allocated), measure allocations-per-operation directly, and reduce allocation count in a real function without changing its observable behavior — connecting directly back to C track Lesson 2's stack/heap distinction, now viewed through Go's compiler-automated lens.

## Prerequisites

C track Lesson 2 (memory model — Go's escape analysis automates exactly the stack-vs-heap decision that lesson covered manually for C), Lesson 2 of this track (profiling — this lesson uses the same allocation-profiling tools, now focused specifically on understanding *why* allocations happen, not just where).

## Learn

**Escape analysis, precisely.** C track Lesson 2 established that returning a pointer to a local (stack) variable is undefined behavior in C, because the stack frame is reclaimed on function return. Go's compiler solves this differently: it performs escape analysis at compile time, determining whether a value's lifetime could extend beyond its allocating function's stack frame (e.g. because a pointer to it is returned, stored in a longer-lived structure, or captured by a closure that outlives the function) — if so, the compiler allocates it on the heap instead of the stack, automatically and transparently, specifically to make the C-track bug structurally impossible in Go. `go build -gcflags="-m"` reports these decisions explicitly, showing "escapes to heap" or "moved to heap" for each value the compiler had to make this determination for.

**Why heap allocations aren't free, even with automatic memory management.** Stack allocation is essentially free (a pointer bump, per C track Lesson 2's mechanics). Heap allocation requires the allocator to find and reserve space, and — critically for Go specifically — every heap-allocated object becomes something the garbage collector must eventually track and potentially scan during a collection cycle. A program that allocates heavily, even if each individual allocation is small, creates ongoing GC pressure: more frequent collection cycles, and CPU time spent during those cycles that isn't available for your actual program logic — this is precisely why Lesson 2's allocation-profile comparison showed a real CPU-time improvement from an allocation-reducing fix, not just a memory-usage improvement.

**Reducing allocations without changing observable behavior — the actual skill this lesson builds.** Common, real techniques: pre-allocating a slice with a known or estimated capacity (`make([]T, 0, expectedSize)`) instead of letting repeated `append` calls trigger multiple reallocations as the slice grows; reusing buffers across calls (e.g. via `sync.Pool`) instead of allocating fresh ones each time; avoiding unnecessary pointer indirection when a value type would do (a pointer to a small struct often forces heap allocation via escape analysis, where the value type itself might not need to). Each of these changes the *implementation* without changing what the function actually computes or returns — the discipline this track has emphasized throughout: verify correctness is preserved, not just that performance improved.

**Why premature allocation-optimization is a real anti-pattern too.** Not every allocation matters — a function called once per user request, allocating a modest amount, is rarely worth obsessing over compared to a function called millions of times in a hot inner loop. Lesson 2's profiling discipline is what tells you which allocations are actually worth addressing (the ones showing up as significant in a real profile of a real, representative workload) versus which would be optimization effort spent on something that was never actually a bottleneck.

## Attempt

1. Write a small Go function and compile it with `go build -gcflags="-m"`, examining the output for "escapes to heap" messages. Identify at least one value that escapes (e.g. because it's returned as a pointer) and one that doesn't (a purely local value never referenced outside the function), confirming your understanding of what triggers escape by predicting each outcome before checking the compiler's actual determination.

2. Write a function that repeatedly appends to a slice without pre-allocating capacity (`var s []int; for i := 0; i < N; i++ { s = append(s, i) }`), and benchmark it (Lesson 1's methodology) alongside a version that pre-allocates capacity (`s := make([]int, 0, N)`). Use an allocation profile (Lesson 2) to compare allocations-per-operation between the two versions, and report the actual difference.

3. Implement a `sync.Pool`-based buffer reuse pattern for a function that would otherwise allocate a fresh buffer on every call (e.g. a function building up a string or byte slice repeatedly in a hot loop simulating many requests), and benchmark the before/after allocation counts and timing, confirming the pool-based version shows measurably fewer allocations under repeated calls.

4. Take a struct that's currently passed and returned by pointer (`*MyStruct`) where the struct is small and doesn't need to outlive the function scope, and change it to be passed/returned by value instead. Use `-gcflags="-m"` to confirm the value version avoids the heap escape the pointer version triggered, and benchmark both to measure the actual performance difference.

## Verify

Present your `-gcflags="-m"` output for step 1's escaping and non-escaping values with your correct predictions, your step 2 allocation-profile comparison (pre-allocated vs. not) with actual numbers, and your step 3/4 before/after benchmark and allocation results.

## Failure drill

Take your step 2 pre-allocation fix and deliberately pre-allocate with a capacity *smaller* than what's actually needed (e.g. `make([]int, 0, 10)` when you actually append 10,000 items). Benchmark this "under-allocated" version against both the original (no pre-allocation) and the correctly-sized pre-allocation versions. Confirm the under-allocated version's allocation count is much closer to the no-pre-allocation baseline than to the correctly-sized version — because Go's slice growth still triggers reallocations once the initial small capacity is exceeded, just delayed slightly compared to starting from zero capacity. Explain why this demonstrates that pre-allocation's benefit specifically depends on reasonably accurately estimating the actual needed capacity — an under-sized guess provides much less benefit than expected, a genuinely easy mistake to make if the "expected size" estimate itself wasn't grounded in real data about typical usage.

## Transfer

If TARDOC's transcription result processing or Lead Sourcer's per-lead scoring logic builds up strings or slices in a loop (a plausible pattern given text/data processing at the core of both), describe what you'd check first using this lesson's tools (`-gcflags="-m"` for escape analysis, an allocation profile per Lesson 2) before attempting any allocation-reduction optimization, and state explicitly why checking first (rather than optimizing based on a general "pre-allocating is usually good practice" assumption) matters, per this lesson's premature-optimization caution.

## Done when

You can correctly predict and verify Go's escape analysis decisions for a new function using `-gcflags="-m"`, you've measured a real allocation-count and timing improvement from proper slice pre-allocation and from buffer reuse via `sync.Pool`, and you've demonstrated — via the failure drill — that pre-allocation's benefit depends on a reasonably accurate capacity estimate, not just the presence of pre-allocation itself.
