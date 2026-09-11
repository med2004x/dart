# Lesson 8: Reliability and Recovery

## Objective

Design a system component with an explicit failure budget and recovery strategy — not just "handle errors," but reason quantitatively about acceptable failure rates and build recovery mechanisms (retries, circuit breakers, graceful degradation) matched to the actual reliability target.

## Prerequisites

System-engineering Lesson 5 (timeouts/retries — this lesson's circuit breaker specifically builds on that lesson's retry-storm discussion), production-engineering track (SLO concepts, if covered before this — otherwise this lesson introduces the core idea directly).

## Learn

**Availability targets, and what they actually mean operationally.** Computer architecture Lesson 8 (reliability modeling, math-for-engineering) established the "nines" framework: 99.9% availability allows roughly 8.76 hours of downtime per year; 99.99% allows roughly 52.6 minutes. Choosing a target isn't just picking an impressive-sounding number — it directly determines what engineering investment is justified (per-component redundancy, active-active failover, extensive chaos testing) versus overkill for a given system's actual business requirements. A target chosen without reference to actual business impact of downtime is arbitrary and can lead to either wasted engineering effort (over-engineering reliability nobody needs) or under-investment relative to real risk.

**Circuit breakers: a more sophisticated response to a failing dependency than plain retries.** System-engineering Lesson 5 covered timeouts and backoff for individual retries. A circuit breaker goes further: it tracks a dependency's recent failure rate, and once failures exceed a threshold, it "opens" — for a period, calls to that dependency fail immediately (without even attempting the network call) rather than continuing to try and wait for timeouts, giving the struggling dependency room to recover instead of continuing to pile on load (directly addressing Lesson 5's cascading-failure concern, but proactively rather than just via a well-tuned timeout on each individual call). After a cooldown period, the circuit breaker allows a small number of "trial" requests through (a "half-open" state) — if they succeed, it closes again (normal operation resumes); if they still fail, it stays open longer.

**Graceful degradation: designing for partial functionality under partial failure, rather than "everything or nothing."** If a search feature depends on a recommendation service that's currently down, a well-designed system can serve basic, unpersonalized results (a degraded but still useful response) rather than failing the entire search request because one non-essential dependency is unavailable — this directly connects back to system-engineering Lesson 1's request-boundary lesson: correctly identifying which dependencies are truly essential (their failure should fail the request) versus which are enhancements (their failure should degrade, not fail, the response) is exactly the design judgment that makes graceful degradation possible.

**Recovery time objective (RTO) and recovery point objective (RPO), the two numbers that define what "recovery" actually needs to achieve.** RTO: how long can the system be down before recovery must be complete? RPO: how much data loss (measured in time — "we can lose up to N minutes of data") is acceptable in a disaster scenario? These numbers directly drive concrete engineering decisions — a tight RPO requires frequent backups or synchronous replication (distributed Lesson 4's tradeoff, revisited here with a business-driven number attached); a tight RTO requires automated, tested failover rather than a manual, ad hoc recovery process.

## Attempt

1. For a real component in TARDOC or Mahall, propose a concrete availability target (e.g. 99.5%) based on actual reasoning about business impact (what does an hour of downtime actually cost or disrupt, given your real business context), not an arbitrary number — write out your reasoning explicitly.

2. Implement a circuit breaker wrapping a call to a simulated flaky dependency (one you can control to fail on demand for testing). Configure a failure-rate threshold, an open-state cooldown, and half-open trial behavior per Learn. Test it through a full cycle: healthy calls succeed normally (closed state), induced failures trip the breaker open (subsequent calls fail immediately, without even attempting the real call — verify this by confirming the underlying flaky dependency isn't even invoked while open), and after the cooldown, a successful trial call closes the breaker again.

3. Implement graceful degradation for one real or plausible feature: identify a genuinely non-essential dependency (per system-engineering Lesson 1's boundary framework) and implement a fallback path that serves a reduced-but-functional response when that dependency is unavailable, rather than failing the whole request. Test it by simulating the dependency's failure and confirming the request still succeeds with the degraded response, rather than erroring out entirely.

4. Propose explicit RTO and RPO numbers for a real component (e.g. TARDOC's database), and reason about whether your current or a plausible backup/replication strategy would actually meet those numbers — using database-internals Lesson 6's recovery/WAL concepts and distributed Lesson 4's replication tradeoffs to ground the reasoning concretely, not just asserting a number without connecting it to an actual mechanism capable of achieving it.

## Verify

For step 2, show your circuit breaker's actual state transitions (closed → open → half-open → closed) logged across a real test run with controlled failure injection, confirming the underlying dependency call is genuinely skipped while the breaker is open (not just that the caller sees a fast failure, but that you've confirmed the actual downstream call didn't happen). For step 3, show the actual degraded response your fallback path produces, confirming the overall request still succeeds despite the simulated dependency failure.

## Failure drill

Configure your step 2 circuit breaker's failure-rate threshold far too sensitively (e.g. opening after just 1 failure, when the dependency's normal, healthy behavior includes occasional transient failures within an acceptable range) and rerun a test with realistic, occasional transient failures mixed into otherwise-successful calls. Confirm the breaker now trips open unnecessarily, blocking legitimate calls to a dependency that's actually healthy overall, purely because of an overly sensitive threshold — directly analogous to Lesson 5's too-short-timeout failure mode, now applied to circuit breaker tuning specifically. Explain why circuit breaker thresholds, like timeouts, need to be calibrated against the dependency's actual, measured normal failure rate, not set to their most conservative-sounding value by default.

## Transfer

If TARDOC's or Mahall's actual availability, given their current single-VPS deployment (per your project history), is realistically closer to 99% or 99.9% than to 99.99%, describe honestly what that implies about acceptable annual downtime, and whether your step 1 proposed target for a real component is actually consistent with what the current infrastructure can realistically deliver — if there's a gap between your aspirational target and current realistic capability, state explicitly what specific infrastructure investment (redundancy, automated failover, a second availability zone) would be required to close it, rather than treating the target as achieved just because it was stated.

## Done when

You've proposed and justified a concrete, business-reasoned availability target for a real component, you've implemented and tested a working circuit breaker through its full state cycle with verified suppression of the underlying failing call while open, you've implemented and tested graceful degradation for a genuinely non-essential dependency, and you've honestly reasoned about whether your current or proposed infrastructure can actually meet stated RTO/RPO numbers rather than treating them as aspirational statements disconnected from real mechanism.
