# Project 10 - Concurrency, Locks, And Isolation

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Observe blocking, snapshots, serialization failures, and deadlock detection with
two real sessions.

## Open Sessions

Open two PowerShell windows:

```bash
.\scripts\connect.ps1
```

Run `session-a.sql` one statement at a time in A and `session-b.sql` one statement
at a time in B.

## Experiments

1. `SELECT ... FOR UPDATE` blocks a conflicting update.
2. inspect waiting state from `pg_stat_activity`.
3. compare two selects under Read Committed.
4. compare under Repeatable Read.
5. create a two-row deadlock.
6. avoid deadlock by locking IDs in consistent order.
7. run an atomic `balance = balance - amount` update.

## Rules

Never leave an accidental transaction open. End every session with:

```sql
ROLLBACK;
```

Serializable and repeatable-read failures require retrying the complete
transaction, not only the final statement.

## Failure Drill

Read a balance into application logic, calculate a new absolute value, and write
it concurrently from two sessions. Compare with one atomic relative update.

## Done Means

You can identify the blocker, waiter, held transaction, and correct recovery.

