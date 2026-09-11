# Lesson 1: Ownership and Borrowing

## Objective

Understand Rust's ownership system as a compile-time solution to exactly the memory-safety problems C track Lessons 2-4 covered manually — translate real, unsafe C patterns into safe Rust, and learn to read and fix borrow-checker errors by understanding the rule being enforced, not by trial-and-error editing.

## Prerequisites

C track Lessons 2-4 (memory model, pointers, allocation — this lesson's entire framing is "here's how Rust prevents, at compile time, every bug those lessons demonstrated by hand").

## Learn

**The core idea: exactly one owner, enforced by the compiler, not by convention.** C track Lesson 4 established that ownership in C is a discipline you maintain manually — nothing stops two pointers from both believing they own the same allocation, leading to double-frees or use-after-free. Rust's ownership system makes this a compile-time rule: every value has exactly one owner at any point in the program, and when that owner goes out of scope, the value is automatically dropped (freed) — no garbage collector needed, and no manual `free()` call to forget, because the compiler tracks ownership statically and inserts the drop automatically at the right point.

**Move semantics: what happens when ownership transfers.** Assigning a value to a new variable, or passing it to a function, *moves* ownership by default (for types that don't implement `Copy`) — the original variable becomes invalid, and the compiler will refuse to compile any subsequent use of it. This directly prevents the "two pointers both think they own this" scenario from C track Lesson 4: after a move, there's no second, still-valid reference to the same data lying around to accidentally free or use.

**Borrowing: temporary, checked access without transferring ownership.** Since moving ownership on every use would be impractical (you often want to read a value without taking it over permanently), Rust provides borrowing: `&T` (an immutable, shared reference) or `&mut T` (a mutable, exclusive reference). The borrow checker enforces, at compile time: any number of immutable borrows can coexist, OR exactly one mutable borrow can exist — never both an immutable and mutable borrow simultaneously, and never multiple mutable borrows simultaneously. This is the direct compile-time prevention of OS Lesson 2's data race — you cannot have two simultaneous mutable accesses to the same data in safe Rust, full stop, checked before the program ever runs, not detected at runtime via a race detector the way Go's `-race` flag catches it after the fact.

**Why borrow-checker errors are the compiler protecting you, not an obstacle to work around.** A genuinely common mistake for newcomers to Rust is treating a borrow-checker error as a puzzle to solve by adding `.clone()` everywhere until it compiles, without understanding *why* the original code violated the ownership/borrowing rules. This defeats the entire purpose — the error is telling you about a real potential correctness issue (a use-after-move, or a potential simultaneous mutable/immutable access) that, in C, would have compiled fine and potentially failed unpredictably at runtime instead. Reading and understanding the specific rule being violated, not just making the error disappear, is the actual skill this lesson builds.

## Attempt

1. Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually). Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.

2. Deliberately reproduce C track Lesson 4's use-after-free bug's *intent* in Rust (attempt to use a value after it's been moved elsewhere) and confirm the Rust compiler refuses to compile it, with an error message explicitly citing the move — read the actual compiler error text and confirm it correctly identifies both the move point and the invalid subsequent use.

3. Write a function that takes an immutable borrow (`&Vec<i32>`) and one that takes a mutable borrow (`&mut Vec<i32>`), and deliberately attempt to call both simultaneously on the same vector (e.g. try to hold an immutable borrow across a call that requires a mutable one) — confirm the compiler rejects this, and read the error message to understand exactly which borrow-checker rule was violated.

4. Fix each borrow-checker error you triggered in steps 2-3 by restructuring the code correctly (not by reaching for `.clone()` reflexively) — e.g., ending an immutable borrow's scope before requesting a mutable one, or restructuring which function owns a value versus which merely borrows it — and explain, for each fix, specifically which ownership/borrowing rule your restructuring satisfies that the original violated.

## Verify

Show your actual Rust compiler error messages for steps 2-3 (not paraphrased — the real `rustc` output, which includes specific, often genuinely helpful guidance), and your corrected code for step 4 alongside a one-sentence explanation of which specific ownership rule each fix satisfies.

## Failure drill

Take one of your step 4 fixes and, instead of the structurally correct fix, apply the "just add `.clone()`" workaround instead — confirm this also compiles, but reason through (or construct a test that demonstrates) that cloning creates a fully independent copy, meaning any mutation to the cloned value will *not* be reflected in the original, potentially producing a logic bug where the C track Lesson 4-equivalent code (or your correctly-borrowed Rust fix) would have correctly shared and mutated the same underlying data. Explain why `.clone()`-as-a-reflex is a real anti-pattern specifically because it makes the code compile without addressing whether sharing or copying semantics were actually what the logic required — the compiler only checks memory safety, not whether your chosen fix preserves the intended *behavior*.

## Transfer

If TARDOC or Mahall has any performance-critical component currently written in C (or where C was considered, per your project history's DART curriculum reflecting systems-programming interest), describe, using this lesson's compile-time guarantees, what specific class of bug (from C track Lessons 2-4's coverage) a Rust rewrite of that component would eliminate by construction — not just "Rust is safer" as a general claim, but the specific mechanism (ownership, borrowing) that prevents the specific bug class relevant to that component's actual code patterns.

## Done when

You've translated a real C data structure into idiomatic Rust using ownership instead of manual memory management, you've triggered and correctly read real borrow-checker errors for both a use-after-move and a simultaneous mutable/immutable borrow violation, and you've fixed each with a structurally correct restructuring (not a reflexive `.clone()`) while being able to explain specifically why the reflexive fix would have been behaviorally different, not just less idiomatic.
