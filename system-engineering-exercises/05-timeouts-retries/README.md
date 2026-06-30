# Project 05 - Timeouts And Retries

## Goal

Call a slow dependency with enforced deadlines and bounded retry behavior.

## Starter

The server has:

- `/fast`
- `/slow`, which waits 500 ms
- `/flaky`, which fails before eventually succeeding

The client currently has no useful timeout policy.

## Checkpoints

1. Add a 100 ms per-attempt timeout.
2. Classify timeout versus permanent HTTP errors.
3. Retry only temporary failures.
4. Limit attempts to three.
5. Add exponential backoff beginning at 20 ms.
6. Add random jitter.
7. Enforce a 500 ms total operation deadline.

## Required Tests

- fast request succeeds once
- slow request times out within bound
- 400 is not retried
- 503 is retried
- attempts never exceed three
- total elapsed time remains bounded

## Failure Drill

Remove all timeouts, then call `/slow` repeatedly. Explain resource retention.
Next, retry a non-idempotent POST and explain the unknown-outcome problem.

## Done Means

Every external wait, retry count, and total operation duration has a testable
upper bound.

