# Project 10 - Rate Limits And Overload

## Goal

Bound abusive or accidental request volume and communicate when clients may
retry.

## Starter

Implement a token bucket per actor:

```text
capacity: 10 tokens
refill: 2 tokens/second
request cost: 1 token
```

When empty:

```text
429 Too Many Requests
Retry-After: seconds
```

## Checkpoints

1. identify actor before limiting.
2. bound the number of stored actor buckets.
3. synchronize bucket access.
4. use monotonic elapsed time.
5. return retry information.
6. expose accepted/rejected counters.
7. define behavior across multiple app instances.
8. distinguish rate limiting from concurrency limiting.

## Failure Drills

1. key only by client-provided user ID.
2. store unlimited IP entries forever.
3. use one global limit for all customers.
4. retry rejected requests immediately.
5. add app instances with independent limits and claim a global guarantee.

## Done Means

Load is bounded per trusted identity, memory is bounded, and distributed
limitations are documented.

