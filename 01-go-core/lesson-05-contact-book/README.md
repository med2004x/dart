# Exercise 05 - Contact Book

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

CRUD means Create, Read, Update, and Delete. Most data-backed applications are
combinations of these four operations.

This exercise keeps data in a slice so you can see the algorithms before HTTP
and databases add more moving parts.

## Define The Contract Of Each Operation

Do not begin with function bodies. First define inputs, outputs, and failure.

| Operation | Inputs | Outputs | Failure signal |
|---|---|---|---|
| Add | contacts, new contact | updated contacts | none yet |
| Find | contacts, ID | contact, bool | `false` |
| Update email | contacts, ID, email | updated contacts, bool | `false` |
| Delete | contacts, ID | updated contacts, bool | `false` |
| List | contacts | contacts or printed view | none |

A slice passed into a function is a small descriptor pointing to an underlying
array. Changing an element may be visible to the caller. Appending may produce a
new underlying array. Returning the resulting slice makes ownership explicit.

## Create

Pseudocode:

```text
FUNCTION add contact(contacts, new contact)
    append new contact to contacts
    RETURN resulting contacts
```

Call:

```go
contacts = addContact(contacts, newContact)
```

Store the returned slice. `append` may return a different slice descriptor.

## Read

Search by stable ID, not by position:

```text
FUNCTION find contact by ID(contacts, wanted ID)
    FOR each contact
        IF contact ID equals wanted ID
            RETURN contact and true
    RETURN empty contact and false
```

Index 2 can change after deletion. Contact ID 2 should continue to identify the
same logical contact.

## Update

To mutate a struct stored in a slice, use the index:

```go
for index := range contacts {
    if contacts[index].ID == wantedID {
        contacts[index].Email = newEmail
        return contacts, true
    }
}
```

Why not update the range value?

```go
for _, contact := range contacts {
    contact.Email = newEmail
}
```

`contact` is a copy of each struct. Changing the copy does not replace the slice
element.

## Delete

For a beginner, build a new result:

```text
FUNCTION delete contact(contacts, wanted ID)
    result = empty contact slice
    found = false

    FOR each contact
        IF contact ID equals wanted ID
            found = true
            CONTINUE to next iteration
        append contact to result

    IF not found
        RETURN original contacts and false
    RETURN result and true
```

This is easy to reason about and preserves the order of remaining contacts.

## Worked Example: Playlist CRUD

```go
package main

import "fmt"

type Song struct {
    ID    int
    Title string
}

func addSong(songs []Song, song Song) []Song {
    return append(songs, song)
}

func findSong(songs []Song, wantedID int) (Song, bool) {
    for _, song := range songs {
        if song.ID == wantedID {
            return song, true
        }
    }
    return Song{}, false
}

func removeSong(songs []Song, wantedID int) ([]Song, bool) {
    result := make([]Song, 0, len(songs))
    found := false

    for _, song := range songs {
        if song.ID == wantedID {
            found = true
            continue
        }
        result = append(result, song)
    }

    if !found {
        return songs, false
    }
    return result, true
}

func main() {
    songs := []Song{}
    songs = addSong(songs, Song{ID: 1, Title: "First"})

    song, found := findSong(songs, 1)
    fmt.Println(song, found)

    songs, removed := removeSong(songs, 1)
    fmt.Println(songs, removed)
}
```

This demonstrates create, read, and delete. Write the playlist update operation
yourself before adapting the pattern to contacts.

## Your Program

Build a contact type with ID, name, email, and phone, then implement:

- add
- find by ID
- update email
- delete
- final listing

Use a slice, not a map. Return a boolean for operations that may not find an ID.

## Implementation Structure

Use plain functions for this exercise, not methods.

Reason: the earlier exercises are still building comfort with passing values
into functions and returning updated values. Methods can come later, after the
slice flow is clear.

Your `main.go` should have this shape:

1. `package main`
2. imports
3. `Contact` type definition
4. helper functions for the contact operations
5. `main` function that creates data and calls the helper functions

Create these helper functions:

| Function | Responsibility | Called from |
|---|---|---|
| `addContact` | add one contact to the slice | `main` |
| `findContactByID` | search for one contact by ID | `main`, optionally update/delete if you choose |
| `updateContactEmail` | change one contact's email by ID | `main` |
| `deleteContact` | remove one contact by ID | `main` |
| `printContacts` | print all contacts | `main` |

`main` should not contain the search, update, or delete loops directly. It
should call the helper functions and print the result of each operation.

Do not create:

- a `ContactBook` type yet
- methods like `contact.UpdateEmail(...)`
- a menu system
- user input prompts
- file saving
- maps
- HTTP handlers

Keep the exercise focused on one thing: operating on a slice of contacts.

## Function Requirements

### `addContact`

This function receives the current contact slice and one new contact.

It returns the updated contact slice.

It does not need to check whether the ID already exists in this exercise.

### `findContactByID`

This function receives the contact slice and an ID.

It returns:

- the matching contact when found
- a success/failure boolean

If no contact has that ID, it must report failure. It must not print the failure
itself. Let `main` decide what to print.

### `updateContactEmail`

This function receives the contact slice, an ID, and the new email.

It changes only the email of the matching contact.

It returns:

- the updated contact slice
- a success/failure boolean

If the ID is missing, it returns the original slice and reports failure.

### `deleteContact`

This function receives the contact slice and an ID.

It removes the matching contact if found.

It returns:

- the updated contact slice
- a success/failure boolean

If the ID is missing, it returns the original slice and reports failure.

### `printContacts`

This function receives the contact slice and prints each contact clearly.

It should not add, update, or delete anything.

## `main` Requirements

The `main` function is the demonstration script for the exercise.

It should:

1. create the initial empty slice
2. create at least three contact values
3. call `addContact` for each contact
4. call `findContactByID` with an existing ID
5. call `findContactByID` with a missing ID
6. call `updateContactEmail` with an existing ID
7. call `updateContactEmail` with a missing ID
8. call `deleteContact` with an existing ID
9. call `deleteContact` with a missing ID
10. call `printContacts`

For every operation that returns a boolean, `main` should print whether the
operation succeeded or failed.

## Build In Checkpoints

1. Define the contact and create three values.
2. Implement add and print the returned length.
3. Implement find and test existing/missing IDs.
4. Implement update and verify only one record changes.
5. Implement delete and verify length and remaining IDs.
6. Run all operations in sequence.

```bash
gofmt -w .\main.go
go run .
```

## Test Table

| Case | Expected |
|---|---|
| add to empty slice | length becomes 1 |
| find first/middle/last | correct contact, true |
| find missing ID | empty contact, false |
| update existing ID | one email changes, true |
| update missing ID | no changes, false |
| delete existing ID | length decreases by one, true |
| delete missing ID | original content, false |
| delete from empty slice | empty result, false |

## Failure Experiments

1. Update the `range` value instead of `contacts[index]`. Print the slice after
   the loop and explain why it did not change.
2. Call `append` without assigning its result. Add enough records to force
   growth and explain why relying on the old descriptor is wrong.
3. Delete by slice index, then reorder the slice. Explain why IDs and indexes
   are different concepts.

## Design Questions

This version permits duplicate IDs and invalid emails. That is intentional
unfinished design. Exercise 08 introduces errors and validation.

Before moving on, write answers:

- Who should assign new IDs?
- Should updating an email validate its format?
- What should happen if two contacts use the same ID?
- Should deleting a missing contact be an error or an ordinary false result?

## You Understand This Exercise When

You can state each CRUD contract, explain copy-versus-element mutation, and
implement the same operations for products, tasks, or songs.

References:

- [Structs](https://go.dev/tour/moretypes/2)
- [Slices](https://go.dev/tour/moretypes/7)
- [`range`](https://go.dev/tour/moretypes/16)
