# Lesson 6: Unsafe Rust and FFI

## Objective

Use `unsafe` Rust deliberately and minimally — wrapping a small C API via FFI (Foreign Function Interface), and writing a safe abstraction around genuinely unsafe operations with explicitly documented invariants the unsafe code depends on, validated with the same sanitizer discipline C track Lessons 4-5 introduced.

## Prerequisites

C track Lessons 1-5 (this lesson calls real C code, so understanding what's on the other side of the FFI boundary matters directly), Rust Lesson 4's failure drill (which already introduced `unsafe` as an explicit, auditable escape hatch — this lesson goes deeper on using it responsibly).

## Learn

**What `unsafe` actually unlocks — a precise, bounded list, not a general "turn off safety checks" switch.** `unsafe` blocks permit exactly five additional operations beyond safe Rust: dereferencing raw pointers, calling `unsafe` functions (including FFI calls into C), accessing/modifying mutable `static` variables, implementing `unsafe` traits, and accessing union fields. Everything else about Rust's safety guarantees (borrow checking on ordinary references, exhaustiveness checking, and so on) remains fully in effect *even inside* an `unsafe` block — `unsafe` doesn't disable Rust's type system, it specifically unlocks these five operations the compiler cannot otherwise prove are safe.

**FFI: calling C code from Rust, and why every such call is inherently unsafe from Rust's perspective.** When Rust calls a C function, Rust's compiler has no visibility into what that C function actually does — it cannot verify the C code respects Rust's ownership/borrowing rules, doesn't have a buffer overflow (C track Lesson 3), or doesn't leak memory (C track Lesson 4), because none of that is checked by Rust's compiler at all; C track's lessons on manual memory discipline apply in full to whatever's happening on the other side of this boundary. This is precisely why every FFI call is wrapped in `unsafe` — Rust is explicitly telling you "I cannot verify this is safe, you (the programmer) are asserting it is."

**Writing a safe abstraction around unsafe internals: the standard, correct pattern.** Rather than scattering `unsafe` FFI calls throughout a codebase (each one a place where Rust's guarantees don't apply), the idiomatic pattern is to write a thin, carefully-reviewed `unsafe` layer wrapping the C API, and expose a *safe* Rust API on top of it — the safe API's implementation contains `unsafe` code, but its *callers* interact with ordinary, safe Rust and get all of Rust's usual compile-time guarantees, because the wrapper's author has taken on the responsibility of upholding whatever invariants the unsafe internals actually require.

**Documenting invariants explicitly — the discipline that makes an `unsafe` abstraction trustworthy.** Every `unsafe` block should be accompanied by a comment (by convention, often literally prefixed `// SAFETY:`) explaining precisely why the operation is actually safe given the current context — e.g. "SAFETY: `ptr` is guaranteed non-null because we just checked it above" or "SAFETY: this FFI call's documented contract guarantees `buf` is not accessed after this function returns." This isn't optional style — it's the actual mechanism by which a future maintainer (or you, months later) can verify the unsafe code's correctness without having to re-derive the entire safety argument from scratch, and it directly parallels C track Lesson 6's capstone requirement to document ownership contracts explicitly, since C provides no compiler enforcement of them either.

## Attempt

1. Write a small C library (reusing or extending C track Lesson 6's capstone string-builder library is a natural choice) with a clear, documented API, and compile it as a static or dynamic library callable from Rust.

2. Write Rust FFI bindings for this C library (`extern "C"` function declarations matching the C API's signatures), and confirm you can call the C functions from Rust — wrapped in `unsafe` blocks as required, since raw FFI calls require it.

3. Write a safe Rust wrapper struct around the C library's functionality (e.g. a `SafeStringBuilder` struct whose methods internally call the unsafe C FFI functions, but whose public API takes and returns ordinary, safe Rust types) — implement Rust's `Drop` trait for this wrapper to ensure the underlying C resource is correctly freed automatically when the Rust wrapper goes out of scope, directly connecting Rust's ownership model (Lesson 1) to C's manual resource management (C track Lesson 4), bridging the two disciplines at this specific boundary.

4. Document the safety invariants for every `unsafe` block in your wrapper with explicit `// SAFETY:` comments, and validate your wrapper's correctness using AddressSanitizer (build with Rust's `-Z sanitizer=address` on nightly, or run the underlying C library's own test suite with AddressSanitizer per C track Lesson 4's methodology) to confirm no memory-safety violations occur across the FFI boundary despite Rust's compiler being unable to verify this on its own.

## Verify

Show your Rust FFI bindings and your safe wrapper's public API (confirming it exposes no raw pointers or `unsafe` requirement to its own callers), your `// SAFETY:` comments for every unsafe block explaining the specific reasoning, and your sanitizer validation results confirming no memory issues across the FFI boundary.

## Failure drill

Deliberately violate one of your documented safety invariants in your C library (e.g. modify the C code to occasionally return a pointer that's actually already been freed, violating an invariant your Rust wrapper's `// SAFETY:` comment assumed held) and observe what happens when your Rust wrapper is used normally — since Rust's compiler cannot verify C-side invariants, this violation will not be caught at compile time, and you should observe either a crash, memory corruption, or (worse) code that appears to work by chance, exactly like C track Lesson 5's undefined-behavior discussion, now happening specifically because the unsafe FFI boundary let a violated invariant through undetected. Explain why this demonstrates the real, serious responsibility that comes with writing `unsafe` code: Rust's compiler protected every caller of your *safe* wrapper API right up until the wrapper's own internal assumption about the C library's behavior turned out to be wrong — the safety of the entire abstraction rests entirely on the correctness of the human-reasoned `// SAFETY:` justification, which is exactly why that reasoning needs to be genuinely rigorous, not just present as a comment.

## Transfer

If any of TARDOC's performance-critical processing were ever implemented as a Rust component calling into an existing C library (a plausible scenario given systems-programming ambitions in your project history), describe what specific safety invariants you'd need to document and verify at that FFI boundary, and what sanitizer-based validation strategy (per this lesson's step 4) you'd want in your CI pipeline (system-engineering track, supply-chain security Lesson 6) to catch regressions in those invariants over time, rather than relying on a one-time manual review.

## Done when

You've written working Rust FFI bindings to a real C library and wrapped them in a safe, ergonomic Rust API with correct automatic resource cleanup via `Drop`, every `unsafe` block has an explicit, genuinely justified `// SAFETY:` comment, you've validated the wrapper with sanitizer tooling, and you've directly demonstrated — via the failure drill — what happens when an unsafe abstraction's underlying safety invariant is violated on the C side, understanding that Rust's guarantees at this boundary are only as strong as the human reasoning documented in your safety comments.
