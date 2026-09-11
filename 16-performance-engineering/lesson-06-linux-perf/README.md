# Lesson 6: Linux Performance Investigation

## Objective

Use system-level tools (beyond Go-specific profiling) to diagnose CPU, memory, disk, and network bottlenecks, and correlate system-level evidence with application-level symptoms to produce a real, evidence-based incident diagnosis — the capstone-style integration of this track's measurement discipline applied to whole-system investigation.

## Prerequisites

Linux tools Lesson 5 (strace/profiling basics), computer architecture Lessons 4/7 (cache and storage latency numbers, needed to interpret system-level measurements meaningfully), Lessons 1-5 of this track (the measurement rigor and Go-specific tools this lesson now combines with system-wide tooling).

## Learn

**Why application-level profiling (Lesson 2) alone sometimes isn't enough.** A Go CPU profile shows where *your program's* CPU time goes, but says nothing about whether the *machine itself* is CPU-saturated by other processes, swapping due to memory pressure (OS Lesson 4), or I/O-bound waiting on a slow disk or network call your Go profiler can't see into (since, per Linux tools Lesson 5's core point, a CPU profiler only samples when your process is actually running on a core, not when it's blocked). System-level tools are what let you distinguish "my code is genuinely CPU-bound and needs optimization" from "my code is fine, but the system it's running on is starved of some resource."

**`top`/`htop`: system-wide CPU and memory at a glance.** Beyond your own process, these show overall CPU utilization (per-core, revealing whether load is balanced or concentrated on one core — relevant if your Go program isn't parallelizing effectively), memory usage including how much is used by the OS page cache (OS Lesson 4) versus actual application memory, and load average (a rough indicator of how many processes are runnable/waiting for CPU time, connecting to OS Lesson 3's scheduling concepts).

**`vmstat`: a compact summary combining CPU, memory, and I/O activity over time.** Particularly useful for spotting swapping (the `si`/`so` columns — if nonzero, the system is actively swapping memory pages to disk, an extremely expensive operation per computer architecture Lesson 7's latency numbers, and a strong signal of genuine memory pressure) and for correlating CPU busy time against I/O wait time (the `wa` column — high `wa` alongside low CPU usage is a strong signal the bottleneck is I/O, not CPU, echoing Linux tools Lesson 5's failure drill about CPU profilers missing I/O-bound bottlenecks entirely).

**`iostat`: disk-specific throughput and latency.** Shows per-device read/write throughput and, critically, `await` (average time for I/O requests to complete) — a rising `await` under load is direct evidence of disk-level bottleneck, connecting concretely to computer architecture Lesson 7's sequential-vs-random and SSD-vs-HDD latency discussion, now observed on real, live infrastructure rather than in a controlled benchmark.

**Correlating system-level and application-level evidence: the actual diagnostic skill.** A real incident diagnosis combines both: application logs/metrics (system-engineering Lesson 9) showing *when* and *what* was slow, and system-level tools (this lesson) showing *why* — e.g., application latency spiked at 2:47pm, and `vmstat`/`iostat` for that same window show elevated `wa` and `await`, pointing to a disk bottleneck as the likely root cause, rather than a bug in the application code itself. This correlation — timestamps lining up between two independent sources of evidence — is what turns a plausible hypothesis into a well-supported diagnosis.

## Attempt

1. On a real Linux machine (a VPS, a local VM, or your development environment), run `top`/`htop`, `vmstat 1` (repeated every second), and `iostat -x 1` simultaneously while your system is at normal, idle-ish load, and record baseline values for each — this baseline is what you'll compare against under induced load in the next steps.

2. Induce a CPU-bound load (e.g. run a tight computational loop across multiple goroutines, or use a stress-testing tool like `stress-ng --cpu N`) and observe the same three tools during the load. Confirm CPU usage rises accordingly in `top`, and note whether `vmstat`'s `wa` column stays low (as expected for a genuinely CPU-bound, not I/O-bound, workload).

3. Induce an I/O-bound load instead (e.g. `stress-ng --io N` or a script performing heavy, uncached disk reads/writes) and observe the same tools. Confirm `iostat`'s `await` rises and `vmstat`'s `wa` column shows elevated I/O wait time, directly distinguishing this scenario's system-level signature from step 2's CPU-bound signature.

4. Induce a memory-pressure scenario (allocate memory approaching or exceeding available RAM, carefully, on a test system you're comfortable risking instability on) and observe `vmstat`'s `si`/`so` (swap in/out) columns becoming nonzero, directly connecting to OS Lesson 4's virtual-memory/demand-paging discussion — confirm the system's overall responsiveness degrades noticeably once swapping becomes significant, giving you a felt, not just measured, sense of why memory pressure is a serious performance problem.

## Verify

Present your baseline and induced-load measurements for all three scenarios (CPU-bound, I/O-bound, memory-pressure) side by side, showing the distinct signature each produces across `top`/`vmstat`/`iostat` — a real reference table you produced yourself, distinguishing what each bottleneck type actually looks like in these tools' output.

## Failure drill

Take a workload that's genuinely I/O-bound (per step 3) and attempt to diagnose it using *only* `top`'s CPU percentage (ignoring `vmstat`'s `wa` column and `iostat` entirely) — confirm that CPU usage alone looks unremarkable or even low during this I/O-bound workload, which, viewed in isolation, could easily lead to an incorrect conclusion that "the system isn't under significant load" when it's actually experiencing a real, measurable I/O bottleneck. Explain why this single-tool, incomplete-evidence mistake is exactly the kind of misdiagnosis this lesson's correlation-across-tools discipline exists to prevent — a genuinely common failure mode when someone reaches for the most familiar tool (`top`) without knowing to check the others for a full picture.

## Transfer

Write a short, evidence-based incident diagnosis for a real or plausible slowdown scenario in TARDOC (e.g. "transcription processing became slow during a specific window") using this lesson's full toolkit: state what you'd check first (application logs/metrics for the affected time window, per system-engineering Lesson 9), what system-level tools you'd run to correlate against that window, and what specific signature in each tool's output would confirm or rule out CPU-bound, I/O-bound, and memory-pressure explanations respectively — a concrete diagnostic plan, not just a general statement that you'd "investigate."

## Done when

You've directly observed and recorded the distinct system-level signatures of CPU-bound, I/O-bound, and memory-pressure scenarios using real tooling on a real machine, you've demonstrated — via the failure drill — why relying on a single tool (`top` alone) can produce an incomplete or misleading picture, and you've written a concrete, evidence-based diagnostic plan for a real scenario that correlates application-level and system-level evidence rather than relying on either alone.
