# Project 11 - Merge Intervals

## Problem

Merge overlapping maintenance windows.

Input:

```text
[1,4], [3,6], [8,10]
```

Output:

```text
[1,6], [8,10]
```

## Algorithm

First sort intervals by start, then end.

```text
result = empty
FOR each sorted interval
    IF result empty OR interval starts after last result ends
        append interval
    ELSE
        extend last result end to maximum of both ends
RETURN result
```

## Trace

| Current | Last result | Action | Result |
|---|---|---|---|
| [1,4] | none | append | [1,4] |
| [3,6] | [1,4] | overlap, extend | [1,6] |
| [8,10] | [1,6] | separate | [1,6],[8,10] |

Decide whether touching intervals such as `[1,3]` and `[3,5]` overlap. Time
interval APIs often use half-open ranges `[start,end)` where they only touch.

## Complexity

Sorting dominates: `O(n log n)` time. The merge pass is `O(n)`. Output can hold
`O(n)` intervals.

## Tasks

1. validate start <= end.
2. avoid mutating caller input.
3. define touching policy.
4. test nested, equal, disjoint, unsorted, empty.
5. return a clear validation error.

## Done Means

You can trace why only the last merged interval needs comparison after sorting.

