# Lesson 6: Capstone — A Small C Library

## Objective

Integrate lessons 1-5 into a small, genuinely reusable C library with a documented ownership contract, a test suite, and clean sanitizer runs — the C-track equivalent of computer architecture Lesson 8's integration capstone.

## Prerequisites

Lessons 1-5, completed. No new theory here — this is deliberately an integration exercise, where gaps in individually-understood pieces (compilation, memory layout, pointers, ownership, UB) tend to surface when they have to work together under real constraints.

## Learn

There is no new material. The engineering discipline this lesson actually tests is: can you design and document an ownership contract *before* writing the implementation, rather than debugging your way to one afterward. This is the practical skill C forces you to develop explicitly, since (per Lesson 4) the language provides no mechanism to express or enforce ownership — only convention and documentation can.

A minimal but real library needs: a clear public API (a header file, since C separates interface from implementation via `.h`/`.c`, directly using Lesson 1's separate-compilation model), an explicit statement of who owns each pointer the API hands out or accepts, and tests that exercise both the happy path and the failure modes from Lessons 4-5.

## Attempt

Build a small **generic dynamic string builder** library (`strbuf.h` / `strbuf.c`) — deliberately different from Lesson 4's integer vector, so you're transferring the ownership/allocation pattern to a new problem rather than repeating the same code.

Required API surface (design the exact signatures yourself, but cover all of this):

1. `strbuf_create()` — allocates and returns a new, empty string builder.
2. `strbuf_append(sb, const char *s)` — appends a null-terminated string, growing the internal buffer (via `realloc`) as needed.
3. `strbuf_cstr(sb)` — returns a pointer to the current null-terminated contents. Document explicitly: is this pointer owned by the caller (must be freed separately) or borrowed (valid only until the next `strbuf_append` or `strbuf_free` call, per Lesson 4's ownership discipline)? Pick one and document it clearly in the header — this is the single most important design decision in the whole exercise.
4. `strbuf_free(sb)` — frees all memory owned by the string builder.

Requirements tying back to each earlier lesson:

- Compile with `-Wall -Wextra` (Lesson 1) and fix every warning, not just the ones that look serious.
- Document the memory layout implications of your design (Lesson 2) — does `strbuf_cstr`'s returned pointer become invalid after a subsequent `strbuf_append` (because `realloc` may have moved the buffer)? State this explicitly in the header comment, since it's exactly the kind of contract C's type system cannot express on its own.
- Implement the internal growth logic using pointer arithmetic where natural (Lesson 3), not just array indexing throughout.
- Follow the ownership discipline from Lesson 4 strictly — write down the ownership rule for every pointer the API returns before implementing the function that returns it.
- Ensure no signed overflow or invalid shift exists in your capacity-growth logic (Lesson 5) — a naive doubling strategy (`cap = cap * 2`) can overflow for very large buffers; decide whether you guard against this or explicitly document it as an out-of-scope limitation.

Write a test program exercising: creating and appending to a builder, growing past several reallocation boundaries, retrieving the string and confirming correctness, and freeing it cleanly.

## Verify

Run your test program under both AddressSanitizer and UBSan simultaneously (`-fsanitize=address,undefined`) and confirm a completely clean run — no leaks, no use-after-free, no UB — across a test sequence that appends enough strings to force multiple internal reallocations.

## Failure drill

Deliberately violate your own documented ownership contract in a *separate* test function: if you documented `strbuf_cstr`'s pointer as borrowed (invalidated by the next `strbuf_append`), write a test that holds onto that pointer across an `append` call and then reads it anyway. Run under AddressSanitizer and confirm it catches the violation with a specific diagnostic. This demonstrates the actual value of writing the contract down explicitly: the sanitizer can catch a violation of memory safety, but only a human (or a static analysis tool reading your documentation) can catch a violation of the *ownership contract specifically*, which is why documenting it clearly matters even though the language itself won't enforce it.

## Transfer

Compare your `strbuf_cstr` ownership decision to how Go's `strings.Builder` or a similar standard-library pattern in a language you know handles the equivalent "get the current contents" operation — does that language's design make the borrowed-vs-owned question moot (e.g., because strings are immutable and copying is cheap and automatic), and if so, explain specifically what C feature's absence (no garbage collector, manual memory management) is what forced you to make this design decision explicitly in the first place.

## Done when

Your string builder library has a clearly documented ownership contract in the header, compiles cleanly with `-Wall -Wextra`, passes a real test suite under combined AddressSanitizer and UBSan with zero issues, and you've demonstrated (via the failure drill) that you understand the specific gap between "memory-safe" and "contract-correct" — a sanitizer can catch the former, only careful design and documentation catches the latter.
