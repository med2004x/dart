# Project 05 - Sorting Records

## Problem

Sort tasks by priority ascending, then creation order, then ID.

## What Sorting Means

A comparator answers which of two values should come first.

Rules:

```text
lower priority number first
if priority ties, earlier CreatedAt first
if time ties, lower ID first
```

The tie-breakers create deterministic total order.

## Use The Standard Library

Do not implement a sorting algorithm yet. Use `sort.Slice` because sorting
correctly and efficiently is established library behavior.

Comparator pseudocode:

```text
IF priorities differ
    return left priority < right priority
IF creation times differ
    return left time before right time
return left ID < right ID
```

## Complexity

General comparison sorting is typically `O(n log n)` time. `sort.Slice` mutates
the provided slice, so copy first if caller order must remain unchanged.

## Tasks

1. implement deterministic comparator.
2. decide mutate versus copy contract.
3. test each tie-breaker independently.
4. test empty and one-item slices.
5. verify original input policy.

## Failure Drill

Use only priority. Run tied values and explain why output contract is incomplete.

## Done Means

You can compare any two tasks by hand and predict which comes first.

