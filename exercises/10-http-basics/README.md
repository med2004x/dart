# Exercise 10 - HTTP Basics

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

An HTTP server is a long-running program that:

1. listens on a network address
2. accepts a request
3. chooses a handler
4. lets the handler write a response
5. repeats for later requests

This exercise is about the request/response boundary, not JSON or storage yet.

## Beginner Bridge: HTTP Is A Conversation

HTTP is a request followed by a response.

The client asks:

```text
method + path + headers + optional body
```

The server answers:

```text
status code + headers + optional body
```

In Go, the handler receives both sides of that conversation:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

- `r` is what the client sent.
- `w` is how your code writes the response.

Similar routes:

| Path | Meaning | Good first response |
|---|---|---|
| `/health` | is the server alive? | `200 ok` |
| `/version` | what version is running? | `200 version 1` |
| `/hello` | simple demo route | `200 hello` |

Do not mix in JSON, databases, auth, or CRUD yet. If you cannot explain one
plain text request and response, the later API exercises will feel random.

## One Request, Step By Step

Client request:

```http
GET /health HTTP/1.1
Host: localhost:8080
```

Server response:

```http
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

ok
```

Important parts:

- method: `GET`
- path: `/health`
- status: `200`
- headers: response metadata
- body: `ok`

The path identifies a resource. The method describes the requested operation.

## Handler Parameters

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
}
```

`r` contains client input:

```go
r.Method
r.URL.Path
r.Header
r.Body
```

`w` sends output:

```go
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
w.WriteHeader(http.StatusOK)
fmt.Fprintln(w, "ok")
```

Headers and status must be written before the body. Writing a body first usually
commits an implicit `200 OK`.

## Routing

A `ServeMux` maps request paths to handlers:

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", healthHandler)
mux.HandleFunc("/hello", helloHandler)
```

The server passes each matching request to the selected handler.

Using an explicit mux is easier to reason about than adding routes to the global
default mux.

## Starting The Server

```go
server := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
}

if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

`ListenAndServe` blocks while the server runs. Code after it does not execute
until the server stops or startup fails.

The read-header timeout prevents a client from holding a connection forever
while slowly sending headers. External operations need time bounds.

## Method Checks

A route is not only a path. If `/health` supports `GET`, reject other methods:

```text
IF request method is not GET
    include Allow: GET header
    send 405 Method Not Allowed
    stop handler
```

`404 Not Found` means no matching resource. `405 Method Not Allowed` means the
resource exists but does not support that method.

## Worked Example: Status And Version Server

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.Header().Set("Allow", http.MethodGet)
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintln(w, "version 1")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/version", versionHandler)

    server := &http.Server{
        Addr:              ":8080",
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Println("listening on http://localhost:8080")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
```

This is a complete server for one analogous route. Your exercise adds health,
hello, and explicit not-found behavior.

## Your Program

Build a server on port 8080:

- `GET /health` returns `200` and `ok`
- `GET /hello` returns `200` and `hello from go`
- unsupported methods return `405`
- unknown paths return `404`

Use an explicit `http.ServeMux` and a configured `http.Server`.

## Build In Checkpoints

1. Start a server with one inline handler.
2. Move the handler to a named function.
3. Add a second route.
4. Add method checks.
5. Test an unknown path.
6. Add a read-header timeout and startup logging.

Start the server:

```powershell
gofmt -w .\main.go
go run .
```

In a second PowerShell window:

```powershell
Invoke-WebRequest -Uri http://localhost:8080/health
Invoke-WebRequest -Uri http://localhost:8080/hello
Invoke-WebRequest -Uri http://localhost:8080/missing -SkipHttpErrorCheck
```

`-SkipHttpErrorCheck` exists in newer PowerShell versions. With Windows
PowerShell 5.1, use `curl.exe` to inspect expected error responses:

```powershell
curl.exe -i http://localhost:8080/missing
curl.exe -i -X POST http://localhost:8080/health
```

Stop the server with `Ctrl+C`.

## Test Table

| Method | Path | Expected |
|---|---|---|
| GET | `/health` | 200, `ok` |
| GET | `/hello` | 200, greeting |
| POST | `/health` | 405 |
| GET | `/missing` | 404 |
| GET | `/health?full=true` | define whether query changes behavior |

## Failure Experiments

1. Write the body before `WriteHeader(404)`. Inspect the actual status.
2. Start a second server on port 8080. Read the address-in-use error.
3. Remove the `return` after `http.Error`. Observe whether extra output is
   written.
4. Use a path without a leading slash in route registration. Read the panic.

## Production Boundary

This server still lacks graceful shutdown, request IDs, structured logging,
rate limiting, and full timeout configuration. The exercise is correct for
learning handlers, but it is not production-ready.

## You Understand This Exercise When

You can trace a request from socket acceptance to route selection and response,
distinguish 404 from 405, and explain when response status becomes committed.

References:

- [`net/http`](https://pkg.go.dev/net/http)
- [`http.Server`](https://pkg.go.dev/net/http#Server)
- [Writing Web Applications](https://go.dev/doc/articles/wiki/)
