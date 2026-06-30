# Project 09 - Observability

## Goal

Make one request traceable through structured logs and measurable through rate,
errors, latency, and saturation.

## Required Completion Log

```json
{
  "timestamp": "2026-06-29T12:00:00Z",
  "level": "info",
  "message": "request completed",
  "requestId": "req-123",
  "method": "GET",
  "path": "/health",
  "status": 200,
  "durationMs": 4
}
```

## Checkpoints

1. Generate or validate one request ID.
2. Return it in `X-Request-ID`.
3. Capture actual response status.
4. Log one JSON completion event.
5. Add request/error counters and a latency histogram.
6. Add active-request saturation.
7. Expose a local `/metrics` JSON endpoint.

## Analyze Logs

```powershell
go run . 2> .\server.log
```

Then:

```powershell
$events = Get-Content .\server.log |
    ForEach-Object { $_ | ConvertFrom-Json }

$events |
    Group-Object status |
    Select-Object Name, Count
```

## Failure Drills

Trigger:

- 404
- 405
- recovered panic returning 500
- slow request
- invalid incoming request ID

Never log API keys, tokens, passwords, or request bodies containing personal
data.

## Done Means

A client response ID leads to exactly one completion event with status and
duration, and the aggregate signals show the incident.

