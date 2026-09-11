# Lesson 2: Values, Addresses, and Memory Layout

## Objective

Understand exactly where a running C program's data lives — stack, static/global storage, heap — and why that placement determines an object's lifetime, not the type of the variable itself.

## Prerequisites

Lesson 1 (compilation — an executable's static/global data is laid out by the linker at this stage). Computer architecture Lesson 6 (virtual memory) is useful background: every address discussed here is a virtual address in the process's own address space.

## Learn

**Four regions, four different lifetime rules.**

- **Static/global storage.** Global variables and `static` locals exist for the entire program's lifetime, allocated once, at a fixed address determined at link time (within the process's virtual address space). Uninitialized globals go in `.bss` (zero-initialized by the OS at load time, doesn't take space in the binary file itself); initialized globals go in `.data`.
- **Stack.** Local variables inside a function. Allocated when the function is entered (mechanically, by decrementing the stack pointer register — Lesson 2 of computer architecture covered this register), freed automatically when the function returns. Fast (just pointer arithmetic, no allocator call), but the memory becomes invalid the instant the function returns — this is the entire mechanism behind the classic "returning a pointer to a local variable" bug.
- **Heap.** Explicitly allocated with `malloc` and explicitly freed with `free` (Lesson 4). Lives until you free it, regardless of which function allocated it — this is what makes heap allocation necessary for data that must outlive the function that created it.
- **Code/text.** The compiled instructions themselves, typically read-only at runtime — this is why writing through a function pointer to "modify code" segfaults on modern systems with proper memory protection.

**Why returning a pointer to a local variable is undefined behavior, not just bad style.** When a function returns, its stack frame is considered free — the memory isn't erased, but the next function call can and will overwrite it. A pointer to a former local variable technically still holds the old address, but reading through it after the function returns is reading memory the language no longer guarantees anything about. It might appear to work in a quick test (nothing's overwritten it yet) and then break in production the moment a different call pattern reuses that stack space — a classic "works on my machine, fails intermittently in production" bug with a precise mechanical cause.

**Pointer identity.** A pointer is just an address — a number. Two pointers are equal if and only if they hold the same address, regardless of what they point to conceptually. This matters when comparing pointers vs. comparing the values they point to (`p1 == p2` compares addresses; `*p1 == *p2` compares the pointed-to values) — a very common source of subtle bugs when the two comparisons are conflated.

## Attempt

1. Write a C program with a global variable, a `static` local inside a function, a plain (automatic) local inside a function, and a heap-allocated `int`. Print `&`(address of) each one. Run it multiple times and observe which addresses stay the same across runs and which change (with ASLR — Address Space Layout Randomization — enabled by default on modern systems, most will vary run to run, but their *relative order/region* should stay consistent — e.g. stack addresses should be consistently numerically distinct from heap addresses).

2. Write a function that returns a pointer to one of its local (automatic, non-static) variables. Call it, then call a second, different function that also uses local stack variables, and print the value through the first, now-stale pointer both before and after the second call. Observe the value change or become garbage after the second call overwrites that stack region — this makes the "reads work until sokmething reuses that memory" failure concrete rather than theoretical.

3. Use a debugger (`gdb` or `lldb`) to set a breakpoint inside a function with several local variables, and inspect the stack frame directly (`info frame` and `info locals` in gdb, or the equivalent). Confirm the addresses you see for local variables are numerically close together and distinct from the heap/global addresses observed in step 1.

4. Draw (on paper, before running anything) the expected memory layout — which region each variable lives in — for a small program with a mix of global, static-local, automatic-local, and malloc'd variables. Then run the program from step 1 with actual address printing and confirm your prediction was directionally correct (relative ordering of regions, not exact addresses, which vary due to ASLR).

## Verify

For step 2, print the pointer's dereferenced value at three points: immediately after the first function returns, after calling the second function, and explain explicitly which of these reads is defined behavior and which is not, even if the "not defined" read happens to print a plausible-looking number.

## Failure drill

Take step 2's stale-pointer bug and compile it with a sanitizer enabled: `gcc -fsanitize=address -g -o prog main.c`. Run it and observe AddressSanitizer catching and reporting the stack-use-after-scope violation explicitly, with a clear error message pointing at the exact line — compare this to the earlier "silent garbage value" behavior with no sanitizer. Explain why the unsanitized version doesn't necessarily crash (the memory is still technically mapped and readable, it's just no longer meaningfully yours) while the sanitized version catches it as an error regardless of whether it happened to still "look correct."

## Transfer

In Go, the compiler performs escape analysis and automatically moves a local variable to the heap if a pointer to it escapes the function (e.g. is returned) — meaning Go's equivalent of this lesson's central bug is structurally prevented by the compiler, not by programmer discipline. Run `go build -gcflags="-m"` on a small Go function that returns a pointer to a local variable, and find the "escapes to heap" message in the output — this is the compiler making exactly the decision a C programmer has to make manually (and can get wrong) explicit and automatic.

## Done when

You can correctly predict, before running, which memory region a given variable lives in, you've directly observed a stale-stack-pointer bug produce a real symptom (not just read about it), and you can explain why AddressSanitizer can catch this class of bug even when the raw, unsanitized program doesn't visibly crash.
