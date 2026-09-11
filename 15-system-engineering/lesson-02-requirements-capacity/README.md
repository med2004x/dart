# Lesson 2: Requirements and Capacity Planning

## Objective

Translate a vague performance goal ("it should be fast," "it should handle our users") into concrete numbers — expected request rate, latency targets, data volume — using the mathematics track's queueing theory to reason about what capacity is actually required, rather than guessing.

## Prerequisites

Mathematics track's queueing lesson (math-for-engineering step — this lesson directly applies that lesson's `ρ`, `L`, `W` formulas to a real capacity-planning scenario), computer architecture Lesson 7 (storage/latency numbers — needed to reason concretely about what a "fast" response actually requires).

## Learn

**Why "it should be fast" isn't a specification.** Without a number, you can't know whether a design meets the requirement, can't compare two design options objectively, and can't tell whether a load test result is good or bad. Real capacity planning starts with translating business context into numbers: how many users, how many requests per user per unit time, what's an acceptable response time (and at what percentile — recall database-internals/statistics Lesson 1's point that p50 and p99 are very different numbers, and "fast" usually implicitly means p95 or p99, not average).

**Estimating request rate from business numbers.** If you know (or can reasonably estimate) daily active users and typical actions per user per day, you can derive requests per second — and critically, real traffic isn't uniform across the day, so a rough peak-to-average ratio (e.g. "peak traffic is roughly 3-5x the daily average, concentrated in business hours") matters for sizing capacity correctly; sizing only for the daily average will leave you under-provisioned exactly when it matters most.

**Using queueing theory to reason about required capacity, concretely.** Given an estimated arrival rate `λ` and a target average latency `W`, the math-for-engineering queueing lesson's M/M/1 formulas let you solve for the required service rate `μ` (and therefore the minimum server capacity needed) to hit that latency target — and, more importantly, that lesson's core insight (utilization `ρ` approaching 1 causes latency to blow up nonlinearly) tells you *why* simply provisioning for "average load" is dangerous: you need enough headroom that peak load doesn't push `ρ` close enough to 1 to enter the steep part of the latency curve.

**Data volume and storage growth.** Beyond request rate, capacity planning includes reasoning about data growth: if each user action creates roughly N bytes of persistent data, and you're adding M new users per month, you can project storage growth and reason about database-internals-track concerns (B+ tree index size, buffer pool sizing) scaling with data volume, not staying fixed as usage grows — a system correctly sized for today's data volume can degrade purely from data growth even with unchanged request patterns.

## Attempt

1. For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

2. Using the math-for-engineering queueing lesson's M/M/1 formulas, and your step 1 peak `λ` estimate, solve for the minimum service rate `μ` required to keep average time-in-system `W` under a target (e.g. 2 seconds) — show your work using the actual formulas, not just a rough guess.

3. Reason about utilization headroom: given your derived `μ` from step 2, compute what your peak `ρ` would actually be, and explicitly discuss — using the queueing lesson's `ρ→1` blowup — why provisioning capacity to exactly hit your target `W` at the *estimated* peak load is risky if actual peak load turns out to be even modestly higher than estimated (a very realistic possibility, since traffic estimates are rarely exactly accurate).

4. Project data growth: given an estimate of average data size created per transcription request (a reasonable guess is fine, state it explicitly) and your estimated daily volume from step 1, project total data volume at 6 months and 1 year of operation, and reason about whether that volume would still comfortably fit the kind of buffer-pool/index-size assumptions from the database-internals track, or whether it's approaching a scale where those systems would need more deliberate tuning.

## Verify

Present your full capacity-planning writeup: request-rate derivation with stated assumptions, the queueing-formula-based service-rate requirement, the utilization-headroom reasoning, and the data-growth projection — each step should show its actual arithmetic, not just a final number asserted without the derivation.

## Failure drill

Take your step 2 capacity plan (provisioned for your best-estimate peak `λ`) and recompute `W` if actual peak load turns out to be 50% higher than estimated (a realistic margin of estimation error) — using the same `μ` you provisioned for. Report the new `ρ` and the resulting `W`, and confirm, using the queueing formula directly, that this modest underestimate produces a disproportionately large latency degradation, not a modest one — directly demonstrating why capacity planning needs deliberate headroom, not just point estimates, and why "close to my estimate" and "close to my target latency" are not the same thing once you're near the steep part of the utilization curve.

## Transfer

If TARDOC's actual clinic count or transcription volume differs from this lesson's example numbers, redo the capacity-planning exercise (steps 1-3) using your real, current numbers (or your best actual estimates, if exact figures aren't readily available), and state explicitly whether your current infrastructure (a single Contabo VPS, per your project history) has headroom relative to what this analysis suggests, is roughly matched, or is already closer to the danger zone than you'd want, based on the utilization-headroom reasoning from step 3.

## Done when

You've translated a vague performance goal into concrete numbers using real (or realistically estimated) business data, you've applied the queueing formulas correctly to derive a required capacity, and you've demonstrated — using your own recomputed numbers under a 50% traffic-estimate error — why headroom beyond your best-guess estimate is a genuine engineering requirement, not just conservative padding.
