# Lesson 6: Virtual Memory

## Objective

Explain why virtual memory exists, how address translation works, and what a page fault and a TLB miss actually are.

## Prerequisites

Lesson 4 (caches — the TLB is a cache, and page faults interact with the memory hierarchy).

## Learn

**The problem virtual memory solves.** Without it, every running program would need to know the physical addresses of real RAM, would collide with every other running program's memory, and a buggy program could read or corrupt another program's data or the OS itself. Virtual memory gives every process its own private address space (0 to some max, e.g. 2^47 on x86-64) that the process addresses as if it owned all of memory, while the OS and CPU cooperate to map those virtual addresses to real physical RAM, transparently and per-process.

**Pages.** Memory is divided into fixed-size chunks called pages (commonly 4KB). A **page table** (per process, maintained by the OS) maps virtual page numbers to physical page frame numbers. The CPU's Memory Management Unit (MMU) does this translation in hardware on every memory access.

**Address translation, concretely.** A virtual address splits into a page number (upper bits) and an offset within the page (lower bits). The MMU looks up the page number in the page table to get a physical frame number, then combines that with the unchanged offset to get the physical address. Page tables are typically multi-level (a tree, not a flat array) because a flat table for a 47-bit address space with 4KB pages would itself be enormous.

**The TLB (Translation Lookaside Buffer)** is a small, fast cache (same idea as Lesson 4's caches, applied to translations instead of data) inside the CPU that stores recently used virtual-to-physical mappings, so most memory accesses don't need a full page table walk. A TLB miss forces a page table walk — slower, though still normally serviced from cache-resident page table entries, not from DRAM every time.

**Page faults.** If a virtual address has no valid mapping (page not present — e.g. it's been swapped to disk, or it's simply unallocated/invalid), the MMU triggers a page fault, which traps into the OS. The OS then either loads the page from disk and resumes the process (a "soft" recoverable fault — this is how demand paging and swap work), or kills the process with a segmentation fault if the access was genuinely invalid (e.g. a null pointer dereference, or writing to unmapped memory).

**Why this matters:** this is the actual mechanism behind "segfault," why processes are memory-isolated from each other by default, why `mmap`-based file I/O works, and why touching a huge sparse allocation for the first time (page faults on first touch) is measurably slower than touching already-resident memory — relevant if you've ever seen a mysterious first-access latency spike in a service.

## Attempt

1. On Linux, write a small C or Go program that allocates a large buffer (e.g. 500MB with `mmap` or a large slice in Go) and does nothing with most of it, then touches (writes to) only a small portion. Use `/proc/self/status` (look at `VmRSS` — resident set size) before and after the touch to observe that RSS only grows for pages actually touched, not the full allocation. This demonstrates demand paging directly: virtual address space was reserved, but physical pages are assigned lazily.

2. Deliberately write to a null or invalid pointer in a small C program (`int *p = NULL; *p = 5;`) and run it. Observe the segmentation fault. Explain in your own words what the MMU/OS did: attempted translation failed because there was no valid mapping, triggering a trap the OS resolved by killing the process (unlike the demand-paging case in step 1, where the OS resolved the fault by mapping a page).

3. If you have access to `perf` on Linux: run `perf stat -e dTLB-load-misses` on a program with poor memory locality (reuse the strided-access benchmark from Lesson 4) versus one with good locality, and note the TLB miss count difference — poor locality causes more distinct pages to be touched per unit of useful work, which increases TLB pressure the same way it increases cache pressure.

## Verify

For step 1, produce the actual `VmRSS` numbers before allocation, after allocation (should still be low — pages not yet touched), and after touching the small portion (should increase by roughly the touched size, rounded up to page granularity, not by the full 500MB).

## Failure drill

In step 1, instead of touching a small portion, touch the entire 500MB buffer (write to every page). Observe `VmRSS` now growing to match the full allocation. Explain the difference from the partial-touch case purely in terms of "a page fault occurs the first time each page is written, and the OS maps a physical frame at that point" — there's no other mechanism at play, just more pages being faulted in.

## Transfer

Explain in one paragraph why containers (Docker) can each believe they have access to a large amount of memory on a host with less total physical RAM than the sum of container limits, connecting this to demand paging and the fact that unused virtual address space costs nothing until touched.

## Done when

You can explain address translation (virtual address to page number to physical frame) without notes, you've observed demand paging directly via `VmRSS`, and you can explain the difference between a recoverable page fault (demand paging) and a fatal one (invalid access / segfault).
