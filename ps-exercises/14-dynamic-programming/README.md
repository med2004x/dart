# Project 14 - Dynamic Programming

## Problem

Count ways to reach step `n` when each move climbs one or two steps.

For `n=4`:

```text
1+1+1+1
1+1+2
1+2+1
2+1+1
2+2
```

Result: 5.

## Why Direct Recursion Repeats Work

```text
ways(5) asks ways(4) and ways(3)
ways(4) asks ways(3) and ways(2)
```

`ways(3)` is recalculated.

## Recurrence

To reach step `n`, the final move came from:

- `n-1` with one step
- `n-2` with two steps

Therefore:

```text
ways(n) = ways(n-1) + ways(n-2)
```

Base cases:

```text
ways(0) = 1
ways(1) = 1
```

## Bottom-Up Algorithm

```text
previous two = 1
previous one = 1
FOR step from 2 through n
    current = previous one + previous two
    shift previous values
RETURN previous one
```

Dynamic programming stores/reuses solutions to overlapping subproblems.

## Complexity

Naive recursion grows exponentially. Bottom-up visits each step once: `O(n)`
time and `O(1)` extra space with two stored values.

## Tasks

1. reject negative n.
2. implement naive recursion for small n.
3. count recursive calls.
4. implement memoized version.
5. implement bottom-up constant-space version.
6. test 0, 1, 2, 5, and overflow boundary.

## Done Means

You can identify the state, recurrence, base cases, and evaluation order.

