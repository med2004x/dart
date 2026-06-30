# Exercise 03 - Return Values

## Goal

Practice functions that calculate a result and return it to the caller.

The function should compute. The caller should decide what to print or do next.

## Required Functions

| Function | Inputs | Return |
|---|---|---|
| `add` | two integers | their sum |
| `multiply` | two integers | their product |
| `discountedPrice` | price, discount amount | final price |
| `isPassingScore` | score | whether score is at least 50 |

## Main Requirements

In `main`:

1. Call each function with at least two inputs.
2. Store at least one returned value in a variable.
3. Print the returned results from `main`.

## Constraints

- Do not print from inside these functions.
- Do not mutate global state.
- Keep each function focused on one calculation.

## Prove It Works

Use output to prove:

- different arguments produce different results
- `isPassingScore` returns true at 50 and false below 50
- `discountedPrice` handles a zero discount
