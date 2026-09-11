# Exercise 13 - Middleware And Authentication

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

Middleware composes HTTP handlers. It applies behavior around a handler without
copying that behavior into every endpoint.

Request flow:

```text
client
  -> logging middleware
    -> authentication middleware
      -> private handler
    <- authentication result
  <- logging result
<- response
```

The order of wrapping determines the order of execution.

## Beginner Bridge: Middleware Is A Wrapper

Middleware is code that runs around a handler.

Think of a handler as the core action:

```text
private page
```

Authentication middleware wraps it:

```text
check API key -> private page
```

Logging middleware can wrap that:

```text
log request -> check API key -> private page
```

If auth fails, it writes `401` and stops. Stopping means it does not call the
next handler.

Similar wrappers:

| Middleware | Runs before handler? | May stop request? |
|---|---|---|
| logging | yes | usually no |
| authentication | yes | yes |
| rate limiting | yes | yes |
| compression | often after | no |

The central question in every middleware is: "Do I call `next.ServeHTTP`, and
under which conditions?"

## Handler And HandlerFunc

The core interface:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

A function with the matching parameter shape can be adapted:

```go
http.HandlerFunc(myFunction)
```

Middleware accepts a handler and returns another handler:

```go
func middleware(next http.Handler) http.Handler
```

The returned handler can do work before calling `next.ServeHTTP`, after it, or
both.

## Build One Middleware Mechanically

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("method=%s path=%s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

Read from outside inward:

1. `loggingMiddleware` receives the next handler.
2. It creates a new handler function.
3. The new function logs the request.
4. It delegates to `next`.

If it never calls `next.ServeHTTP`, the inner handler never runs.

## Authentication Short-Circuits

Pseudocode:

```text
FUNCTION API key middleware(next)
    RETURN handler that:
        read X-API-Key header
        IF key is missing or wrong
            send 401
            RETURN without calling next
        call next handler
```

Authentication answers "who or what presented credentials?" Authorization
answers "may that identity perform this action?" This exercise performs only a
simple authentication check.

## Wrap Only Protected Routes

```go
publicHandler := http.HandlerFunc(public)
privateHandler := authMiddleware(http.HandlerFunc(private))

mux.Handle("/public", publicHandler)
mux.Handle("/private", privateHandler)
```

Do not wrap the whole mux with auth when health and public routes must remain
accessible.

Logging can wrap the whole mux:

```go
root := loggingMiddleware(mux)
```

Then pass `root` as the server's handler.

## Middleware Order

These are not equivalent:

```go
loggingMiddleware(authMiddleware(private))
authMiddleware(loggingMiddleware(private))
```

In the first, rejected requests are logged. In the second, auth can reject before
logging runs. Decide from the intended policy, not aesthetics.

## Worked Example: Required Client Version

This analogous middleware checks a shared request requirement:

```go
func requireClientVersion(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-Client-Version") != "1" {
            http.Error(w, "unsupported client version", http.StatusBadRequest)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func dashboard(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "dashboard")
}

func example() http.Handler {
    return requireClientVersion(http.HandlerFunc(dashboard))
}
```

The middleware owns a cross-route transport policy. The dashboard owns its
endpoint response.

## Your Program

Build:

- `GET /public` without authentication
- `GET /private` protected by `X-API-Key: secret123`
- request logging for every route
- 401 for missing or invalid keys
- method checks for each route

The hardcoded key is acceptable only for this local exercise. Real secrets come
from protected configuration, not source control.

## Build In Checkpoints

1. Implement and test both handlers without middleware.
2. Write logging middleware and wrap the whole mux.
3. Write auth middleware but do not attach it yet.
4. Wrap only the private handler.
5. Test missing, wrong, and correct headers.
6. Reverse middleware order and explain the log difference.

```bash
gofmt -w .\main.go
go run .
```

Second window:

```bash
curl -i http://localhost:8080/public
curl -i http://localhost:8080/private
curl -i http://localhost:8080/private -H "X-API-Key: wrong"
curl -i http://localhost:8080/private -H "X-API-Key: secret123"
```

## Test Matrix

| Route | Header | Expected | Logged? |
|---|---|---:|---|
| public | none | 200 | yes |
| private | none | 401 | yes |
| private | wrong key | 401 | yes |
| private | valid key | 200 | yes |
| missing route | any | 404 | yes |

## Failure Experiments

1. Omit `return` after 401. Confirm whether the private handler still runs.
2. Forget `next.ServeHTTP` on success. Explain the empty response.
3. Wrap the whole mux in auth. Show why `/public` is no longer public.
4. Log the API key. Explain why credentials must never enter logs.
5. Put auth checks directly in both handlers, then change the header name.
   Observe the duplication risk.

## Production Boundary

An API key sent over plain HTTP can be intercepted; production requires TLS.
Hardcoded shared keys have poor rotation and identity properties. Production
authentication also needs secret management, rate limiting, audit policy, and
often authorization checks.

Plain string logs are acceptable for this focused exercise. Production logs
should be structured and include a request ID without credentials or PII.

## You Understand This Exercise When

You can expand the wrapper execution order by hand, explain every path that does
or does not call `next`, and attach policy only to the routes that require it.

References:

- [`http.Handler`](https://pkg.go.dev/net/http#Handler)
- [`http.HandlerFunc`](https://pkg.go.dev/net/http#HandlerFunc)
- [`http.Request.Header`](https://pkg.go.dev/net/http#Request)
