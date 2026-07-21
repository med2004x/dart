# Exercise 01 - Function Calls

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Practice creating and calling simple functions that do not return values.

This exercise is about call order: code inside a function runs only when that
function is called.

## Required Functions

| Function | Responsibility |
|---|---|
| `printHeader` | print a title for the program |
| `printDivider` | print a visual separator |
| `printFooter` | print an ending line |

## Main Requirements

In `main`, call the functions in this order:

1. `printHeader`
2. `printDivider`
3. print at least two normal lines directly from `main`
4. `printDivider`
5. `printFooter`

## Constraints

- Use plain functions.
- Do not use parameters yet.
- Do not return values yet.
- Keep all code in `main.go`.

## Prove It Works

Run the program and confirm:

- the header appears first
- the footer appears last
- the divider appears twice
- changing call order changes output order
