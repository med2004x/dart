# Project 10 - Binary Search

## Problem

Find a target ID in a sorted slice.

## Prerequisite

Binary search requires sorted input. If input is unsorted, the movement logic is
invalid.

## Algorithm

Compare with the middle value and discard half:

```text
low = 0
high = last index

WHILE low <= high
    middle = low + (high-low)/2
    IF middle value equals target
        return middle
    IF middle value < target
        low = middle + 1
    ELSE
        high = middle - 1
return not found
```

## Trace

IDs `[2, 4, 7, 9, 12]`, target 9:

| Low | High | Middle | Value | Action |
|---:|---:|---:|---:|---|
| 0 | 4 | 2 | 7 | low=3 |
| 3 | 4 | 3 | 9 | found |

## Complexity

Each comparison halves remaining candidates: `O(log n)` time and `O(1)` extra
space for iterative implementation.

For 1,024 sorted items, at most about 11 checks are needed.

## Tasks

1. return index and found.
2. test first/last/middle/missing/empty.
3. define duplicate behavior.
4. implement first occurrence of duplicates.
5. compare comparison counts with linear search.

## Failure Drill

Use `low < high` and test a one-element remaining range. Forget `+1`/`-1` and
observe an infinite loop.

## Done Means

You can trace low, high, middle, and discarded range at every iteration.

