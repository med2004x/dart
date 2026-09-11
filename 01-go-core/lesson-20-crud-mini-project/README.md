# Exercise 20 - CRUD Mini Project

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Build a larger CRUD program that combines the slice, method, validation, and
error ideas from the previous CRUD exercises.

You will build a simple course enrollment manager.

## Data Model

Create:

| Type | Fields |
|---|---|
| `Course` | ID, title, capacity, enrolled count |
| `CourseStore` | slice of courses |

## Required Methods

Put CRUD behavior on `CourseStore`.

| Method | Job |
|---|---|
| `AddCourse` | validate and add a course |
| `FindCourseByID` | return course plus found boolean |
| `UpdateTitle` | change a course title |
| `UpdateCapacity` | change capacity safely |
| `EnrollStudent` | increase enrolled count when capacity allows |
| `DeleteCourse` | remove a course |
| `PrintCourses` | print all courses |

## Validation Rules

Enforce these rules:

| Rule | Failure |
|---|---|
| course title cannot be blank | error |
| capacity must be positive | error |
| duplicate course ID is not allowed | error |
| capacity cannot be set below enrolled count | error |
| cannot enroll past capacity | error |
| missing course ID | error or false, depending on method contract |

Choose whether each method returns `bool` or `error`, but be consistent:

- use `bool` only when found/not-found is the whole story
- use `error` when there are multiple failure reasons

## How The Go Design Should Work

### Store Owns State

`CourseStore` owns the slice. Methods that add, update, enroll, or delete should
operate on the store, not on unrelated global variables.

### Methods Own Use Cases

`main` should not contain the loops for add, update, enroll, or delete.

`main` should:

- create the store
- call methods
- print success or failure
- print final state

### Failed Operations Must Not Mutate State

If validation fails, the store should stay the same.

Examples:

- duplicate add should not overwrite existing course
- invalid capacity update should not change capacity
- over-capacity enrollment should not increase enrolled count

## Main Program Requirements

In `main`:

1. Create an empty store.
2. Add at least three valid courses.
3. Try to add a duplicate course.
4. Try to add a course with invalid capacity.
5. Find an existing course.
6. Try to find a missing course.
7. Update a course title.
8. Update a course capacity.
9. Try to set capacity below enrolled count.
10. Enroll students until one course is full.
11. Try to enroll one more student in the full course.
12. Delete one course.
13. Try to delete a missing course.
14. Print final courses.

## Proof Cases

| Case | Expected |
|---|---|
| add valid course | course appears |
| add duplicate ID | error, no overwrite |
| add invalid capacity | error, no add |
| update title | title changes only for matching course |
| update capacity valid | capacity changes |
| update capacity below enrolled | error, no change |
| enroll within capacity | enrolled count increases |
| enroll past capacity | error, no change |
| delete existing | course removed |
| delete missing | failure reported |

## Common Mistakes To Catch

- letting `main` contain store mutation loops
- using methods with value receivers when the store must change
- changing state before validation succeeds
- allowing capacity below enrolled count
- returning success when nothing changed

## Done When

You can explain every method contract:

- what it receives
- what it changes
- what it returns on success
- what it returns on failure
- why the receiver is a pointer or value
