# Project 07 - Sliding Window

## Problem

Find the consecutive `k` days with the highest total sales.

Input:

```text
sales = [4, 2, 7, 1, 8]
k = 3
```

Windows:

```text
[4,2,7] total 13
[2,7,1] total 10
[7,1,8] total 16
```

Result starts at index 2 with total 16.

## Direct Solution

Calculate every window from scratch. There are roughly `n` windows and each
adds `k` values: `O(n*k)`.

## Sliding Window Algorithm

Adjacent windows share most values:

```text
new total = old total - value leaving + value entering
```

Pseudocode:

```text
validate 1 <= k <= length
sum first k values
best = first sum
best start = 0

FOR right from k to last index
    subtract value at right-k
    add value at right
    IF current total > best
        update best and start
RETURN start, best
```

## Complexity

First window costs `O(k)`, then each move costs `O(1)`: total `O(n)` time and
`O(1)` extra space.

## Tasks

1. implement valid-window behavior.
2. return a clear failure for invalid `k`.
3. define earliest-window tie behavior.
4. test negatives, k=1, k=n, empty input.

## Failure Drill

Subtract the wrong leaving index and trace the first window where totals diverge.

## Done Means

You can name exactly which value enters and leaves at every step.

