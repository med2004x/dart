# Project 01 - Decompose And Trace

## Problem

Given daily temperatures, return how many are below zero, equal to zero, and
above zero.

Input:

```text
[-2, 0, 5, -1, 3]
```

Output:

```text
below=2 zero=1 above=2
```

## Algorithm From Zero

You need three counters and one pass.

Pseudocode:

```text
below = 0
zero = 0
above = 0

FOR each temperature
    IF temperature < 0
        below = below + 1
    ELSE IF temperature == 0
        zero = zero + 1
    ELSE
        above = above + 1

RETURN all three counters
```

## Hand Trace

| Value | Below | Zero | Above |
|---:|---:|---:|---:|
| start | 0 | 0 | 0 |
| -2 | 1 | 0 | 0 |
| 0 | 1 | 1 | 0 |
| 5 | 1 | 1 | 1 |
| -1 | 2 | 1 | 1 |
| 3 | 2 | 1 | 2 |

## Complexity

Every value is visited once: `O(n)` time. Only three counters are stored:
`O(1)` extra space.

## Tasks

1. Implement `classify`.
2. test empty input.
3. test all-negative/all-zero/all-positive.
4. verify counts always sum to input length.
5. print a trace while learning, then remove debug output.

## Done Means

You can write pseudocode and predict every counter before running the function.

