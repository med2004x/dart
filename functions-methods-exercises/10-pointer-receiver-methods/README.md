# Exercise 10 - Pointer Receiver Methods

## Goal

Practice methods that change struct state.

When a method must modify the original struct, use a pointer receiver.

## Required Data

Create a `Counter` type with:

- name
- value

## Required Methods

| Method | Responsibility |
|---|---|
| `Increment` | increase value by one |
| `Add` | increase value by a supplied amount |
| `Reset` | set value back to zero |
| `Value` | return the current value |

## Main Requirements

In `main`:

1. Create one counter.
2. Increment it twice.
3. Add a positive amount.
4. Try to add zero or a negative amount and decide the behavior.
5. Print the value after each operation.
6. Reset it and print again.

## Constraints

- Mutating methods should use pointer receivers.
- The `Value` method can use a value receiver.
- Do not use global variables.

## Prove It Works

Your output should prove that method calls changed the original counter, not a
copy.
