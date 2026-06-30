# Project 08 - Reliability And Recovery

## Goal

Implement separate liveness/readiness behavior and drain in-flight work during
shutdown.

## Starter

The server provides:

- `/health`
- `/ready`
- `/slow`

It currently exits through the default server path without graceful shutdown.

## Checkpoints

1. Track readiness with an atomic boolean.
2. Register `os.Interrupt` with `signal.NotifyContext`.
3. Mark readiness false before shutdown.
4. call `Server.Shutdown` with a five-second deadline.
5. Let the two-second `/slow` request finish.
6. Log shutdown start, drain result, and total duration.
7. Add a fake dependency readiness flag.

## Manual Test

Start `/slow` from one terminal:

```powershell
curl.exe -i http://localhost:8083/slow
```

Press `Ctrl+C` in the server terminal before it completes.

## Failure Drill

Replace graceful shutdown with immediate `os.Exit(1)`. Repeat the in-flight
request and record the client behavior.

Then mark liveness unhealthy whenever the fake dependency fails. Explain why
restarting this process cannot repair the dependency.

## Done Means

The service stops receiving new traffic and drains bounded in-flight work before
exit.

