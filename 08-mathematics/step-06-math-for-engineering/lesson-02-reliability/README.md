# Lesson 2: Reliability Modeling

## Objective

Apply probability to system reliability: compute availability for series/parallel component configurations, and connect this directly to SLA math and redundancy design decisions.

## Prerequisites

Lesson 1 of probability (independence — reliability calculations for independent components rely on it directly), Lesson 1 of this step (queueing — both lessons in this final step apply probability to concrete infrastructure decisions).

## Learn

**Availability.** The fraction of time a system is operational, typically expressed as a percentage or in "nines" — 99% ("two nines") allows about 3.65 days of downtime per year; 99.9% ("three nines") allows about 8.76 hours per year; 99.99% ("four nines") allows about 52.6 minutes per year. Each additional nine is an order-of-magnitude reduction in allowed downtime, which is why going from three nines to four nines is a much harder engineering problem than the small-looking percentage difference (99.9% → 99.99%) suggests.

**Series systems** (both components must work for the system to work — e.g. a request that must pass through both a load balancer and a database, in sequence, with no fallback for either): if component A has availability `a` and component B has availability `b`, and they fail independently, the system's availability is `a·b` — always less than or equal to the worse of the two individual components. This is a direct consequence of the independence multiplication rule from probability Lesson 1, and it means chaining dependent services in series without redundancy makes overall reliability strictly worse than any single link, no matter how reliable each individual link is.

**Parallel (redundant) systems** (the system works if *at least one* component works — e.g. two independent database replicas, either of which can serve a read): the system *fails* only if both fail, so `P(system fails) = P(A fails)·P(B fails) = (1-a)(1-b)`, and system availability is `1 - (1-a)(1-b)`. This is always greater than or equal to the better of the two individual components — redundancy is the concrete mechanism that lets a system exceed any single component's reliability.

Worked example: two independent components each with 99% availability (`a=b=0.99`).

*In series:* `0.99 × 0.99 = 0.9801` — 98.01%, worse than either component alone.

*In parallel:* `1 - (1-0.99)(1-0.99) = 1 - 0.01×0.01 = 1 - 0.0001 = 0.9999` — 99.99%, dramatically better than either component alone, because both would need to fail simultaneously.

This is the quantitative justification for redundancy as a reliability strategy, and it also quantifies exactly why a system with many components chained in series (a long dependency chain with no fallback at any link) is fragile — the availabilities multiply down, not just add up as risk intuitively might suggest.

## Attempt

1. A request path passes through 4 components in series (load balancer, API server, database, cache), each independently available 99.9% of the time. Compute the overall series availability. State the resulting downtime in hours/year using the "nines" reference table in Learn (or interpolate).

2. Take the database component from step 1 and make it redundant: two independent database replicas in parallel, each still at 99.9% individually, serving as one logical "database" component in the series chain. Recompute the overall series availability from step 1 with this one component upgraded to its parallel-redundant availability, and compare the resulting overall system availability to step 1's.

3. Determine how many *independent* 99% components you'd need in parallel to achieve at least 99.999% ("five nines") combined availability. Solve `1 - (0.01)ⁿ ≥ 0.99999` for the smallest integer `n` (note: `(1-a)ⁿ = (0.01)ⁿ` for n components each with 1% failure probability, generalizing the two-component formula).

4. Implement a small reliability calculator in Go or Python: functions `seriesAvailability(availabilities []float64) float64` and `parallelAvailability(availabilities []float64) float64` that generalize the two-component formulas to n components, and verify them against your hand calculations from steps 1-3.

## Verify

Your code's output for steps 1-3 must match your hand-computed values. Report the actual downtime-per-year figures (not just percentages) for at least the step 1 and step 2 results, since "99.9% vs 99.6%" is much less viscerally clear than "8.76 hours/year vs 35 hours/year."

## Failure drill

Take the 4-component series system from step 1 and make the *weakest* possible design choice: instead of each component being independent, assume the database and cache components share a single point of failure (e.g. they're both hosted on the same physical server, so if that server fails, both fail together — meaning they are not actually independent, violating the assumption the series formula relies on). Explain in your own words why applying the independent-series formula to this scenario would give an overly optimistic (too high) availability estimate, and what the real-world implication is: shared infrastructure (a single VPS host, a single availability zone, a single network path) silently breaks the independence assumption that makes redundancy math work, which is why real high-availability designs deliberately spread redundant components across independent failure domains (different hosts, different availability zones), not just multiple instances on the same underlying infrastructure.

## Transfer

If TARDOC's deployment (a single Contabo VPS, per your project history) currently has any single points of failure — a single database instance, a single VPS host, a single region — use this lesson's series/parallel formulas to estimate, even roughly, what availability you're implicitly targeting today, and what specific redundancy (a second replica, a second host) would move the needle most given the series/parallel math, rather than guessing at what "more reliable" would require.

## Done when

You can compute series and parallel availability for a multi-component system by hand and in code, you can solve for the number of redundant components needed to hit a reliability target, and you can explain — using the shared-infrastructure failure drill as the concrete case — why redundancy only delivers its theoretical benefit when components genuinely fail independently.
