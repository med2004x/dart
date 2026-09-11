# Lesson 1: Queueing Theory

## Objective

Apply probability (Poisson processes, exponential distributions) to model request queues, and derive the practical rule-of-thumb behind why systems degrade sharply as utilization approaches 100% — directly relevant to capacity planning for TARDOC or Mahall under real traffic.

## Prerequisites

Lesson 2 of probability (Poisson and exponential distributions — this lesson's whole model is built from them).

## Learn

**The M/M/1 queue** is the simplest useful queueing model: Poisson arrivals (rate `λ`), exponential service times (rate `μ`, so average service time is `1/μ`), 1 server, infinite queue capacity. Despite its simplicity, it captures the core nonlinear behavior every real queueing system exhibits.

**Utilization** `ρ = λ/μ` — the fraction of time the server is busy. For the queue to be stable (not grow unboundedly), you need `ρ < 1` (arrivals must be slower than service, on average).

**Average number in system** (waiting + being served): `L = ρ/(1-ρ)`. **Average time in system** (Little's Law, a genuinely general result not specific to M/M/1): `W = L/λ`, which for M/M/1 works out to `W = 1/(μ-λ) = (1/μ)/(1-ρ)`.

**Why this formula is the single most important intuition in this lesson.** As `ρ → 1` (utilization approaches 100%), the denominator `(1-ρ) → 0`, so `L` and `W` blow up toward infinity. This isn't a linear degradation — a queue at 90% utilization has roughly 9x the average queue length of one at 50% utilization (`0.9/0.1=9` vs `0.5/0.5=1`), and at 99% utilization it's roughly 99x. This is the rigorous justification behind the common operational wisdom "don't run your systems at high utilization" — it's not caution for its own sake, it's the direct, provable consequence of this formula's shape near `ρ=1`.

**Little's Law**, `L = λW`, deserves separate mention because it holds for *any* stable queueing system, regardless of arrival distribution, service distribution, or number of servers — a genuinely general, assumption-light result. It says: average number of items in a system equals the arrival rate times the average time each item spends in the system. This lets you estimate one of the three quantities (`L`, `λ`, `W`) from the other two even when you don't know or can't assume the detailed statistical structure of your actual system.

## Attempt

1. A server processes requests at an average rate `μ = 100/sec` (exponential service time). Requests arrive at `λ = 60/sec` (Poisson). Compute `ρ`, `L` (average number in system), and `W` (average time in system) using the M/M/1 formulas.

2. Recompute `L` and `W` for the same `μ=100` but with `λ` increased to 80, then 95, then 99. Tabulate `ρ, L, W` for all four `λ` values (60, 80, 95, 99) and observe explicitly how sharply `L` and `W` grow as `ρ` approaches 1 — compute the ratio of `L` at `λ=99` to `L` at `λ=60` to make the nonlinearity concrete with actual numbers, not just "it gets worse."

3. Implement a discrete-event simulation of an M/M/1 queue in Go or Python: generate Poisson-distributed arrival times (equivalently, exponential inter-arrival times) and exponential service times for a given `λ` and `μ`, simulate the queue for, say, 100,000 arrivals, and empirically measure the average time in system. Compare your simulated `W` against the theoretical formula for at least two different `ρ` values (e.g. `ρ=0.6` and `ρ=0.9`).

4. Apply Little's Law directly: if you know (or can estimate) that a system processes 50 items/second on average (`λ=50`) and observe that the average item spends 0.2 seconds in the system (`W=0.2`), compute `L` (average number of items in the system at any time) using `L=λW`, without needing to know anything about the arrival or service distributions.

## Verify

Report your simulated `W` from step 3 alongside the theoretical `W` for both `ρ` values — they should agree closely (within a few percent, given enough simulated arrivals) at `ρ=0.6`, and should still agree reasonably at `ρ=0.9` though with more variance in a finite simulation run, since the theoretical values themselves are larger and more sensitive near high utilization.

## Failure drill

Push your simulation from step 3 to `ρ=0.99` (e.g. `λ=99, μ=100`) and observe both the theoretical `L`/`W` and your simulated results becoming very large and, in the simulation, noticeably unstable/high-variance across repeated runs even with 100,000 arrivals. Explain why this happens: near `ρ=1`, the system is only barely stable, so a temporary run of faster-than-average arrivals (which will happen by chance in any finite Poisson simulation) can create a large queue backlog that takes a long time to clear, and this sensitivity is not a simulation bug, it's the real, provable behavior of queues operating near saturation.

## Transfer

If TARDOC's Celery workers, or any request-handling component in Mahall, have a known or estimable processing rate and observed arrival rate, use the M/M/1 formulas (as a rough first approximation, acknowledging real systems often aren't exactly Poisson/exponential) to estimate what utilization you're currently running at, and state — using this lesson's `ρ→1` blowup as justification — what utilization level you'd want to stay under for acceptable latency, rather than just picking a round number like "80%" without the underlying reason.

## Done when

You can compute `ρ`, `L`, and `W` for an M/M/1 queue from `λ` and `μ`, you've confirmed via simulation that theory and empirical results agree, and you can explain — using your own step 2 ratio calculation as evidence — why queueing systems degrade nonlinearly as utilization approaches 100%, rather than just citing it as received wisdom.
