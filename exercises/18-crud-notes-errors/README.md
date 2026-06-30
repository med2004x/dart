# Exercise 18 - CRUD Notes With Errors

## Goal

Build CRUD again, but add validation and error returns.

The earlier CRUD exercises used booleans for missing IDs. This one separates
different failure reasons:

- missing note
- invalid title
- invalid body
- duplicate ID

## Data Model

Create a `Note` type with:

| Field | Meaning |
|---|---|
| ID | stable note identifier |
| Title | short title |
| Body | note content |
| Archived | whether note is archived |

## Storage

Use a slice:

```text
[]Note
```

No files yet. No JSON yet. This is only about CRUD plus errors.

## Required Functions

Use plain functions.

| Function | Job |
|---|---|
| `validateNote` | check whether a note can be saved |
| `addNote` | add a note or return an error |
| `findNoteByID` | return note plus found boolean |
| `updateNoteBody` | update body or return an error |
| `archiveNote` | set archived to true or return an error |
| `deleteNote` | delete note or return an error |
| `printNotes` | print all notes |

## Error Rules

Define your own exact error messages, but enforce these rules:

| Situation | Result |
|---|---|
| blank title | error |
| blank body | error |
| duplicate ID on add | error |
| update missing ID | error |
| archive missing ID | error |
| delete missing ID | error |

Do not print inside validation or CRUD functions. Return the error and let
`main` print what happened.

## How The Go Flow Works

### Bool Versus Error

Use `bool` when there is only one ordinary question:

```text
was it found?
```

Use `error` when the caller needs to know what went wrong:

```text
title is blank
body is blank
ID already exists
note not found
```

### Add

Add should validate before appending.

Flow:

```text
validate note
if invalid, return original notes and error
check duplicate ID
if duplicate, return original notes and error
append note
return updated notes and nil error
```

### Update

Update should find by ID first. If the note exists, validate the new body before
changing the stored note.

Failed update must leave the slice unchanged.

### Delete

Delete should return an error for a missing ID. This is different from earlier
exercises where missing delete returned `false`.

## Main Program Requirements

In `main`:

1. Add at least three valid notes.
2. Try to add a note with a blank title.
3. Try to add a duplicate ID.
4. Find one existing note.
5. Update one existing note body.
6. Try to update a missing note.
7. Archive one existing note.
8. Delete one existing note.
9. Try to delete a missing note.
10. Print final notes.

## Proof Cases

| Case | Expected |
|---|---|
| valid add | note added |
| blank title add | error, no add |
| duplicate ID add | error, no add |
| update existing | body changes |
| update missing | error, no change |
| archive existing | archived becomes true |
| delete existing | note removed |
| delete missing | error, no change |

## Common Mistakes To Catch

- appending before validation
- returning an error but still changing the slice
- printing inside helper functions
- treating every error as the same failure
- ignoring returned errors in `main`

## Done When

You can show that every failed operation leaves the notes unchanged.
