# Exercise 04 - Multiple Returns

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Practice returning a useful value plus a boolean that tells whether the value was
found or valid.

This is the same pattern used by search functions.

## Required Functions

| Function | Inputs | Return |
|---|---|---|
| `findNumber` | slice of integers, wanted integer | matching number, found boolean |
| `findName` | slice of strings, wanted name | matching name, found boolean |
| `safeDivide` | numerator, denominator | result, success boolean |

## Main Requirements

In `main`:

1. Search for an existing number.
2. Search for a missing number.
3. Search for an existing name.
4. Search for a missing name.
5. Divide by a nonzero denominator.
6. Try to divide by zero.

## Constraints

- Do not print inside the search/divide functions.
- Let `main` inspect the boolean and print success or failure.
- Do not panic on missing values or division by zero.

## Prove It Works

Your output should show both success and failure paths for every function.
