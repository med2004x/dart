# Exercise 02 - Parameters

## Goal

Practice passing values into functions.

Parameters are inputs. A function with parameters should not need hardcoded data
inside its body for the values it is supposed to receive.

## Required Functions

| Function | Parameters | Responsibility |
|---|---|---|
| `printGreeting` | name | print a greeting for that name |
| `printUserBadge` | name, role | print both values clearly |
| `printOrderLine` | product name, quantity | print one order line |

## Main Requirements

In `main`:

1. Call `printGreeting` for at least two different names.
2. Call `printUserBadge` for at least two different users.
3. Call `printOrderLine` for at least three different products.

## Constraints

- Use parameters for changing values.
- Do not create separate functions like `printSaraGreeting`.
- Do not return values yet.

## Prove It Works

Your output should prove that the same function can behave differently depending
on the arguments passed to it.
