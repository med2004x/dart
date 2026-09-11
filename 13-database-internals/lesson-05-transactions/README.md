# Lesson 5: Transactions and Isolation

## Objective

Reproduce concrete concurrency anomalies (lost updates, write skew) under low isolation levels, and understand precisely which mechanism each isolation level uses to prevent which specific anomaly — not just memorize the isolation-level names.

## Prerequisites

OS Lesson 2 (data races — transaction anomalies are the database-level analog of the same underlying problem: concurrent access to shared state without adequate coordination).

## Learn

**Why isolation is a spectrum, not a binary.** Full serializability (transactions behave as if they ran one at a time, in some order, with no interleaving) is the strongest, safest guarantee — and the most expensive to provide, since it requires the most restrictive concurrency control. Real database engines offer multiple isolation levels precisely because stronger isolation costs more (more locking, more blocking, more aborted-and-retried transactions), and many workloads don't need the strongest guarantee for every operation — understanding which anomalies each level actually prevents is what lets you choose correctly rather than either over-paying for unneeded strictness or under-protecting against a real risk.

**Lost update.** Two transactions each read the same value, each independently compute a new value based on what they read, and each write back — the second write overwrites the first, silently losing its effect, even though neither transaction did anything individually wrong. Classic example: two concurrent "increment counter by 1" operations, both reading the same starting value 10, both computing 11, both writing 11 — the counter should be 12, but ends at 11, with one increment silently lost.

**Write skew.** A subtler anomaly where two transactions each read overlapping data, each makes a decision based on what they read, and each writes to *different* rows — individually, each transaction's write is consistent with what it read, but the combination violates an invariant that spans both rows. Classic example: a system requiring at least one of two on-call doctors to always be scheduled; both doctors independently check "is the other doctor still on call?", both see yes, both decide it's safe to remove themselves, both do so — the invariant (at least one doctor on call) is violated even though neither transaction's individual read-then-write logic was wrong in isolation.

**Isolation levels and what each specifically fixes.** *Read Uncommitted* (rarely used in practice) allows reading another transaction's uncommitted changes ("dirty reads") — essentially minimal isolation. *Read Committed* (PostgreSQL's default) prevents dirty reads — you only ever see committed data — but still permits lost updates and write skew, since it doesn't prevent the underlying interleaving that causes them. *Repeatable Read* additionally guarantees that if you read the same row twice within one transaction, you'll see the same value both times (preventing a different anomaly, non-repeatable reads) — PostgreSQL's implementation of this level, via MVCC (Multi-Version Concurrency Control — each transaction sees a consistent snapshot of the database as of when it started, rather than acquiring locks on every read) also happens to prevent lost updates, though this is a property of PostgreSQL's specific MVCC implementation, not a universal guarantee of "Repeatable Read" as an abstract isolation level across all database systems. *Serializable* (the strongest) prevents write skew as well, typically via detecting conflicting read/write patterns at commit time and forcing one of the conflicting transactions to abort and retry.

## Attempt

1. Using PostgreSQL (from the postgresql-engineering track) with two separate `psql` sessions (simulating two concurrent transactions), reproduce a lost update at `READ COMMITTED` isolation (PostgreSQL's default): both sessions `BEGIN`, both `SELECT` the same counter value, both independently compute value+1, both `UPDATE` and `COMMIT` — confirm the counter only increased by 1 total, not 2, demonstrating the lost update directly.

2. Repeat the same experiment but with both sessions using `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ` before their `BEGIN`. Confirm the outcome differs — typically, PostgreSQL's MVCC-based Repeatable Read will cause the second transaction's `UPDATE` to fail with a serialization/concurrent-update error when it tries to commit, forcing your application code to retry, rather than silently losing an update the way `READ COMMITTED` did.

3. Reproduce write skew: create two rows representing the "on-call doctor" scenario from Learn (e.g. a small table with two doctor rows, each with an `on_call boolean`). In two concurrent sessions at `REPEATABLE READ`, have each session check "is at least one *other* doctor currently on call" and, if so, set its own row's `on_call` to false. Run both concurrently (coordinate the timing manually between two terminal sessions) and confirm both can succeed, leaving zero doctors on call — violating the invariant, despite `REPEATABLE READ` having prevented the earlier lost-update case.

4. Repeat step 3's write-skew scenario at `SERIALIZABLE` isolation instead. Confirm PostgreSQL now detects the conflicting pattern and forces one of the two transactions to fail at commit time with a serialization error, preventing the invariant violation — directly demonstrating the specific additional protection `SERIALIZABLE` provides beyond `REPEATABLE READ`.

## Verify

For each of steps 1-4, report the actual final database state (the counter value, or the on-call doctor table's final contents) and the exact error message (if any) PostgreSQL returned — real, observed output, not predicted behavior, since isolation-level behavior varies by database engine and this lesson is specifically about PostgreSQL's actual implementation.

## Failure drill

Take your step 2 fix (`REPEATABLE READ` preventing the lost update via a commit-time error) and confirm your *application code* actually needs to handle that error explicitly — rerun the scenario but have the "losing" transaction's code simply ignore the serialization error rather than catching it and retrying. Confirm the increment is effectively lost anyway, not because the database allowed the anomaly, but because the application discarded the database's correct rejection instead of retrying the transaction. Explain why choosing a stronger isolation level is necessary but not sufficient on its own — the database can correctly detect and reject the unsafe interleaving, but the application must still contain retry logic for the rejection to actually result in correct final behavior, rather than an error that's silently swallowed and produces the exact same lost update the isolation level was supposed to prevent.

## Transfer

If TARDOC's billing logic or subscription-sweep scheduler (mentioned in your project history) ever reads a value, computes something based on it, and writes it back (e.g. incrementing a usage counter, or checking-then-updating a subscription status), describe, using this lesson's lost-update and write-skew examples directly, what isolation level PostgreSQL is likely running at by default for that code path (probably `READ COMMITTED`, PostgreSQL's default, unless explicitly overridden), and whether the specific read-then-write pattern in that code is vulnerable to a lost update under concurrent execution — and if the Celery beat scheduler could ever run overlapping instances of the same task (worth checking explicitly, since scheduler misconfiguration allowing overlapping runs is a realistic way this exact class of concurrency bug could occur silently in production).

## Done when

You've directly reproduced a lost update at `READ COMMITTED` and confirmed `REPEATABLE READ` prevents it (with an explicit commit-time error, not silent success), you've directly reproduced write skew at `REPEATABLE READ` and confirmed `SERIALIZABLE` prevents it, and you can explain — using the failure drill — why choosing the correct isolation level alone isn't sufficient without matching retry logic in the application.
