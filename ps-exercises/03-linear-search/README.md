# Project 03 - Linear Search

## Problem

Find a task by ID in an unsorted slice.

## Algorithm

Linear search checks items from beginning to end until one matches.

Pseudocode:

```text
FOR each task
    IF task ID equals wanted ID
        return task and true
return empty task and false
```

## Trace

IDs `[4, 9, 2]`, wanted `2`:

| Step | Current ID | Match |
|---:|---:|---|
| 1 | 4 | no |
| 2 | 9 | no |
| 3 | 2 | yes, return |

## When To Use It

Use linear search when:

- input is small
- data is not sorted
- building another index is unnecessary
- you search only once or rarely

Worst case visits all `n` tasks: `O(n)` time. Extra space is `O(1)`.

## Tasks

1. implement first match.
2. test first, middle, last, and missing IDs.
3. test empty input.
4. define duplicate-ID behavior.
5. write `findAllByStatus` to return multiple matches.

## Failure Drill

Return failure inside the loop after the first nonmatch. Search for the last
task and explain why it is never visited.

## Done Means

You can trace the exact number of comparisons for any input.

