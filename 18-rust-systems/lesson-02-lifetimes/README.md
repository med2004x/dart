# Lesson 2: Lifetimes and APIs

## Objective

Understand lifetime annotations as the compiler's way of tracking how long a borrowed reference remains valid, design function APIs that return borrowed data correctly, and develop the judgment for when returning an owned value is simply the better, simpler design choice.

## Prerequisites

Lesson 1 (ownership and borrowing — lifetimes are specifically about reasoning precisely about how long a borrow, from Lesson 1, remains valid).

## Learn

**What a lifetime actually describes.** Every reference (`&T`) is valid for some scope — the lifetime — and the compiler needs to verify that a reference is never used after what it points to has been dropped (directly preventing C track Lesson 2's stale-pointer bug at compile time, the same underlying guarantee Lesson 1's move semantics provide for owned values, now extended to borrowed ones). Most of the time, the compiler infers lifetimes automatically without you writing anything — lifetime annotations (`'a` syntax) only become necessary when the relationship between input and output reference lifetimes is ambiguous enough that the compiler can't infer it on its own, most commonly when a function returns a reference derived from one of several possible input references.

**Why a function returning a borrowed reference needs to specify whose lifetime it's tied to.** Consider a function `fn longest(x: &str, y: &str) -> &str` (returns whichever string is longer) — the compiler cannot know, just from the signature, whether the returned reference's validity is tied to `x`'s lifetime, `y`'s, or something else, because the actual answer depends on runtime logic (which string turns out to be longer). The lifetime annotation `fn longest<'a>(x: &'a str, y: &'a str) -> &'a str` explicitly tells the compiler: the returned reference is valid for at most as long as *both* input references are valid — letting the compiler correctly reject any caller usage that would use the returned reference after either input has gone out of scope.

**This is not the compiler being overly cautious — it's preventing a real bug class.** Without this check, a caller could easily write code that uses the returned reference after one of the original strings has been dropped — exactly C track Lesson 2's stale-pointer bug, except now happening across a function boundary instead of within a single function's stack frame, which makes it considerably harder to catch by inspection alone in a language without this compile-time check.

**When to return owned data instead of a borrowed reference — a real design judgment, not just "borrowing is always better."** Returning a borrow (`&str`) avoids a copy, which matters for performance-sensitive code, but it ties the caller's usage to the lifetime of whatever the borrow is derived from — sometimes genuinely awkward for the caller (they must ensure the original data outlives every use of your returned reference, which can ripple lifetime constraints through their own code in ways that are hard to satisfy cleanly). Returning an owned value (`String` instead of `&str`) costs a copy but frees the caller from any lifetime entanglement — for many APIs, especially ones not in an extremely hot path, this simplicity is worth the modest performance cost, and choosing correctly between the two is a real API-design skill, not a matter of always picking whichever seems more "idiomatically Rust."

## Attempt

1. Write the `longest` function from Learn exactly as described, with the explicit lifetime annotation, and confirm it compiles and works correctly for various input pairs. Then remove the lifetime annotation and confirm the compiler rejects it, reading the actual error message to see how the compiler explains the ambiguity it cannot resolve on its own.

2. Write a function that deliberately tries to return a reference to a value created *inside* the function itself (a local variable, not a borrowed parameter) — confirm the compiler rejects this outright, with no lifetime annotation able to fix it, and explain in your own words why this is fundamentally different from the `longest` case: no lifetime annotation can express "valid as long as a value that no longer exists once the function returns," since that value's lifetime genuinely ends at the function boundary — this connects directly back to C track Lesson 2's stale-stack-pointer bug, which this Rust code would have to be rewritten to return an *owned* value to fix, not merely reworded with a different lifetime annotation.

3. Design and implement a small API (e.g. a simple text-parsing function) twice: once returning borrowed string slices (`&str`) referencing the original input, and once returning owned `String`s. Write a caller for each version that demonstrates a realistic use case, and note specifically what lifetime constraints the borrowed version imposes on the caller (e.g. the caller must keep the original input string alive for as long as they use the parsed results) that the owned version does not.

4. Benchmark (performance track Lesson 1's methodology) both versions from step 3 on a realistic workload (parsing many, or large, inputs) and report the actual measured performance difference — confirming, with real numbers, what the borrowed version's avoided-copy benefit is actually worth for this specific case, rather than assuming it matters without checking.

## Verify

Present your step 2 compiler error (confirming no lifetime annotation resolves it) with your explanation of why, and your step 3/4 borrowed-vs-owned comparison with actual benchmark numbers, connecting the measured performance difference to a real design recommendation for this specific case.

## Failure drill

Take your step 3 borrowed-reference API version and construct a caller scenario where the lifetime constraint genuinely becomes awkward — e.g. the caller wants to store the parsed results in a struct that needs to outlive the original input string's own scope (a realistic pattern: parse once, use the results later, after the original input might have gone away). Attempt to write this caller code and confirm the compiler rejects it due to the lifetime mismatch, then rewrite the caller using the owned-`String` version of your API instead and confirm it now compiles without issue. Explain, using this concrete, blocked-then-unblocked example, why the "always prefer borrowing to avoid a copy" instinct is not universally correct — for this specific caller pattern, the borrowed API's lifetime constraint is a genuine obstacle the owned version simply doesn't have, and recognizing this tradeoff (not defaulting to either choice) is the actual design skill.

## Transfer

If you were designing a Rust rewrite of a text/data-processing component from TARDOC or Lead Sourcer (per your project history's interest in systems-level work), describe, using this lesson's borrowed-vs-owned tradeoff, which approach you'd choose for a function that processes and returns substrings of a large input (e.g. extracting fields from a transcription result or a scraped webpage) — reasoning about both the performance characteristics (how large is the input, how hot is this code path) and the caller ergonomics (does the caller need the results to outlive the original input) rather than defaulting to either choice without justification.

## Done when

You've written and correctly annotated a function needing explicit lifetime relationships, you've triggered and understood a compiler rejection that no lifetime annotation could fix (the return-a-local-reference case), and you've benchmarked and compared borrowed versus owned API designs for a real case, with a concrete example (via the failure drill) of a caller pattern where the borrowed version's lifetime constraint becomes a genuine obstacle the owned version avoids.
