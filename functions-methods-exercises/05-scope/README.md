# Exercise 05 - Scope

## Goal

Practice where variables exist and where they do not.

Functions have their own local variables. A variable created inside one function
is not automatically available inside another function.

## Required Functions

| Function | Responsibility |
|---|---|
| `buildUsername` | receive first and last name, return one username string |
| `calculateTotal` | receive subtotal and tax, return total |
| `printReceipt` | receive customer name and total, print both |

## Main Requirements

In `main`:

1. Create input variables.
2. Call `buildUsername` and store the return value.
3. Call `calculateTotal` and store the return value.
4. Pass stored values into `printReceipt`.

## Constraints

- Do not use package-level variables for exercise data.
- Do not make `printReceipt` recalculate the total.
- Do not make `calculateTotal` know the customer name.

## Failure Checks

After the working version, deliberately try:

- using a variable outside the function where it was created
- redeclaring a variable with `:=` when you meant to assign with `=`

Read the compiler error, then restore the working code.
