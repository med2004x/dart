# Exercise 07 - File Notes

## What You Are Learning

Until now, data existed only in memory. When the process ended, the operating
system reclaimed that memory.

This exercise adds persistence:

```text
Go values in memory <-> JSON bytes <-> file on disk
```

It also introduces a boundary where failure is normal. Files can be missing,
unreadable, malformed, or unwritable. Every file operation therefore returns an
error that must be classified.

## Beginner Bridge: Persistence Is A Separate Action

Changing a Go variable changes memory only. It does not change a file.

```go
notes = append(notes, Note{ID: 1, Text: "learn files"})
```

At this point, the operating system still has no reason to update `notes.json`.
Saving is a second action:

```text
Go values -> JSON bytes -> file on disk
```

Loading is the reverse action:

```text
file on disk -> JSON bytes -> Go values
```

Similar programs:

| Program | Go value | File |
|---|---|---|
| Settings app | `Preferences` | `settings.json` |
| Game save | `SaveGame` | `save.json` |
| Bookmark tool | `[]Bookmark` | `bookmarks.json` |

The dangerous beginner mistake is pretending disk cannot fail. Disk is outside
your program. Missing files, broken JSON, and permission errors are ordinary
cases that your code must name and handle.

## Memory And Disk Are Different States

Suppose the program starts with:

```go
notes := []Note{}
```

Appending changes memory:

```go
notes = append(notes, Note{ID: 1, Text: "learn files"})
```

It does not change a file. Saving is a separate operation.

Likewise, reading a file produces bytes. It does not automatically create Go
structs. JSON decoding is another separate operation.

## Exported Fields And JSON Tags

`encoding/json` only serializes exported fields, whose names begin with uppercase
letters:

```go
type Note struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}
```

The Go value:

```go
Note{ID: 1, Text: "learn files"}
```

can become:

```json
{"id":1,"text":"learn files"}
```

The tags control the JSON names. Lowercase Go fields would silently disappear
from normal JSON encoding.

## Bytes Connect Files And JSON

File functions read and write `[]byte`:

```go
data, err := os.ReadFile(filename)
```

JSON functions convert between bytes and Go values:

```go
data, err := json.MarshalIndent(notes, "", "  ")
err := json.Unmarshal(data, &notes)
```

`&notes` gives `Unmarshal` an address it can update. Without an address, it
would receive a copy and could not fill the caller's variable.

## Loading Has More Than One Failure

Missing file on the first run is expected. Corrupt JSON is not.

Pseudocode:

```text
FUNCTION load notes(filename)
    read file bytes

    IF error means file does not exist
        RETURN empty notes and no error

    IF another read error occurred
        RETURN no notes and that error

    decode bytes into notes
    IF decoding failed
        RETURN no notes and decoding error

    RETURN notes and no error
```

Do not treat every read error as "first run." Permission failure and missing file
have different causes and require different action.

In Go, classify the missing-file case with:

```go
errors.Is(err, os.ErrNotExist)
```

## Saving Has Two Stages

```text
Go notes -> JSON encoding -> file write
```

Both can fail, so check both errors.

```text
FUNCTION save notes(filename, notes)
    encode notes as JSON bytes
    IF encoding failed
        RETURN error

    write bytes to filename
    IF writing failed
        RETURN error

    RETURN no error
```

`0644` is a Unix-style permission mode commonly used for files. On Windows,
`os.WriteFile` accepts it but Windows access control remains authoritative.

## Assigning The Next ID

`len(notes) + 1` is unsafe after deletion:

```text
existing IDs: 1, 3
length: 2
len + 1: 3 -> duplicate
```

A simple algorithm:

```text
next ID = 1
FOR each note
    IF note ID is at least next ID
        next ID = note ID plus 1
```

This is still local learning logic. A database normally owns concurrency-safe ID
generation.

## Worked Example: Persist Preferences

```go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
)

type Preferences struct {
    Theme    string `json:"theme"`
    FontSize int    `json:"fontSize"`
}

func loadPreferences(filename string) (Preferences, error) {
    data, err := os.ReadFile(filename)
    if errors.Is(err, os.ErrNotExist) {
        return Preferences{Theme: "light", FontSize: 14}, nil
    }
    if err != nil {
        return Preferences{}, fmt.Errorf("read preferences: %w", err)
    }

    var preferences Preferences
    if err := json.Unmarshal(data, &preferences); err != nil {
        return Preferences{}, fmt.Errorf("decode preferences: %w", err)
    }
    return preferences, nil
}

func savePreferences(filename string, preferences Preferences) error {
    data, err := json.MarshalIndent(preferences, "", "  ")
    if err != nil {
        return fmt.Errorf("encode preferences: %w", err)
    }

    if err := os.WriteFile(filename, data, 0644); err != nil {
        return fmt.Errorf("write preferences: %w", err)
    }
    return nil
}
```

`fmt.Errorf("read preferences: %w", err)` adds context while preserving the
original error for `errors.Is` and `errors.As`.

## Your Program

Build:

1. A note type with exported ID and text fields plus JSON tags.
2. `loadNotes(filename) ([]Note, error)`.
3. `saveNotes(filename, notes) error`.
4. `addNote(notes, text) []Note` with a non-duplicate ID.
5. A startup-load, add, save, print flow.

## Build In Checkpoints

1. Marshal one note and print the JSON.
2. Unmarshal fixed JSON and print the Go value.
3. Save JSON to a temporary file.
4. Load it back.
5. Handle a missing file as an empty starting state.
6. Add the ID logic.
7. Run twice and verify that prior notes remain.

```powershell
gofmt -w .\main.go
go run .
Get-Content .\notes.json
go run .
Get-Content .\notes.json
```

## Test Table

| Situation | Expected |
|---|---|
| file missing | empty notes, no error |
| valid empty JSON array `[]` | empty notes, no error |
| valid notes JSON | decoded notes |
| malformed JSON | decoding error |
| path is a directory | read or write error |
| two program runs | second run keeps first run's data |

## Failure Experiments

1. Write `{broken` into `notes.json` and run. The program must report a decode
   error, not silently erase the file.
2. Ignore the save error and use an invalid path. Explain why printing "saved"
   would be a lie.
3. Make the struct fields lowercase and inspect the JSON output.
4. Use `len(notes)+1`, create IDs 1 and 3, and show the duplicate.

## Production Boundary

`os.WriteFile` can leave a partially written file if the process or machine
fails during writing. Production file persistence often writes a temporary
file, flushes it, and atomically renames it. Concurrent processes also require
coordination. This exercise teaches serialization and error flow, not a
production database.

## You Understand This Exercise When

You can trace values through memory, bytes, JSON, and disk; distinguish missing
files from broken files; and explain every returned error.

References:

- [`os`](https://pkg.go.dev/os)
- [`encoding/json`](https://pkg.go.dev/encoding/json)
- [JSON and Go](https://go.dev/blog/json)
