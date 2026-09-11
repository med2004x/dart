# Exercise 11 - JSON API

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise turns an HTTP body into a typed Go value and turns a typed Go
value back into an HTTP body:

```text
request bytes -> JSON decoder -> Go struct
Go struct     -> JSON encoder -> response bytes
```

An API is a contract. Field names, required values, methods, status codes, and
error shapes are observable behavior, not implementation details.

## Beginner Bridge: JSON Is The Wire Shape, Structs Are The Go Shape

Clients do not send Go structs. They send bytes. In this exercise, those bytes
are JSON:

```json
{"username":"adam","role":"admin"}
```

Your handler converts those bytes into a Go value:

```go
var input createUserInput
err := json.NewDecoder(r.Body).Decode(&input)
```

Then it converts a Go value back into JSON bytes:

```go
json.NewEncoder(w).Encode(user)
```

Similar APIs:

| API | Input JSON | Go input struct |
|---|---|---|
| create book | title | `createBookInput` |
| create user | username, role | `createUserInput` |
| create note | text | `createNoteInput` |

The main design rule: clients provide editable fields only. Server-owned fields
such as IDs are created by the server, not trusted from the request body.

## Model The Wire Format

Incoming JSON:

```json
{"username":"adam","role":"admin"}
```

Go type:

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Role     string `json:"role"`
}
```

Fields are exported so `encoding/json` can access them. Tags define exact JSON
keys.

Missing JSON fields receive Go zero values. `"username"` missing and
`"username":""` both produce an empty string unless the API models presence
separately. Validation must decide whether that is allowed.

## Decode Requires A Destination Address

```go
var input createUserInput
err := json.NewDecoder(r.Body).Decode(&input)
```

Execution:

1. `input` starts with zero values.
2. `&input` gives the decoder its memory address.
3. The decoder reads request bytes from `r.Body`.
4. Matching JSON fields replace fields in `input`.
5. A syntax or type mismatch returns an error.

Without `&`, the decoder cannot populate the caller's value.

Use a separate input type when clients must not choose server-owned fields:

```go
type createUserInput struct {
    Username string `json:"username"`
    Role     string `json:"role"`
}
```

Do not decode a client-supplied `id` and then pretend the server owns IDs.

## Encode A Response

Set headers and status before encoding:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
if err := json.NewEncoder(w).Encode(user); err != nil {
    // The response may already be partially committed; log the server error.
}
```

For small exercises, the encode error is uncommon but still exists. Production
handlers need a consistent response helper and logging policy.

## Choose Status Codes Deliberately

| Situation | Status |
|---|---:|
| successful read | 200 OK |
| resource created | 201 Created |
| empty or malformed client input | 400 Bad Request |
| unsupported method on known path | 405 Method Not Allowed |
| unexpected server failure | 500 Internal Server Error |

Do not return `200` with an error message in the body. Clients should not parse
English text to discover whether the operation failed.

## Handler Pseudocode

```text
FUNCTION create user handler(response, request)
    IF method is not POST
        send 405 and stop

    decode JSON into create-user input
    IF decoding fails
        send 400 and stop

    trim and validate username
    IF invalid
        send 400 and stop

    create server-owned user value
    set Content-Type to application/json
    set status 201
    encode user as response JSON
```

Every failure branch returns immediately. Otherwise the handler may continue
and attempt to send both failure and success responses.

## Worked Example: Create A Book

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
)

type Book struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

type createBookInput struct {
    Title string `json:"title"`
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input createBookInput
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    if err := decoder.Decode(&input); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    input.Title = strings.TrimSpace(input.Title)
    if input.Title == "" {
        http.Error(w, "title is required", http.StatusBadRequest)
        return
    }

    book := Book{ID: 1, Title: input.Title}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(book); err != nil {
        log.Printf("encode book response: %v", err)
    }
}
```

`DisallowUnknownFields` catches misspelled keys such as `"titel"`. Whether an
API should reject or ignore unknown fields is a contract decision.

## Your Program

Build:

- `GET /health` returning `{"status":"ok"}`
- `POST /users` accepting username and role
- server-assigned ID 1
- username validation
- JSON response with status 201
- 400 for malformed or invalid input
- 405 for unsupported methods

Storage is intentionally excluded. The goal is the JSON/HTTP contract.

## Build In Checkpoints

1. Encode a fixed struct to the response.
2. Decode a valid request body and print the typed input.
3. Return the decoded values in JSON.
4. Add server-owned ID.
5. Add malformed JSON and field validation paths.
6. Add health JSON and method checks.

Run:

```bash
gofmt -w .\main.go
go run .
```

Test from another window:

```bash
$body = @{
    username = "adam"
    role = "admin"
} | ConvertTo-Json

Invoke-WebRequest `
    -Uri http://localhost:8080/users `
    -Method Post `
    -ContentType "application/json" `
    -Body $body
```

Malformed JSON with native curl:

```bash
curl -i -X POST http://localhost:8080/users `
  -H "Content-Type: application/json" `
  --data "{broken"
```

## Test Table

| Request | Expected |
|---|---|
| valid POST JSON | 201 and JSON user |
| empty username | 400 |
| missing username | 400 |
| malformed JSON | 400 |
| wrong JSON type, such as numeric username | 400 |
| GET `/users` | 405 with `Allow: POST` |
| GET `/health` | 200 and JSON |

## Failure Experiments

1. Make struct fields lowercase and inspect the response.
2. Decode without `&input`. Read the decoder error.
3. call `WriteHeader` after `Encode`. Inspect the resulting 200 status.
4. Remove `return` after invalid JSON and observe the conflicting response path.
5. Let the client submit an ID. Explain why that violates ownership.

## Production Boundary

Real handlers should limit body size, set read/write timeouts, use a stable JSON
error format, avoid exposing internal errors, and test with `httptest`. Those
controls come after the basic contract is understood.

## You Understand This Exercise When

You can trace bytes into a struct and back, explain exported fields and tags,
separate client-owned from server-owned fields, and defend every status code.

References:

- [`encoding/json`](https://pkg.go.dev/encoding/json)
- [`net/http`](https://pkg.go.dev/net/http)
- [JSON and Go](https://go.dev/blog/json)
