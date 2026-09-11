# Lesson 7: Capstone — High-Performance Systems Component

## Objective

Build one serious, benchmarked systems component in Rust, integrating ownership/borrowing, traits/error handling, concurrency, and (if relevant to your chosen component) async or unsafe/FFI — the Rust track's equivalent of the integration capstones throughout this curriculum, with an explicit emphasis on documenting both safety and performance invariants.

## Prerequisites

Lessons 1-6, completed. No new theory — this lesson forces the individually-practiced pieces (ownership, lifetimes, traits/enums, concurrency, async, unsafe/FFI) to compose in one real, substantial piece of software.

## Learn

There is no new material. The specific integration challenge this capstone tests: Rust's compile-time guarantees (Lessons 1, 4) are genuinely powerful, but composing them correctly across a larger system — where ownership needs to flow sensibly through multiple modules, where concurrent access patterns (Lesson 4) need to be designed holistically rather than bolted on, where error types (Lesson 3) need to compose cleanly across boundaries — requires design judgment beyond what any single lesson's isolated exercise demonstrates.

Rust's own ecosystem strongly rewards this kind of integration exercise specifically because Rust systems-programming crates (Tokio for async, various storage-engine and networking crates) are generally high-quality, well-documented references — using them as inspiration or dependencies for your capstone, rather than reimplementing everything from absolute scratch, is a reasonable and encouraged choice, consistent with this curriculum's general principle of building real things using strong existing tools where using them is itself the valuable exercise.

## Attempt

Choose **one** substantial component, integrating multiple lessons' concepts genuinely, not superficially:

**Option A: A small, concurrent in-memory key-value store.** Combines Lesson 1 (ownership of stored values), Lesson 4 (safe concurrent access via `Arc<Mutex<T>>` or a more sophisticated concurrent data structure), Lesson 3 (a proper `Result`-based error type for operations like "key not found"), and optionally Lesson 5 (an async network-facing API using Tokio, directly connecting to networking track Lesson 8's TCP-based key-value store option, now with Rust's stronger compile-time concurrency guarantees).

**Option B: A high-performance parser/serializer for a real data format.** Combines Lesson 2 (careful lifetime management for zero-copy parsing, borrowing directly from the input buffer rather than allocating new strings for every parsed field — a genuine, common Rust performance pattern), Lesson 3 (a well-designed enum representing the parsed data's structure, with proper error handling for malformed input), and benchmarked comparison (performance track Lesson 1) against a naive, allocation-heavy version to quantify the zero-copy approach's actual benefit.

**Option C: A Rust wrapper around a real C library, exposed as a safe, idiomatic Rust crate.** Combines Lesson 6 (FFI and unsafe abstraction) most centrally, with Lesson 1 (correct ownership/`Drop` semantics wrapping the C library's manual resource management) and Lesson 3 (translating C's error-code conventions into idiomatic `Result`-based Rust error handling).

Whichever option you choose, the requirements are:

1. **Genuine integration**, not superficial use of each concept — your final report should be able to point to specific places where, e.g., a lifetime annotation from Lesson 2 and a concurrency pattern from Lesson 4 had to be reasoned about together, not just each concept demonstrated in isolation within the same file.

2. **Real benchmarks** (performance track Lesson 1's rigor: multiple runs, reported variance) comparing your component against a reasonable baseline (a naive implementation, or an equivalent Go/C implementation if you have one from earlier tracks) — quantifying what Rust's approach actually bought you, in real numbers, for this specific component.

3. **Documented safety and performance invariants.** For any `unsafe` code (per Lesson 6's discipline), explicit `// SAFETY:` justifications. For performance-critical design decisions (e.g. choosing borrowed over owned data per Lesson 2, or static over dynamic dispatch per Lesson 3), a brief written justification citing your actual benchmark evidence, not just an assumption that the "more idiomatic" choice was automatically correct.

## Verify

Present your component's actual benchmark results (with proper variance reporting) against your chosen baseline, and your safety/performance invariant documentation — both should be complete enough that a reader unfamiliar with your specific reasoning could understand not just what your component does, but why it's built the way it is.

## Failure drill

Take one specific design decision in your capstone that you made based on an assumption rather than a measurement (a genuinely common and realistic situation — not every choice can be benchmarked before making it) and actually benchmark it now, after the fact. Report whether your assumption held up under real measurement, or whether it turned out to be wrong (either outcome is a legitimate, useful capstone finding) — and if wrong, describe what you'd change now that you have real evidence rather than an assumption. Connect this explicitly to performance track Lesson 1's core discipline: even experienced systems programmers' intuitions about performance are frequently wrong, and this capstone-scale exercise is meant to give you a real, personally-experienced example of that, not just the abstract warning.

## Transfer

Compare your capstone component's design to a real, production Rust crate solving a similar problem (e.g. compare a key-value store capstone to `sled` or a similar embedded database crate; compare a parser capstone to `serde`'s design patterns) — identify at least one specific design decision the production crate makes differently from yours, and reason about what additional requirement (broader use-case support, more extensive error handling, backward-compatibility constraints) likely drove that difference, given that your simplified capstone didn't need to handle the same breadth of real-world usage.

## Done when

You've built a genuinely integrated Rust component (not concepts demonstrated separately in the same project) with real, properly-measured benchmark results against a reasonable baseline, complete safety and performance invariant documentation, and you've identified — via the failure drill — at least one design assumption you'd made that real benchmarking either confirmed or overturned, with an honest account of which it was.
