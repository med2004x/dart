# Project 02 - One-Pass Aggregation

## Problem

Given response times, return minimum, maximum, total, and average.

## Algorithm

Aggregation converts many values into a smaller summary.

Pseudocode:

```text
IF input is empty
    return not-found

minimum = first value
maximum = first value
total = 0

FOR each value
    add value to total
    IF value < minimum
        minimum = value
    IF value > maximum
        maximum = value

average = decimal total / count
RETURN summary and found
```

Initialize min/max from the first value, not zero. Zero is wrong for all-positive
minimums and all-negative maximums.

## Trace

Input `[30, 10, 50]`:

| Value | Min | Max | Total |
|---:|---:|---:|---:|
| start | 30 | 30 | 0 |
| 30 | 30 | 30 | 30 |
| 10 | 10 | 30 | 40 |
| 50 | 10 | 50 | 90 |

Average is 30.

## Complexity

One pass: `O(n)` time. Fixed summary: `O(1)` extra space.

## Tasks

1. define a `Summary` struct.
2. return `(Summary, bool)` for empty input.
3. preserve decimal average.
4. test one value, negatives, duplicates, empty input.

## Failure Drill

Initialize min/max to zero and test `[10, 20]` and `[-20, -10]`.

## Done Means

You can derive and defend every initial accumulator value.

