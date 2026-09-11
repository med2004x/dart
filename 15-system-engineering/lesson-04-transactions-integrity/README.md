# Lesson 4: Transactions and Integrity Across Boundaries

## Objective

Understand how to maintain data integrity when a single logical operation spans multiple systems (a database plus an external API call, or multiple database transactions across module boundaries) — since database transactions alone (database-internals Lesson 5) cannot protect an operation that crosses outside the database itself.

## Prerequisites

Database-internals Lesson 5 (transactions/isolation — the guarantee this lesson explains the *limits* of), distributed Lesson 2 (idempotency — one of the standard tools for handling the gap this lesson identifies).

## Learn

**Why a database transaction's guarantee stops at the database's edge.** A database transaction gives you atomicity for operations *within that database* — either all the changes commit, or none do. But a real business operation frequently needs to do something *outside* the database as part of the same logical unit — charge a payment via an external API, send a notification, call another internal module's service boundary (system-engineering Lesson 3's modular monolith pattern, if the modules communicate over network calls rather than in-process function calls). The database's ACID guarantees say nothing about what happens if the external call succeeds but the database commit fails, or vice versa — this is a genuine correctness gap that no database feature alone can close.

**The classic problematic pattern: "charge then save," or "save then charge," each broken differently.** If you charge a payment via an external API *first*, then try to save the order record, and the database save fails (a transient database issue, a crash), you've charged the customer with no corresponding order record — a real, damaging inconsistency. If you save the order first, then charge, and the charge fails, you have an order record for a payment that never happened — also wrong, though arguably a less damaging failure direction (an order incorrectly marked as unpaid is easier to detect and remediate than a payment charged with no record of why).

**The outbox pattern: a standard, practical mitigation.** Rather than directly calling the external system as part of the primary operation, write a record of "what needs to happen" (e.g. "charge payment X for order Y") into an outbox table *within the same database transaction* as the order creation itself — since this is now purely a database-internal write, it benefits fully from the database's own atomicity guarantee (either both the order and the outbox entry are saved, or neither is). A separate background process then reads pending outbox entries and performs the actual external call, marking the entry complete once it succeeds — and because this external call can be retried safely (using distributed Lesson 2's idempotency-key pattern for the external call itself, if the external API supports it, as most payment APIs do), the two-step process closes most of the gap a naive single-step "call then save" or "save then call" leaves open.

**This doesn't achieve true distributed transactions — it's a deliberate, practical tradeoff.** The outbox pattern doesn't give you the same atomicity a single-database transaction provides; there's still a window where the order exists but the payment hasn't been charged yet (the background process hasn't run yet) — but this window is now *observable and recoverable* (you can query the outbox for pending entries and know exactly what's incomplete) rather than a silent, undetectable inconsistency, which is the realistic, achievable goal for cross-system integrity, distinct from the stronger but much harder-to-achieve goal of genuine distributed atomicity (which two-phase commit and similar protocols attempt, at real complexity and availability cost, largely out of this lesson's practical scope).

## Attempt

1. Design (on paper) the naive "call external API, then save to database" pattern for a real operation (e.g. TARDOC activating a subscription, involving both a database write and — per your project history — a mailer call as part of subscription lifecycle events). Explicitly trace what happens if the database save fails immediately after the external call (mailer/payment) succeeds.

2. Implement the outbox pattern for this operation: within a single database transaction, insert both the primary record (e.g. the subscription) and an outbox entry describing the external action needed (e.g. "send activation email for subscription X"). Confirm, using database-internals Lesson 5's transactional guarantees, that a simulated failure partway through (e.g. deliberately erroring out before the transaction commits) results in *neither* the subscription nor the outbox entry existing — genuine atomicity for this part, unlike the naive pattern's gap.

3. Implement the background outbox processor: a separate process (or goroutine) that polls for pending outbox entries, performs the actual external action, and marks the entry complete (or failed, with retry logic) — reusing distributed Lesson 2/7's idempotency and retry patterns for the external call itself.

4. Test the full flow end to end, including a simulated failure of the external call itself (not the database write): confirm the outbox entry remains in a "pending" or "failed" state (not silently lost) when the external call fails, and that your background processor correctly retries it, eventually succeeding once the simulated external failure is resolved — demonstrating the "observable and recoverable" property from Learn concretely, via an actual outbox table you can inspect mid-failure.

## Verify

Show the outbox table's actual contents at each stage of a test run: immediately after the primary transaction commits (entry present, marked pending), during a simulated external-call failure (entry still present, marked pending or failed-with-retry-count), and after eventual success (entry marked complete) — real, inspectable state, not just a description of the intended behavior.

## Failure drill

Implement the naive "call external API first, then save" pattern from step 1 for comparison, and deliberately trigger a database failure immediately after the external call succeeds (e.g. by injecting a forced error right before the commit in your test code). Confirm — using your simulated payment/email log versus your database's actual final state — that you now have a real, silent inconsistency: the external action definitely happened, but there's no database record reflecting it, and (this is the key point) no automated way to even detect this inconsistency occurred, since nothing recorded that an external call was ever attempted. Compare this directly to your outbox-pattern version's behavior under the equivalent failure and explain, using both concrete test results, why the outbox pattern's core benefit is making the failure *observable and recoverable*, not eliminating the possibility of failure entirely — the external call can still fail, but you now always have a durable record of exactly what's incomplete.

## Transfer

If TARDOC's actual subscription-activation flow (mailer wired into lifecycle events, per your project history) currently uses something closer to the naive pattern (calling the mailer directly as part of the activation logic, without an outbox-style durable record), describe specifically what inconsistency could currently occur if the mailer call succeeded but a subsequent database operation failed, or vice versa, and what migrating to an outbox-pattern-based design would need to look like for that specific flow.

## Done when

You've implemented and tested the outbox pattern for a real cross-system operation, confirmed the primary write and outbox entry are genuinely atomic via a real database transaction, and demonstrated — via a direct, side-by-side comparison against the naive pattern under the same simulated failure — the specific difference in outcome: a silent, undetectable inconsistency versus an observable, recoverable pending state.
