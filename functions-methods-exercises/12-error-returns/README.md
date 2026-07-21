# Exercise 12 - Error Returns

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Practice functions and methods that return errors when requested work cannot be
completed.

Errors should be returned to the caller. The caller decides what to print.

## Required Data

Create a `BankAccount` type with:

- owner
- balance

## Required Methods

| Method | Responsibility |
|---|---|
| `Deposit` | add a positive amount |
| `Withdraw` | remove a positive amount when funds are enough |
| `Balance` | return the current balance |

## Required Function

Create a function that transfers money from one account to another.

## Main Requirements

In `main`, demonstrate:

1. valid deposit
2. zero or negative deposit
3. valid withdrawal
4. withdrawal with insufficient funds
5. valid transfer
6. invalid transfer

## Constraints

- Return errors for invalid operations.
- Do not print inside account methods.
- A failed transfer must not partially change only one account.

## Prove It Works

Print balances before and after each operation so failed operations are visible.
