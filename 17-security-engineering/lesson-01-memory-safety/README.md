# Lesson 1: Memory Safety

## Objective

Understand the common memory-corruption vulnerability classes (buffer overflow, use-after-free) as exploitable security issues, not just correctness bugs — directly extending C track Lessons 3-5's coverage from "this produces undefined behavior" to "this is exactly the mechanism real attacks use."

## Prerequisites

C track Lessons 3-5 (buffer overflows, allocation bugs, undefined behavior — this lesson reframes those exact bugs through a security lens).

## Learn

**Why memory-safety bugs are security bugs, not just correctness bugs.** C track Lesson 3 demonstrated a buffer overflow as undefined behavior that might silently corrupt adjacent memory. From a security perspective, "adjacent memory" is a critical detail: on the stack, that adjacent memory can include the **return address** — the address the CPU jumps to when the current function returns (computer architecture Lesson 2's calling-convention mechanics). If an attacker can control the overflow's content precisely enough, they can overwrite that return address with an address of their choosing, hijacking control flow when the function returns — this is the foundational mechanism behind classic stack-based buffer overflow exploits, and it's the exact same bug C track Lesson 3 covered, just with an attacker deliberately constructing the overflow's content instead of it being accidental garbage.

**Use-after-free as an exploitation primitive, not just a crash.** C track Lesson 4 covered use-after-free as a bug that might silently work "by luck" if the freed memory hasn't been reused yet. From a security perspective, that "by luck" window is exactly what an attacker targets: if an attacker can control what gets allocated into the freed memory before the dangling pointer is used again, they can potentially control the data the vulnerable code ends up operating on — turning a memory-lifetime bug into a way to inject attacker-controlled data into a code path that assumes it's still reading the original, trusted object.

**Why modern mitigations exist, and what they actually protect against.** ASLR (Address Space Layout Randomization) randomizes where code and data are loaded in memory on each run, making it harder for an attacker to know in advance what address to jump to. Stack canaries place a known value between local variables and the return address; if a buffer overflow corrupts the canary on its way to overwriting the return address, the corruption is detected before the function returns. DEP/NX (Data Execution Prevention / No-Execute) marks memory regions as non-executable unless specifically needed for code, making it harder to simply inject and run attacker-supplied machine code directly. These mitigations raise the difficulty of exploitation significantly but don't make memory-safety bugs harmless — sophisticated exploitation techniques exist specifically to work around each of these, which is why the actual fix is not having the memory-safety bug in the first place, with mitigations serving as defense-in-depth (system-engineering Lesson 10), not a substitute for correct code.

**Why languages with memory safety guarantees (Go, Rust) eliminate entire vulnerability classes structurally.** C track Lesson 2's transfer task noted Go's compiler-enforced escape analysis prevents the stale-stack-pointer bug by construction. This isn't just a correctness convenience — it means an entire category of exploitable vulnerability (stack-based buffer overflow via return-address overwrite, use-after-free) simply cannot occur in safe Go or Rust code the way it can in C, which is a genuine, significant security property of the language choice itself, not merely a stylistic preference.

## Attempt

1. In a controlled, local sandbox (a VM or container you're comfortable potentially crashing, never a shared or production system), compile a deliberately vulnerable C program with a stack buffer overflow (reusing or extending C track Lesson 3's overflow example) with security mitigations *disabled* (`gcc -fno-stack-protector -z execstack -no-pie`) to make the underlying mechanism observable without fighting modern defenses. Use a debugger (Linux tools Lesson 4) to observe the stack layout and confirm you can identify exactly where the return address sits relative to the vulnerable buffer.

2. Demonstrate a controlled crash via return-address corruption: craft an input that overflows the buffer far enough to overwrite the return address with a clearly invalid value (not a working exploit — just enough to prove the mechanism), run it, and confirm the program crashes with a segfault at an address matching your injected (garbage) value — direct, observable proof that you controlled the crash location, the foundational primitive real exploitation builds on.

3. Recompile the same program with stack protection enabled (`gcc -fstack-protector-all`, the modern default in most cases) and rerun the same overflow. Confirm the program now detects the corruption and aborts with a "stack smashing detected" error *before* attempting to use the corrupted return address, directly demonstrating the stack canary mitigation from Learn actually functioning.

4. Fix the underlying vulnerability properly (bounds-checked input handling, per C track Lesson 3's correct string-handling practices) rather than relying on the mitigation alone, and confirm the fixed version handles the same malicious input safely without needing the canary to catch anything, since the actual bug — not just its exploitability — has been eliminated.

## Verify

Show your step 2 crash with the injected garbage address visible in the crash report (e.g. via `dmesg` or the program's own segfault output, per computer architecture Lesson 6), your step 3 stack-canary detection output, and confirm your step 4 fix handles the same malicious input without any crash or detection event at all — the vulnerability genuinely gone, not just caught.

## Failure drill

Take your step 3 canary-protected version and attempt to overflow the buffer with content that overwrites *past* the canary and return address into the caller's own stack frame variables, without ever touching the canary's actual bytes (a more surgical overflow than step 2's blunt approach — this requires knowing precisely how far the canary is from your buffer, which you can determine from step 1's debugger inspection). If you can construct such a case, confirm the canary does *not* detect this specific corruption (since it only checks its own bytes, not everything past it), demonstrating a canary's real, specific limitation: it protects against overflows that happen to corrupt it on the way to the return address, but a sufficiently precise or differently-targeted overflow can still cause harm without ever triggering the canary check. Explain why this is exactly why defense-in-depth (system-engineering Lesson 10) matters — the canary is one layer, not a complete solution, and relying on it alone rather than fixing the underlying bounds-checking issue leaves this kind of gap.

## Transfer

If any of your systems still use C for performance-critical components (or if you're evaluating whether to), describe, using this lesson's direct demonstration of exploitability (not just "undefined behavior is bad" in the abstract), what specific additional review discipline (mandatory bounds-checking, sanitizer-enabled testing per C track Lesson 4/5, avoiding unsafe string functions per C track Lesson 3) that code would need compared to equivalent Go code, and whether the performance benefit of C for that specific component is worth the security review burden this lesson has made concrete.

## Done when

You've directly demonstrated controlling a program's crash location via a stack buffer overflow in a safe, local sandbox, you've confirmed a stack canary mitigation actually detects the same overflow when enabled, you've fixed the underlying vulnerability properly rather than relying on the mitigation, and you've identified — via the failure drill — a real, specific limitation of the canary mitigation, understanding why defense-in-depth doesn't substitute for actually fixing the root cause.
