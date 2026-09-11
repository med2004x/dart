# Lesson 5: Undefined Behavior and Defensive C

## Objective

Recognize the specific categories of undefined behavior (UB) in C, understand why the language defines them this way instead of mandating a specific (if unhelpful) behavior, and build the habit of surfacing UB early with tooling rather than discovering it in production.

## Prerequisites

Lessons 2-4 (stale pointers, buffer overflows, and double-frees are all specific instances of UB — this lesson names the general category and covers additional forms).

## Learn

**What "undefined behavior" formally means.** The C standard defines the language's behavior for well-formed programs operating within specified bounds. Outside those bounds, the standard imposes *no requirement whatsoever* on what happens — not "it crashes," not "it does something reasonable," literally anything is a conforming outcome, including appearing to work correctly. This is a deliberate design choice, not an oversight: it lets compilers assume UB never happens and aggressively optimize based on that assumption, which is faster than inserting runtime checks everywhere — but it means a compiler is free to transform undefined-behavior-containing code in ways that produce surprising, non-local results far from the actual bug.

**Common UB categories, beyond what Lessons 2-4 already covered:**

- **Signed integer overflow.** Unlike unsigned overflow (which wraps, well-defined by the standard), signed overflow is UB. `INT_MAX + 1` is not guaranteed to wrap to `INT_MIN` — a compiler is permitted to assume it never happens and optimize accordingly, which has caused real, documented cases of security checks being silently eliminated by the optimizer (e.g. a check like `if (x + 1 < x)` intended to detect overflow can be optimized away entirely, since the compiler assumes signed overflow can't happen).
- **Invalid shifts.** Shifting by a negative amount, or by an amount ≥ the type's bit width (e.g. `1 << 32` on a 32-bit int), is UB — not a defined wraparound.
- **Strict aliasing violations.** Accessing an object through a pointer of an incompatible type (with some specific exceptions) is UB, because the compiler is permitted to assume pointers of different types never alias the same memory, and may reorder or cache reads/writes based on that assumption.
- **Uninitialized reads.** Reading a local variable before it's assigned is UB — the value isn't just "unpredictable garbage," it's formally undefined, and again the compiler may exploit this assumption in optimization.

**Why "it worked when I tested it" is not evidence of correctness for any of these.** UB's defining property (restated because it's the single most important idea in this lesson) is that a program can appear entirely correct under one compiler, one optimization level, one platform, and then fail under a different one — not because the environment changed the "true" behavior, but because there never was a defined behavior to begin with, only whatever a specific compiler happened to generate for that specific case.

## Attempt

1. Write a small program that deliberately overflows a signed `int` (e.g. `int x = INT_MAX; x = x + 1;`) and print the result. Compile and run it at `-O0` and separately at `-O2`, and compare the printed value — they may differ, which is itself evidence of UB (a well-defined operation would print the same result regardless of optimization level).

2. Write the overflow-check example from Learn (`if (x + 1 < x)` as an attempted overflow detector) and compile it at `-O2` with `-fsanitize=undefined` (UBSan, distinct from AddressSanitizer — it specifically targets UB like signed overflow and invalid shifts). Run it with an input that should trigger overflow and observe UBSan's report identifying the exact operation and line.

3. Write an invalid shift example (`int x = 1; int y = x << 35;` on a platform where `int` is 32 bits) and run it under UBSan. Record the exact diagnostic.

4. Write a function with an uninitialized local variable that's read (not written) before use, in a way the compiler might not immediately flag with basic warnings (e.g. reading it conditionally based on another parameter, so simple flow analysis can't always prove it's always uninitialized). Compile with `-Wall -Wextra` first and note whether it's caught, then with `-fsanitize=memory` if available on your platform (MemorySanitizer specifically targets uninitialized reads; note if it's unavailable on your platform and use `-Wall -Wextra`'s warning as a fallback).

## Verify

For steps 1-3, produce a table: the UB category, whether `-Wall -Wextra` alone caught it (often it does not, for many of these), and what UBSan's specific diagnostic said, including the exact source line it identified.

## Failure drill

Take the signed-overflow example from step 1 and step 2's overflow-check, and specifically compare the check's behavior at `-O0` (likely appears to work, catching the overflow as the programmer intended) versus `-O2` (the compiler may have eliminated the check entirely, assuming overflow can't happen). If your specific compiler/version doesn't demonstrate this exact elimination, research and note that this is a well-documented real-world compiler behavior (search "signed overflow optimized away" for concrete historical examples) even if your specific toy example doesn't trigger it — the point is understanding that the optimizer is permitted to do this, whether or not this particular case happens to trigger it on your machine.

## Transfer

Explain, using this lesson's UB categories as the concrete cases, why languages like Go and Rust define integer overflow behavior precisely (Go: signed overflow wraps, well-defined; Rust: panics in debug builds, wraps in release by default unless using checked arithmetic) rather than leaving it undefined the way C does — what does C's approach buy in exchange (typically: more aggressive optimization potential, since the compiler doesn't need to preserve a specific overflow behavior), and what does it cost (the class of extremely hard-to-reproduce bugs this whole lesson has been about).

## Done when

You've directly observed at least one case where UB produced different output at different optimization levels (not just read that this is possible), you can name which UBSan or AddressSanitizer flag would catch each category covered here, and you can explain in your own words why "no requirement whatsoever" is a more dangerous standard than "some specific, if unhelpful, defined behavior" would be.
