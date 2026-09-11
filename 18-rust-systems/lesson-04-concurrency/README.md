# Lesson 4: Safe Concurrency

## Objective

Build a worker pool using Rust's channels and shared-state primitives, and understand precisely which category of concurrency bug Rust's type system prevents at compile time — connecting directly back to OS Lesson 2's data race, now made structurally impossible rather than merely detectable via a runtime tool like Go's race detector.

## Prerequisites

OS Lesson 2 (threads, data races, mutexes — this lesson's entire point is contrasting Rust's compile-time prevention against that lesson's runtime-detected version of the same bug class), Lesson 1 of this track (ownership/borrowing — the mechanism that makes this lesson's guarantee possible).

## Learn

**`Send` and `Sync`: the two traits that make Rust's concurrency safety compile-time, not just a convention.** A type is `Send` if it's safe to transfer ownership of it to another thread. A type is `Sync` if it's safe for multiple threads to hold *shared references* to it simultaneously. These aren't traits you typically implement yourself — the compiler automatically determines them based on a type's structure, and critically, most standard types that would be unsafe to share (like `Rc`, Rust's non-thread-safe reference-counted pointer) are deliberately *not* `Sync`, meaning the compiler will flatly refuse to compile code that attempts to share one across threads — turning what would be a subtle, timing-dependent data race in C or a runtime-detected one in Go into a compile-time error that can never reach a running program at all.

**Why this is a stronger guarantee than Go's race detector.** OS Lesson 2 demonstrated Go's `-race` flag catching a data race — but only when the racing code path is actually *exercised* during a test run; a race condition in a rarely-hit code path can ship to production undetected if testing didn't happen to trigger it under the right timing. Rust's `Send`/`Sync` checking is a *static* analysis performed on every compile, regardless of whether any specific code path is exercised at runtime — if the code compiles, this specific class of data race is provably absent, not merely "not observed in this particular test run."

**Channels: message-passing concurrency, directly extending OS Lesson 6/Go's own channel idiom.** Rust's standard library provides `mpsc` (multiple-producer, single-consumer) channels, conceptually identical to Go's channels — one or more sender threads send values, a single receiver thread receives them, with the channel handling the necessary synchronization internally. Ownership transfers with each sent value (Lesson 1's move semantics) — meaning once a value is sent, the sending thread literally cannot access it anymore (the compiler enforces this), eliminating any possibility of the sender and receiver both trying to use the same data simultaneously, by construction.

**`Arc<Mutex<T>>`: the idiomatic pattern for genuinely shared, mutable state across threads.** `Arc` (atomically reference-counted pointer, the thread-safe counterpart to the non-thread-safe `Rc` mentioned above) allows multiple threads to share ownership of the same data; wrapping the data in a `Mutex` provides the actual synchronized access (directly analogous to OS Lesson 2's mutex, but here the mutex's lock/unlock is enforced by the type system too — you cannot access the data inside a `Mutex<T>` without going through its locking API, meaning it's structurally impossible to "forget" to lock before accessing the shared data, unlike C or Go where nothing prevents directly touching shared memory without acquiring the intended lock first).

## Attempt

1. Build a worker pool using `mpsc` channels: a fixed number of worker threads, a shared channel for incoming work items, and a channel for collecting results. Submit a batch of work items and confirm all results are correctly collected, with no lost or duplicated work — using Rust's channel-based ownership transfer, not manual synchronization.

2. Implement shared, mutable state across your worker threads using `Arc<Mutex<T>>` (e.g. a shared counter tracking total work completed, updated by every worker), and confirm the final count is exactly correct after all workers finish — directly reproducing OS Lesson 2's counter-increment correctness test, but now in Rust.

3. Deliberately attempt to share a non-`Sync` type (e.g. `Rc<RefCell<T>>`, Rust's single-threaded reference-counting/interior-mutability combo) across threads without the `Arc`/`Mutex` wrapper, and confirm the compiler rejects it outright with an error citing the missing `Send`/`Sync` bound — read the actual compiler error and identify which specific trait bound is unsatisfied.

4. Compare this lesson's Rust worker pool against OS Lesson 2's Go/C-style equivalent: attempt to deliberately introduce a data race in your Rust code by bypassing the `Mutex` somehow (e.g. by using `unsafe` code to get a raw, unsynchronized pointer to the shared counter — you'll need `unsafe` specifically because safe Rust structurally prevents this, which is itself the point) and confirm that doing so requires explicitly opting into `unsafe`, unlike C where the equivalent unsynchronized access requires no special marking at all.

## Verify

Show your step 1 worker pool's correct, complete result collection, your step 2 exactly-correct shared counter final value, and your step 3 actual compiler error text demonstrating the `Send`/`Sync` violation being caught before the program could ever run.

## Failure drill

Complete step 4's `unsafe`-code data race deliberately, and run it under real concurrent load, confirming you *can* still produce a genuine data race in Rust if you specifically opt into `unsafe` and bypass the type system's protections. Explain why this is not a flaw in Rust's safety model but an intentional, explicit escape hatch — `unsafe` code exists for cases (FFI with C, certain performance-critical patterns, implementing new safe abstractions) where the compiler's static analysis is too conservative to prove something is actually safe even though it is — but crucially, the `unsafe` keyword makes this escape hatch grep-able and auditable: a code reviewer (or you, six months later) can search for every `unsafe` block in a codebase and know exactly where Rust's compile-time guarantees are *not* being enforced, a genuinely different and better situation than C, where literally *every* line of code carries this same risk with no marker distinguishing "reviewed carefully because it's doing something inherently risky" from "ordinary code."

## Transfer

If you were implementing a concurrent processing pipeline for TARDOC's transcription jobs or Lead Sourcer's scraping/scoring logic in Rust instead of Go, describe, using this lesson's `Send`/`Sync` compile-time guarantee versus Go's runtime race detector, what specific advantage the Rust version would provide for a codebase with many contributors or infrequently-exercised code paths (where a Go race might exist in code that simply hasn't been tested under the exact timing that would trigger `-race` catching it) — and honestly assess whether that additional safety is worth the real learning-curve and development-velocity cost Rust's stricter compiler imposes, for your actual team size and project stage.

## Done when

You've built a working worker pool using Rust's channels with correct, verified result collection, you've implemented and verified correct shared-state synchronization via `Arc<Mutex<T>>`, you've triggered and read a real `Send`/`Sync` compiler error demonstrating compile-time race prevention, and you've deliberately used `unsafe` to reproduce a genuine data race, understanding why this required an explicit, auditable opt-out rather than being possible by accident the way it is in C.
