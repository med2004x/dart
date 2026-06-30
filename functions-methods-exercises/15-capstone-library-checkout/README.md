# Exercise 15 - Capstone: Library Checkout

## Goal

Combine functions and methods in one small program.

Build an in-memory library checkout model. The goal is to decide which behavior
belongs on a type as a method and which behavior belongs in a separate function.

## Required Data

Create:

- `Book` with ID, title, author, and checked-out state
- `Member` with ID and name
- `Library` with a slice of books

## Required Methods

| Type | Method | Responsibility |
|---|---|---|
| `Book` | read-only summary method | describe one book |
| `Book` | checkout-state method if appropriate | report whether it is available |
| `Library` | add book | add a book to the library |
| `Library` | find book by ID | find a book |
| `Library` | check out book | mark one book checked out |
| `Library` | return book | mark one book available |
| `Library` | list available books | show available books |

## Required Functions

Create at least one plain function for behavior that does not naturally belong
to one receiver. Examples:

- print a checkout receipt using both a member and a book
- compare two books
- print a report from a library and a member

## Main Requirements

In `main`:

1. Create a library.
2. Add at least four books.
3. Create at least two members.
4. List available books.
5. Check out an available book.
6. Try to check out the same book again.
7. Return the book.
8. Try to return a missing book.
9. Print a final library report.

## Constraints

- Use methods for behavior owned by `Book` or `Library`.
- Use a plain function for behavior that combines independent values.
- Return booleans or errors for operations that can fail.
- Do not use files, databases, HTTP, or maps.

## Prove It Works

Your output should prove:

- books can be added
- available books can be listed
- checkout changes state
- double checkout fails
- return changes state back
- missing book operations fail without crashing

## Completion Standard

You are done when you can explain every function and method choice:

- why it is a function or method
- whether it reads or mutates state
- what it returns on success
- what it returns on failure
