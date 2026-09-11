# Capstone 1: Production Network Service

## Objective

Build a complete, production-shaped network service integrating correct networking, persistence, observability, and failure handling — the first of four terminal capstones drawing on nearly every track in this curriculum, proving the pieces compose into something a real team could actually operate.

## Prerequisites

This capstone assumes completion of: networking (sockets through HTTP), database-internals or at minimum postgresql-engineering (persistence), API engineering (contract design, idempotency, auth), system-engineering (request boundaries, timeouts, reliability, observability), and production-engineering (SLOs, incident response). It is deliberately positioned after all of these, not before.

## Learn

There is no new material in this capstone — its entire purpose is integration at a scale beyond any single track's own capstone lesson. Where individual track capstones (networking Lesson 8, database-internals Lesson 7, API engineering Lesson 15, system-engineering Lesson 14) each integrated one track's concepts, this capstone integrates *across* tracks: a real service's networking layer, persistence layer, API contract, and operational tooling all need to work together correctly, which is precisely where cross-track integration gaps — invisible when each track is tested in isolation — tend to surface.

**What "production-shaped" specifically means here, as a concrete bar.** Not a toy demo that works once in a controlled demonstration, but a service you could hand to another engineer with reasonable confidence they could operate it: they'd know how to deploy it (production Lesson 1), how to tell if it's healthy (API engineering Lesson 14, production Lesson 2), what happens when a dependency fails (system-engineering Lesson 8), and what its actual, measured capacity limits are (production Lesson 6) — not assumed, tested.

## Attempt

Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

1. **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.
2. **Correct networking fundamentals** (networking Lessons 1-4): proper framing, correct HTTP semantics, connection reuse.
3. **Real persistence** (postgresql-engineering or database-internals): actual database-backed storage, not an in-memory stand-in, with at least one nontrivial query pattern using proper indexing.
4. **API engineering's full toolkit, integrated** (API engineering Lessons 5-14): validation, idempotency on writes, optimistic concurrency, authentication/authorization with resource-level checks, rate limiting, at least one async operation, and correctly distinguished liveness/readiness endpoints.
5. **System-engineering's reliability layer** (system-engineering Lessons 1, 5, 8): explicit request-boundary decisions, tuned timeouts/retries on any downstream calls, and at least one circuit breaker or graceful-degradation path.
6. **Observability from day one** (system-engineering Lesson 9, production Lesson 2): structured logs, metrics, and tracing sufficient to diagnose an incident purely from telemetry.
7. **A documented SLO** (production Lesson 3): a real, business-justified target with calculated error budget.

## Verify

1. **Load test it** (production Lesson 6): a real, open-loop load test finding your actual capacity and bottleneck, compared against a theoretical capacity estimate (system-engineering Lesson 2).
2. **Diagnose one induced incident** (production Lesson 5): inject a realistic failure, diagnose it blind using only your telemetry, and produce a real postmortem with trackable corrective actions.
3. **Security review** (security-engineering, system-engineering Lesson 10): a trust-boundary audit and at least one demonstrated defense-in-depth control.

Present all of the above as real artifacts — actual load test numbers, an actual incident timeline and postmortem, an actual security review writeup — not descriptions of what you'd do.

## Failure drill

Take your service's most critical failure-handling mechanism (your circuit breaker, or your idempotency-key implementation) and deliberately break it in a way that wouldn't be caught by a test exercising only that mechanism in isolation — for example, break it specifically at the point where it interacts with your observability layer (does a tripped circuit breaker actually get logged/alerted, or does it fail silently from an operator's perspective even though it's functioning correctly from the client's?). Find this gap through the same telemetry-only diagnosis discipline as your induced incident, and fix it. Document why this specific class of gap — correct in isolation, incomplete in integration — is exactly what a capstone at this scope is meant to surface, and why no single track's own lesson-level testing would have caught it.

## Transfer

Compare your capstone's actual operational maturity against TARDOC's real, current production posture (per your project history). Identify at least 3 specific things your capstone does that TARDOC's actual deployment currently doesn't, and assess honestly, for each, whether the gap represents a real, worthwhile near-term investment for TARDOC given its actual scale and risk profile, or whether it's premature relative to TARDOC's current stage.

## Done when

Your service is deployed (or deployable) with a complete, integrated feature set spanning networking through observability, you have real load-test numbers and a real diagnosed-and-postmortem'd synthetic incident, you've found and fixed at least one integration-specific gap invisible to any single component's isolated testing, and you can honestly assess how your capstone's operational maturity compares to a real system you actually operate.
