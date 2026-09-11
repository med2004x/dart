# Lesson 5: Profiling and System Observation

## Objective

Measure before optimizing — use `strace` to observe what syscalls a program actually makes, and a CPU profiler to find where time is actually spent, replacing guesswork with evidence before making any performance change.

## Prerequisites

Linux Lesson 1 (file descriptors and syscalls conceptually), computer architecture Lesson 7 (storage/IO latency numbers — profiling output makes far more sense once you have the rough latency table from that lesson in mind).

## Learn

**Why "measure before optimizing" is the single rule that matters most here.** Intuition about where a program spends its time is frequently wrong — a function that looks expensive by inspection (nested loops, complex logic) may be negligible next to one unglamorous line doing a blocking network call. Optimizing the wrong thing wastes effort and sometimes makes code worse (more complex) for zero measured benefit. Every profiling tool in this lesson exists to replace "I think it's slow because of X" with "the data shows it's actually slow because of Y."

**`strace`: observing every syscall a process makes.** `strace -c ./program` gives a summary (which syscalls were called, how many times, total time spent in each) — often immediately revealing surprising patterns, like a program calling `read()` thousands of times in tiny chunks instead of a few large reads, or repeatedly `stat()`-ing the same file. `strace -T` shows per-call timing. This tool operates entirely at the syscall boundary — it can't see what happens inside your program's own code between syscalls, only every point where your program asks the kernel to do something (read a file, write to a socket, allocate memory via `mmap`/`brk`).

**CPU profiling: sampling where time is spent inside your own code.** Unlike `strace` (syscall boundary), a CPU profiler (Go's built-in `pprof`, or `perf` on Linux generally) periodically samples the program's call stack while it runs, building up a statistical picture of which functions were "on the stack" (actively executing or a caller of what's executing) most often — this becomes a good approximation of where CPU time is actually going, without needing to instrument every function manually.

**Correlating a hypothesis with actual measurement.** The discipline this lesson is really teaching: before profiling, write down your guess for where the bottleneck is. After profiling, compare. Being wrong is the common, expected, useful outcome — it's direct evidence that intuition alone isn't reliable enough to skip measurement, which is the entire justification for treating profiling as a required step rather than an optional one you reach for only when confused.

## Attempt

1. Write a small program that does a mix of CPU-bound work (e.g. a nested loop computing something) and I/O (reading a file, or making a small number of syscalls). Before running anything, write down your prediction: which part do you expect dominates the total runtime?

2. Run `strace -c` on the program and record the syscall summary — total time in syscalls, and which specific syscalls dominate.

3. Profile the CPU-bound portion using a language-appropriate profiler: in Go, `go test -cpuprofile=cpu.out -bench=.` (or `pprof.StartCPUProfile` around the relevant code) followed by `go tool pprof cpu.out`, then use `top` inside pprof's interactive mode to see which functions consumed the most sampled time. In C, `perf record ./program` followed by `perf report`.

4. Compare your step 1 prediction to the actual measured breakdown from steps 2-3. State explicitly whether you were right, and if not, what specifically about the actual bottleneck surprised you relative to your intuition.

## Verify

Produce the actual `strace -c` summary table and the actual profiler's top-N function list side by side, and state in one sentence which one (syscall time or CPU-bound function time) turned out to dominate total runtime for your specific program.

## Failure drill

Deliberately profile a program dominated by I/O wait (e.g. one that mostly sleeps or blocks on a slow read) using only a CPU profiler, with no `strace` or I/O-aware tooling. Observe that the CPU profiler shows very little "hot" code — because the program spends most of its wall-clock time blocked, not executing instructions, and a CPU sampling profiler by design only samples when the process is actually running on a CPU core. Explain why this is not a bug in the profiler but a fundamental scope limitation: a CPU profiler answers "where is CPU time going" specifically, not "where is wall-clock time going" — for an I/O-bound program these are very different questions, and conflating them leads to concluding "there's nothing to optimize" about a program that's actually spending 95% of its time waiting on a slow disk or network call, which strace (or a wall-clock/tracing profiler) would have revealed clearly.

## Transfer

If TARDOC's transcription pipeline or Celery task processing has ever felt slow in a way you addressed by guessing at the cause (e.g. assuming it was CPU-bound transcription work) rather than profiling first, describe what a real `strace -c` or Go pprof run against that specific workload would tell you that your assumption couldn't — and whether, given the multi-key Groq API rotation mentioned in your project history, you'd now predict the bottleneck is more likely CPU-bound local processing or I/O-bound waiting on external API calls, and how you'd confirm which one it actually is rather than continuing to assume.

## Done when

You've measured a real program with both `strace -c` and a CPU profiler and can state which resource (syscall/IO time vs CPU time) actually dominated, you've compared that measurement against a prediction you wrote down beforehand and can honestly report whether you were right, and you can explain — using the failure drill's I/O-bound case as the concrete example — why a CPU profiler alone can be actively misleading for programs that spend most of their time blocked rather than computing.
