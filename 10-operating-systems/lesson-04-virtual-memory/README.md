# Lesson 4: Virtual Memory (OS Perspective)

## Objective

Go deeper than computer architecture Lesson 6's MMU/TLB mechanics into the OS's role specifically: managing page tables per process, handling page faults as software (not just hardware) events, and inspecting real memory mappings on a running Linux process.

## Prerequisites

Computer architecture Lesson 6 (virtual memory — this lesson assumes that lesson's hardware-level address translation as given and focuses on what the OS does around it).

## Learn

**Division of labor: hardware vs. OS.** The MMU (hardware, computer architecture Lesson 6) performs address translation using the page table — but the OS is what *creates and maintains* that page table, decides the initial memory layout of a new process, and handles what happens when translation fails (a page fault traps into OS code, not hardware logic, precisely because deciding "load from disk" vs. "kill the process" vs. "allocate a fresh zero page" requires OS-level context the MMU doesn't have).

**Why a fresh process's memory isn't actually all mapped from the start.** When a process starts, its executable's code and initial data are mapped, but many pages it will eventually use (heap growth via `malloc`→`brk`/`mmap`, stack growth) are *not* physically backed until first touched — this is demand paging, previewed in computer architecture Lesson 6's `mmap`/`VmRSS` experiment, and it's the OS's page-fault handler that does the actual work of mapping a fresh physical page in on first access, not something that happens automatically without OS involvement.

**`/proc/<pid>/maps`: the OS's own record of a process's virtual address space.** This file (readable on any Linux system) lists every mapped memory region for a running process — code, data, heap, stack, shared libraries, memory-mapped files — with their address ranges, permissions (r/w/x), and what backs them (a file, or anonymous memory). This is the direct, inspectable evidence of the abstract "virtual address space" concept: a real table you can read for a real running process.

**Page fault types, from the OS's perspective, not just the CPU-trap mechanism.** A **minor fault** occurs when the page exists in memory (perhaps already mapped by another process sharing the same file, or already in the page cache) but isn't yet mapped into *this* process's page table — resolved quickly, no disk I/O needed. A **major fault** requires reading from disk (the page genuinely isn't in RAM anywhere yet) — orders of magnitude slower, directly connecting to computer architecture Lesson 7's storage latency numbers. `/proc/<pid>/stat` and tools like `ps -o min_flt,maj_flt` expose these counts per process, letting you distinguish "lots of faults but all cheap" from "faults that are actually costing real time."

## Attempt

1. Run `cat /proc/self/maps` (or `/proc/<pid>/maps` for any running process) and identify at least: the executable's code region (permissions should show `r-xp`, executable but not writable — direct evidence of the write-protection on code mentioned in computer architecture Lesson 6), the heap region, the stack region, and at least one shared library mapping. Note their address ranges and confirm the stack is typically at a much higher address than the heap, with a large unmapped gap between them (this gap is intentional — room for both to grow toward each other without pre-committing all that address space).

2. Write a program that allocates a large amount of memory (e.g. 200MB via `malloc` in C or a large slice in Go) but only touches a small portion of it, then re-run `cat /proc/<pid>/maps` while it's running (you'll need to pause the program, e.g. with a `sleep()` call or a breakpoint, to inspect it mid-execution) and confirm the large allocation appears as one mapped region in `/proc/<pid>/maps` — the *virtual* mapping exists as one contiguous entry — even though (per computer architecture Lesson 6's `VmRSS` experiment) most of it isn't physically resident yet.

3. Compare minor vs. major fault counts for two different programs: one that touches a large amount of freshly allocated (but never-before-accessed) memory (expect mostly minor faults, since it's anonymous/zero-fill memory, not something requiring a disk read), and one that reads a large file it hasn't accessed before, that isn't already in the OS's page cache (harder to guarantee reliably, but try reading a large, infrequently-accessed file — expect to see at least some major faults corresponding to actual disk reads). Use `ps -o min_flt,maj_flt -p <pid>` before and after each program runs, or `/usr/bin/time -v` on Linux which reports these directly.

4. Trigger and observe an invalid access (a real page fault the OS resolves by killing the process, not by mapping a page) — reference computer architecture Lesson 6's null-pointer-dereference example — and this time, in addition to observing the segfault, use `dmesg` (or `journalctl -k` on systemd systems) immediately after to find the kernel's own log entry for the fault, which typically records the faulting address and the process that triggered it — this is the OS's own record of the event, distinct from what your program's own crash output shows.

## Verify

For step 3, report the actual minor and major fault counts (not just "some faults happened") for both test programs, and state explicitly which one showed meaningfully more major faults, connecting the result to whether the underlying data was fresh anonymous memory versus disk-backed file content not yet cached.

## Failure drill

Rerun step 3's file-reading test a second time, immediately after the first run, without doing anything to clear caches in between. Observe the major fault count drop sharply (likely to near zero) on the second run. Explain why: the OS's page cache now holds the file's contents in RAM from the first read, so the second read's page faults are resolved as *minor* faults (the data's already in memory, just needs mapping into this process) rather than major faults requiring an actual disk read — this is the same mechanism that makes "warm cache" reads of recently-accessed files dramatically faster than "cold" first reads, and it's a real, common source of confusing benchmark results when a first run and a repeated run of the same test produce very different timings for reasons that have nothing to do with the code being tested.

## Transfer

If TARDOC's Celery workers read the same reference data files repeatedly (config, model files, or lookup tables used across many transcription jobs), state what you'd expect page-cache behavior to do for you automatically — the OS caching frequently-read file pages in RAM across process invocations, without your application code doing any explicit caching itself — and what would defeat this benefit (e.g. reading from a very large number of distinct files that don't fit in available page cache, or running each job in an environment, like certain containerized deployments, where the cache doesn't persist between invocations the way it would on a long-running bare process).

## Done when

You can read and interpret `/proc/<pid>/maps` for a real running process, you've measured a concrete difference in minor vs. major fault counts between memory-only and disk-backed access patterns, and you've directly observed page-cache warming reduce major faults on a repeated read, connecting it to real-world caching behavior you can now recognize rather than treat as a black box.
