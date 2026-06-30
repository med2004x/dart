# Exercise 08 - Function Composition

## Goal

Practice building a larger operation from smaller functions.

One function should not do every job. Split the work so each function has a clear
responsibility.

## Required Program

Build a simple username registration checker.

## Required Functions

| Function | Responsibility |
|---|---|
| `trimUsername` | remove surrounding spaces |
| `isLongEnough` | check minimum length |
| `containsNoSpaces` | reject internal spaces |
| `usernameExists` | check whether username is already used |
| `canRegisterUsername` | combine the smaller checks |

## Main Requirements

In `main`, test usernames that are:

1. valid
2. too short
3. padded with spaces but valid after trimming
4. containing internal spaces
5. already taken

## Constraints

- Smaller functions should not print.
- `canRegisterUsername` should call the smaller functions.
- Do not duplicate validation logic in `main`.

## Prove It Works

Your output should show one line per username with accepted/rejected status and
a short reason.
