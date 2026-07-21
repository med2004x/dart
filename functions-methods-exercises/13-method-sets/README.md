# Exercise 13 - Method Sets

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Practice the difference between value receiver methods and pointer receiver
methods when a value is used through an interface.

## Required Interface

Create an interface for something that can be started and stopped.

## Required Data

Create a `Service` type with:

- name
- running state

## Required Methods

| Method | Receiver kind | Responsibility |
|---|---|---|
| `Name` | value receiver | return service name |
| `Running` | value receiver | return running state |
| `Start` | pointer receiver | change running to true |
| `Stop` | pointer receiver | change running to false |

## Main Requirements

In `main`:

1. Create one service value.
2. Call read-only methods.
3. Call mutating methods.
4. Store the correct value form in an interface variable.
5. Prove start and stop changed the original service.

## Constraints

- Use pointer receivers for mutating methods.
- Do not hide state changes inside global variables.
- Record what fails if you try to use the non-pointer value through the
  interface.

## Prove It Works

Your output should show the service moving from stopped to running and back to
stopped.
