# Lesson 4: Allocation and Resource Ownership

## Objective

Master heap allocation (`malloc`/`free`) and the ownership discipline it requires, and directly observe the two classic failure modes — leaks and double-frees — with tooling that makes them undeniable rather than theoretical.

## Prerequisites

Lesson 2 (memory model — heap is one of the four regions covered there, this lesson goes deep on it), Lesson 3 (pointers — heap objects are always accessed through pointers).

## Learn

**Malloc and free, mechanically.** `malloc(n)` requests `n` bytes from the heap allocator and returns a pointer to the start (or `NULL` on failure — a return value real code must check, and a very common source of bugs when it isn't). `free(p)` returns that memory to the allocator, making it available for reuse. Critically, `free` does not zero out the memory or the pointer — after `free(p)`, `p` still holds the same address (now a **dangling pointer**), and the memory it points to may still contain the old bytes, may be reused by a subsequent allocation, or may be unmapped entirely, depending on allocator internals you don't control.

**Ownership is a discipline, not a language feature.** C provides no enforcement that every `malloc` has exactly one matching `free`. The rule you must maintain by convention: every allocation has exactly one "owner" responsible for freeing it exactly once, and no code should access the memory after that. Violating this produces one of:
- **Memory leak**: allocated memory never freed — the owner lost track of it. Not immediately catastrophic (the OS reclaims all process memory on exit), but fatal for long-running processes (a server that leaks per-request allocations will eventually exhaust memory).
- **Double free**: `free`ing the same pointer twice. Corrupts the allocator's internal bookkeeping structures (most allocators store metadata like block size adjacent to or interleaved with the allocated memory itself) — this can crash immediately, corrupt unrelated later allocations, or in older/simpler allocators, be exploitable as a security vulnerability.
- **Use-after-free**: accessing memory through a pointer after it's been freed — same class of "still compiles, technically dereferences a valid-looking address, but the memory's meaning is gone" bug as Lesson 2's stale stack pointer, except now the freed memory might have already been handed to a completely unrelated allocation.

**Why "just always free everything" isn't sufficient guidance.** The actual discipline that scales is: decide who owns each allocation *before* writing the code, not after debugging a leak. A function that allocates and returns a pointer transfers ownership to the caller (the caller must eventually free it) — this contract needs to be documented, because C's type system cannot express it (a `char*` return type looks identical whether it's owned or borrowed).

## Attempt

1. Write a dynamic integer vector (`typedef struct { int *data; size_t len, cap; } Vector;`) with `vector_push`, `vector_get`, and `vector_free` functions, using `malloc`/`realloc`/`free` internally. Test that it correctly grows past its initial capacity (push more elements than the initial allocation holds) and that all values remain correct after a `realloc`-triggered move.

2. Compile and run your vector's test program with AddressSanitizer and LeakSanitizer: `gcc -fsanitize=address -g -o test test.c`. Confirm a clean run (no leaks, no errors) when every `vector_free` is called correctly.

3. Deliberately introduce a leak: remove one `vector_free` call from your test program (create a `Vector`, use it, but never free it). Rerun with the sanitizer and record the exact leak report — it should identify the allocation site (file and line of the original `malloc`/`vector_push` that grew the buffer).

4. Deliberately introduce a double-free: call `vector_free` twice on the same `Vector`. Run with the sanitizer and record the exact error — compare its specificity (does it tell you both the original free location and the second, offending free location?) to what a crash with no sanitizer looks like on the same bug.

## Verify

Produce three sanitizer reports side by side: the clean run (step 2), the leak (step 3), and the double-free (step 4) — for each, note the specific line numbers the sanitizer identifies and confirm they correctly point at the actual bug you introduced, not just "somewhere in the program."

## Failure drill

Take the double-free bug from step 4 and run it *without* the sanitizer, in a normal (non-debug) build. On many systems this may not crash immediately, may crash with a generic and unhelpful "malloc(): double free detected" abort with no line information, or may silently corrupt a later, unrelated allocation that then fails mysteriously somewhere else in the program. Run it a few times if your platform's behavior seems inconsistent. Explain why debugging this class of bug without a sanitizer is disproportionately harder than the bug itself — the symptom (a crash or corruption somewhere else, possibly much later) is often far removed from the actual cause (the double-free), which is exactly the gap sanitizer tooling closes.

## Transfer

Go and Rust both eliminate manual `free` calls, but through fundamentally different mechanisms: Go uses a garbage collector (memory is freed automatically once nothing references it, at a runtime-determined point), while Rust's ownership system (previewed properly in the Rust track) enforces exactly one owner at compile time, with no runtime GC needed. State, based on this lesson's leak/double-free examples, which specific failure mode each approach eliminates by construction versus which it still allows (e.g., Go can still leak by holding a reference longer than intended, even without manual `free`; ask yourself whether Go's GC prevents that specific case or not).

## Done when

Your dynamic vector implementation is leak-free and correct under real sanitizer testing (not just "seems to work"), you've deliberately produced and diagnosed both a leak and a double-free using sanitizer output rather than guesswork, and you can articulate the ownership rule you followed (who allocates, who frees, when) for your vector implementation specifically, not just in the abstract.
