# Exercise 06 - Slice Parameters

## Goal

Practice functions that receive a slice and answer questions about the data.

These functions should read the slice, not change it.

## Required Functions

| Function | Input | Return |
|---|---|---|
| `countItems` | slice of strings | number of items |
| `totalScores` | slice of integers | sum of scores |
| `highestScore` | slice of integers | highest score and found boolean |
| `containsItem` | slice of strings, wanted item | found boolean |

## Main Requirements

In `main`:

1. Create at least one string slice.
2. Create at least one integer slice.
3. Call every function.
4. Test at least one empty slice case.

## Constraints

- Do not modify the input slices.
- Do not print inside these helper functions.
- Return a boolean for cases where an empty slice means "no result."

## Prove It Works

Your output should prove:

- count works for empty and non-empty slices
- total works for multiple scores
- highest score handles empty input
- contains returns true and false in different cases