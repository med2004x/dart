# Project 15 - Architecture Review

## Goal

Audit a running project through execution, failure injection, and code-boundary
inspection. Findings must cite evidence.

Use project 14 or the PostgreSQL task capstone.

## Review Procedure

1. Run all tests.
2. Trace one successful write.
3. trigger invalid input.
4. trigger unauthorized access.
5. trigger dependency timeout.
6. trigger storage failure.
7. hold a database lock.
8. saturate the connection/worker pool.
9. stop during an in-flight request.
10. restore a backup to a separate target.

## Inspect

- dependency direction
- data ownership
- transaction boundaries
- parameterized SQL
- external-call deadlines
- secrets and log redaction
- readiness/liveness behavior
- migration compatibility
- rollback procedure
- alert actionability

## Finding Format

Use `review-report.md`. Order findings by severity.

Do not state that a control works unless it was executed or the exact
unverified limitation is recorded.

## Done Means

The report separates confirmed defects, untested risks, and passed checks, with
commands and outputs for each.

