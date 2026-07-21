# Exercise 09 - Value Receiver Methods

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Practice methods that read struct state without changing it.

A value receiver is appropriate when the method only needs to inspect the value
or calculate something from it.

## Required Data

Create a `Book` type with:

- title
- author
- page count
- pages read

## Required Methods

| Method | Responsibility |
|---|---|
| `Summary` | return or print a readable book summary |
| `ProgressPercent` | calculate reading progress |
| `Finished` | report whether pages read is at least page count |

## Main Requirements

In `main`:

1. Create at least three books.
2. Call every method.
3. Include one unread book.
4. Include one partially read book.
5. Include one finished book.

## Constraints

- Use methods with value receivers.
- These methods should not modify the book.
- Handle zero page count deliberately.

## Prove It Works

Your output should show correct progress for unread, partial, and finished
books.
