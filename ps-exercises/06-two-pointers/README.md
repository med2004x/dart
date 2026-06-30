# Project 06 - Two Pointers

## Problem

Given a sorted slice and target, find two numbers whose sum equals the target.

Input:

```text
numbers = [1, 3, 4, 7, 10]
target = 11
```

Output can be `1 + 10` or `4 + 7`, according to a documented first-match rule.

## Algorithm

Place one pointer at each end:

```text
left = 0
right = last index

WHILE left < right
    sum = numbers[left] + numbers[right]
    IF sum equals target
        return pair
    IF sum is too small
        move left rightward
    ELSE
        move right leftward
return not found
```

Why it works: the slice is sorted. If the sum is too small, keeping the smaller
left value cannot help; move it upward. If too large, reduce the right value.

## Trace

Target 11:

| Left value | Right value | Sum | Action |
|---:|---:|---:|---|
| 1 | 10 | 11 | found |

Target 8:

| Left | Right | Sum | Action |
|---:|---:|---:|---|
| 1 | 10 | 11 | right-- |
| 1 | 7 | 8 | found |

## Complexity

Each pointer moves inward at most `n` times: `O(n)` time and `O(1)` extra space.
The guarantee requires sorted input. Sorting first costs `O(n log n)` and may
change index meaning.

## Tasks

1. return original values and found bool.
2. test exact pair, no pair, duplicates, negatives, two values, empty.
3. define whether one element can be reused; normally no.
4. reject or document unsorted input.

## Done Means

You can explain which impossible pairs each pointer movement removes.

