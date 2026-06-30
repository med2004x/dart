# Project 04 - Frequency Maps

## Problem

Count support tickets by status and return the most common status.

## Algorithm

A frequency map stores:

```text
value -> number of times seen
```

Pseudocode:

```text
counts = empty map
FOR each status
    counts[status] = counts[status] + 1
RETURN counts
```

Then scan the map to find the largest count.

## Trace

Input `[open, done, open]`:

| Status | Counts after |
|---|---|
| open | open:1 |
| done | open:1, done:1 |
| open | open:2, done:1 |

## Complexity

Average map insert/lookup is treated as `O(1)`. Counting `n` values is `O(n)`
time. If there are `k` distinct statuses, extra space is `O(k)`.

Map iteration order is not guaranteed. Define a tie rule for "most common."

## Tasks

1. normalize status case/whitespace.
2. count frequencies.
3. return most common and count.
4. choose alphabetical tie-break.
5. test empty input and ties.

## Failure Drill

Replace `+=` with `=` and show why repeats are lost. Use map iteration order as
a tie rule and observe nondeterminism.

## Done Means

You can distinguish raw input count from number of distinct keys.

