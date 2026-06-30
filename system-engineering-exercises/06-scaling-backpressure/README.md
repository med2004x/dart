# Project 06 - Scaling And Backpressure

## Goal

Measure one bottleneck and reject excess work instead of allowing unbounded
queues and memory growth.

## Starter

The server exposes `/work`. Every request sleeps for 200 ms. It currently
accepts unlimited concurrent requests.

## Checkpoints

1. Add a buffered-channel semaphore allowing five workers.
2. Return 503 immediately when all workers and queue slots are occupied.
3. Include `Retry-After`.
4. Count active, completed, and rejected requests.
5. Add `/metrics` returning the counters as JSON.
6. Run sequential and concurrent measurements.

Concurrent PowerShell:

```powershell
$jobs = 1..30 | ForEach-Object {
    Start-Job {
        curl.exe -s -o NUL -w "%{http_code}" http://localhost:8082/work
    }
}

$jobs | Receive-Job -Wait -AutoRemoveJob |
    Group-Object |
    Select-Object Name, Count
```

## Failure Drill

Replace the bounded semaphore with one goroutine per request and a growing
in-memory queue. Explain what happens when arrival rate exceeds completion rate.

## Design Decision

Compare:

- optimize the work
- increase bounded worker capacity
- add instances
- queue asynchronous jobs
- reject/rate-limit callers

Choose only after naming the measured resource limit.

## Done Means

Overload produces a controlled response and observable rejection count.

