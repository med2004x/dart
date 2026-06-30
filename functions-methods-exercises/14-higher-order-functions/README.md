# Exercise 14 - Higher-Order Functions

## Goal

Practice functions that receive other functions.

This is useful for filtering, mapping, and customizing behavior without copying
the same loop many times.

## Required Data

Create a `Task` type with:

- title
- priority
- done flag

## Required Functions

| Function | Responsibility |
|---|---|
| `filterTasks` | receive tasks and a keep/reject function |
| `countTasks` | receive tasks and a match function |
| `printTasks` | receive tasks and a formatting function |

## Main Requirements

In `main`, create tasks and use higher-order functions to:

1. filter completed tasks
2. filter high-priority tasks
3. count unfinished tasks
4. print tasks in at least two formats

## Constraints

- Do not write a separate loop for every case.
- The filtering/counting functions should receive behavior as a function
  argument.
- Keep formatting separate from filtering.

## Prove It Works

Your output should show that the same loop function can be reused with different
behavior.
