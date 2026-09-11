# Lesson 4: Debuggers and Core Dumps

## Objective

Use a real debugger (`gdb` or `lldb`) to inspect a program's actual runtime state — stack frames, variable values, registers — rather than debugging by inserting print statements and re-running, and learn to analyze a core dump from a crash that already happened and can't be reproduced live.

## Prerequisites

Computer architecture Lesson 3 (CPU datapath/registers) and C track Lesson 2 (memory model/stack frames) — a debugger's stack-frame and register inspection is directly reading the concepts covered there.

## Learn

**Why a debugger beats print-statement debugging for many bugs.** Print debugging requires guessing in advance which values matter and recompiling every time you guess wrong. A debugger lets you pause execution at any point and inspect *any* value after the fact, including ones you didn't think to print — genuinely faster once you're fluent with it, though print debugging remains reasonable for simple, well-understood cases. The skill this lesson builds is having the debugger available as a real option, not defaulting to print statements purely out of unfamiliarity with the alternative.

**Breakpoints.** `break <file>:<line>` (gdb) or `b <file>:<line>` (lldb) pauses execution when that line is about to run. `break <function_name>` pauses on entry to a function. Once paused, `next`/`n` steps over a line (executing function calls without entering them), `step`/`s` steps into a function call, `continue`/`c` resumes until the next breakpoint or program end.

**Inspecting state.** `print <expr>` (or `p`) evaluates and prints any expression in the current scope, including dereferencing pointers, struct fields, array elements. `backtrace`/`bt` shows the full call stack — every function currently active, in order — directly visualizing the stack frames from C Lesson 2's memory model. `frame <n>` moves your inspection context to a specific frame in that stack (useful for seeing what a *caller's* local variables were, not just the currently paused function's).

**Core dumps.** When a process crashes (segfault, abort), the OS can write its entire memory image to disk at the moment of the crash — a core dump. `gdb <executable> <corefile>` loads this post-mortem, letting you run `backtrace` and `print` exactly as if the program were still paused live at the crash point, except the program is no longer running — this is essential for crashes you can't easily reproduce interactively (e.g. a crash that only happens in production under specific load, hours after the process started).

## Attempt

1. Write a small C (or Go, using `dlv` — Delve — as the debugger, which has an analogous command set) program with a function that computes something incorrectly due to a deliberate logic bug (not a crash, just a wrong-answer bug — e.g. an off-by-one in a loop bound). Set a breakpoint inside the function, run under the debugger, and use `next`/`print` to step through and observe exactly where a variable's value diverges from what it should be, rather than guessing from the final wrong output alone.

2. Set a breakpoint in a function that's called from two different places in your program, and use `backtrace` at that breakpoint on two separate hits (`continue` between them) to confirm the call stack differs — showing the two different call paths that reached the same function.

3. Deliberately cause a segfault (e.g. dereference a null pointer) in a C program. Enable core dumps if not already enabled (`ulimit -c unlimited` on most Linux shells before running), run the program so it crashes and produces a core file, then load it with `gdb <executable> <corefile>` and run `backtrace` to identify the exact line and call stack that led to the crash — without ever running the program interactively under the debugger.

4. Use `frame <n>` after the core-dump backtrace from step 3 to move up the call stack to the *caller* of the function that actually crashed, and `print` one of the caller's local variables — confirming you can inspect state beyond just the immediate crash site, which is often where the actual root cause (e.g. passing a null pointer that shouldn't have been null) is more visible than at the crash site itself.

## Verify

For step 1, produce the exact line/iteration where your stepped-through observation first shows the variable's value diverging from the correct expected value, and confirm this matches your independently reasoned prediction of where the off-by-one actually occurs.

## Failure drill

Attempt step 3's core dump analysis without having compiled with debug symbols (`-g` flag omitted). Observe that `backtrace` now shows only raw addresses (or `??`) instead of function names and line numbers, making the crash far harder to diagnose. Recompile with `-g` and repeat, confirming symbol names and line numbers now appear. Explain in your own words why debug symbols (a separate table mapping machine addresses back to source-level names and line numbers, generated at compile time and normally stripped from optimized production binaries to save space) are what make a debugger's output human-readable at all — without them, a debugger can still show you raw memory and registers, but not the source-level correspondence that makes it actually useful.

## Transfer

If TARDOC or Mahall has ever crashed or misbehaved in production in a way you diagnosed purely by reading logs and re-reasoning about the code rather than using a debugger, describe what a core dump (for a Go program, this would be a Go panic with a stack trace, which Go provides automatically without needing `ulimit -c` — note this as a specific difference from C's model) or an attached live debugger session would have shown you directly, that you instead had to infer indirectly from logs.

## Done when

You've used breakpoints and stepping to catch a logic bug's exact divergence point rather than only observing final wrong output, you've analyzed a real (not simulated) core dump with `gdb` to find a crash's call stack after the fact, and you can explain why debug symbols are necessary for a debugger's output to be useful, having directly seen the difference with and without them.
