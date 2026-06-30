# Project 01 - Request Boundaries

## Goal

Build an HTTP server where one request can be traced from client to handler and
back. Learn to identify the last successful boundary before proposing a fix.

## Starter

`main.go` provides:

- `GET /health`
- `GET /work`
- one request log before the handler
- a five-second server header timeout

Run:

```powershell
go run .
```

Second terminal:

```powershell
curl.exe -i http://localhost:8080/health
curl.exe -i http://localhost:8080/work
```

## Checkpoints

1. Add a generated request ID.
2. Return it as `X-Request-ID`.
3. Log request start and completion.
4. Include method, path, status, and duration.
5. Add `/fail?at=handler` returning 500.
6. Add `/fail?at=dependency` simulating a dependency error.

## Request Trace

Complete with real function names:

```text
curl
-> TCP localhost:8080
-> http.Server
-> ServeMux
-> middleware
-> handler
-> response writer
-> curl
```

## Failure Drills

1. Stop the process and call `/health`.
2. Call an unknown route.
3. Occupy port 8080 with another process.
4. Trigger each `/fail` mode.

For each, record:

- client output
- server output
- last confirmed boundary
- root cause
- correct owner of the fix

## Done Means

You can distinguish connection failure, routing failure, handler failure, and
dependency failure from evidence.

