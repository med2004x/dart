# Lesson 3: Traits, Enums, and Error Handling

## Objective

Model state and behavior explicitly using Rust's enum and trait system — building a typed state machine that makes illegal states unrepresentable — and implement idiomatic error propagation using `Result`, contrasting static and dynamic dispatch.

## Prerequisites

Discrete math Lesson 2 (sets/relations — a well-designed enum-based state machine directly encodes the valid states and transitions from that lesson's relation concepts), API engineering Lesson 12 (job state machines — this lesson revisits that exact problem with Rust's stronger type-system guarantees).

## Learn

**Enums in Rust are algebraic data types, not just named integer constants (unlike C's or Go's enums).** A Rust enum variant can carry different data per variant: `enum JobStatus { Pending, Processing { started_at: Instant }, Completed { result: String }, Failed { error: String } }` — each variant has exactly the data relevant to that state, and critically, the compiler enforces that you *cannot* accidentally access `result` data unless you're actually in the `Completed` variant (accessing it requires pattern matching, which forces you to handle every variant explicitly, or the compiler refuses to compile). This directly and structurally solves API engineering Lesson 12's failure drill concern — an invalid state transition, or accessing state that doesn't exist for the current status, becomes a compile-time impossibility rather than a runtime bug that needs a manually-written guard to prevent.

**Pattern matching with exhaustiveness checking.** A `match` expression over an enum must handle every variant (or include an explicit catch-all `_` pattern) — the compiler will refuse to compile a `match` that forgets a variant. This means if you later add a new variant to your enum (e.g. adding a `Cancelled` status to the job state machine), every `match` expression handling that enum throughout your codebase will fail to compile until you've explicitly decided how to handle the new case — a genuinely powerful correctness guarantee when evolving a system over time, structurally preventing the "we added a new status but forgot to update this one handler" class of bug.

**Traits: shared behavior across different types, without inheritance.** A trait defines a set of methods a type can implement (`trait Shape { fn area(&self) -> f64; }`), and multiple, unrelated types can implement the same trait — similar in spirit to Go's interfaces, but with additional capabilities (default method implementations, associated types) that make traits considerably more expressive for certain patterns.

**Static vs. dynamic dispatch: a real, measurable tradeoff, not just syntax.** Static dispatch (`fn process<T: Shape>(s: T)`, generic over a specific type known at compile time) lets the compiler generate specialized, inlined code for each concrete type used — no runtime overhead, but each distinct type used generates its own compiled code (larger binary, "monomorphization"). Dynamic dispatch (`fn process(s: &dyn Shape)`, a trait object) uses a runtime vtable lookup to call the correct implementation — one compiled function handles any type implementing the trait, smaller binary, but with a real, measurable per-call overhead compared to static dispatch's direct, inlined call.

**`Result<T, E>` and error propagation: making error handling a type-system-enforced part of the API, not an optional convention.** A function that can fail returns `Result<T, E>` rather than throwing an exception (which a caller could forget to catch) or returning a special sentinel value (which a caller could forget to check, unlike Go's multiple-return-value error convention, which relies entirely on discipline). The `?` operator propagates an error up the call stack concisely, but critically, the compiler forces every `Result`-returning call to be explicitly handled (via `?`, `match`, or an explicit `.unwrap()`/`.expect()` if you're deliberately choosing to panic) — silently ignoring a `Result` produces a compiler warning at minimum, a real, structural improvement over error-handling conventions that rely purely on programmer discipline to remember the check.

## Attempt

1. Model API engineering Lesson 12's job state machine as a Rust enum with per-variant data, following the `JobStatus` example in Learn. Write a `match` expression that handles every variant, and confirm the compiler rejects the match if you deliberately comment out handling for one variant.

2. Add a new variant to your enum (e.g. `Cancelled`) after you've already written several `match` expressions handling the original variants elsewhere in your code, and confirm every one of those existing match expressions now fails to compile until you explicitly add handling for the new variant — directly demonstrating the exhaustiveness-checking benefit from Learn with your own evolving code, not just the original design.

3. Define a trait with at least 2 different implementing types, and write both a statically-dispatched generic function and a dynamically-dispatched trait-object function that each use it. Benchmark (performance track Lesson 1's methodology) calling each version many times in a loop, and report the actual measured performance difference between static and dynamic dispatch for this specific case.

4. Implement a function that can fail (e.g. parsing a string into a specific structured type) returning `Result<T, E>` with a custom error type, and a caller using the `?` operator to propagate errors up through at least 2 levels of function calls. Deliberately trigger both the success and failure paths and confirm the error correctly propagates all the way up with the expected error information intact, not lost or replaced with something less specific along the way.

## Verify

Show your step 2 compiler errors demonstrating the exhaustiveness check catching your intentionally incomplete match expressions after adding a new variant, and your step 3 static-vs-dynamic dispatch benchmark numbers with an honest report of the actual measured difference (which may be small for many realistic workloads — report what you actually find, not what you expected to find).

## Failure drill

Take your step 4 error-propagation chain and deliberately use `.unwrap()` instead of the `?` operator or proper `Result` handling at one point in the chain (a common shortcut, especially during prototyping) — confirm the code still compiles (since `.unwrap()` is valid Rust, just risky), then trigger the failure path and observe the program panic and crash entirely, rather than propagating a handleable `Result::Err` up to a caller that could have responded gracefully. Explain why `.unwrap()` is a real, sharp escape hatch from Rust's otherwise-enforced error handling — the compiler cannot prevent you from choosing to panic instead of properly propagating an error, and recognizing `.unwrap()` calls in a codebase (your own, or one you're reviewing) as points where this safety net has been deliberately bypassed is an important code-review skill, not just a stylistic preference.

## Transfer

If you were modeling TARDOC's subscription lifecycle (mentioned in your project history) as a Rust enum instead of however it's currently represented, sketch what the variants and their associated data would look like, and identify at least one transition (e.g. subscription cancellation, or a billing failure state) that your current implementation (in whatever language it's actually written in) handles via a runtime check or convention, that Rust's exhaustive pattern matching would instead enforce at compile time.

## Done when

You've modeled a real state machine as a Rust enum and demonstrated the exhaustiveness-checking guarantee actually catching an incomplete match after adding a new variant, you've measured a real (even if small) performance difference between static and dynamic dispatch for a genuine use case, and you've implemented correct `Result`-based error propagation across multiple call levels while also directly demonstrating, via `.unwrap()`, how that safety guarantee can be deliberately bypassed and why recognizing that bypass matters.
