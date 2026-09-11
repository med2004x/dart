# Lesson 14: Verification and Operational Endpoints

## Objective

Build the operational endpoints (health checks, readiness checks) real infrastructure depends on to make correct decisions about your service — and understand precisely why "liveness" and "readiness" are different questions, not the same check under two names.

## Prerequisites

Lesson 4 (HTTP semantics — health endpoints are ordinary HTTP endpoints using the same status-code conventions), OS Lesson 1/2 (process states — a liveness check is fundamentally asking "is this process's main loop still functioning," directly related to the process-state concepts covered there).

## Learn

**Why health checks matter beyond "nice to have monitoring."** Load balancers (networking Lesson 6) use health checks to decide whether to route traffic to a given instance. Container orchestrators (referenced in OS Lesson 2's `SIGTERM` discussion) use them to decide whether to restart a container. Deployment systems use them to decide whether a new deployment is safe to promote or should be rolled back. A missing or poorly-designed health check doesn't just mean worse observability — it means real infrastructure automation is making decisions blind, or worse, making *wrong* decisions based on a check that doesn't actually reflect the service's real condition.

**Liveness vs. readiness, the critical distinction.** A **liveness** check answers "is this process fundamentally alive and functioning" — typically extremely minimal (e.g. "can the HTTP server respond at all"), and a failure here usually means "restart this instance," since something has gone fundamentally wrong (deadlock, crash-loop, unresponsive process). A **readiness** check answers "is this instance currently able to correctly serve traffic" — which can be false even while the process is perfectly alive (e.g. it hasn't finished starting up yet, or a required downstream dependency like the database is currently unreachable) — and a readiness failure should mean "stop routing traffic here, but don't necessarily restart it," since restarting a process that's failing readiness only because its database is temporarily down wouldn't fix anything and would just cause unnecessary restart churn.

**Why conflating them causes real operational problems.** If your health check only checks liveness (the process can respond) but is used by infrastructure as a readiness signal, a service instance that's alive but has lost its database connection will keep receiving traffic it cannot actually serve correctly, since nothing is telling the load balancer to stop routing to it. Conversely, if a check is overly strict and used as a liveness signal (e.g. failing liveness whenever a downstream dependency is briefly unavailable), infrastructure may restart perfectly healthy process instances repeatedly for a problem restarting them cannot fix, wasting resources and potentially making an already-degraded situation worse through unnecessary restart churn.

**What a good readiness check actually verifies.** Beyond "can I respond at all," a meaningful readiness check should verify the specific dependencies the service actually needs to function correctly — a database connection (a lightweight query, not a full health audit), critical downstream API reachability if the service genuinely cannot function without it, and so on — but should avoid becoming so heavyweight or slow that the check itself becomes a performance or reliability liability (a readiness check that takes 10 seconds to run, or that itself puts meaningful load on the database, is a design mistake in the opposite direction).

## Attempt

1. Implement a `/health/live` endpoint that performs the minimal possible check — confirm the HTTP server can respond at all, with no dependency checks — returning `200 OK` if the process is fundamentally functioning.

2. Implement a `/health/ready` endpoint that additionally checks a real dependency (e.g. a lightweight database ping/query) and returns `200 OK` only if both the process is alive *and* the dependency is reachable, returning `503 Service Unavailable` otherwise with a body indicating which specific dependency check failed.

3. Test the distinction directly: with your service running normally, confirm both endpoints return 200. Then simulate the database becoming unreachable (stop it, or block network access to it) while keeping your service process itself running, and confirm `/health/live` still returns 200 (the process itself is fine) while `/health/ready` correctly returns 503 (it cannot currently serve traffic correctly) — this is the concrete, hands-on demonstration of why these need to be separate endpoints, not one combined check.

4. Simulate a scenario that should fail liveness specifically (not just readiness) — e.g. deliberately induce a deadlock in your service (reusing OS Lesson 7's deadlock construction, applied to a code path your liveness check itself would traverse or that would block the entire process from responding at all) — and confirm `/health/live` now correctly fails too, distinguishing this case from step 3's readiness-only failure.

## Verify

Show the actual response (status code and body) from both endpoints in three states: fully healthy, database-unreachable-only (readiness fails, liveness passes), and process-deadlocked (both fail, or liveness specifically fails/times out) — three clearly distinguished states with real, observed responses for each.

## Failure drill

Configure a hypothetical (or, if you have access to test infrastructure, real) load balancer or container orchestrator health-check configuration to use your `/health/live` endpoint for *readiness* decisions (a plausible misconfiguration, since the two concepts are easy to conflate if not deliberately distinguished, exactly as this lesson has been arguing). Reason through (or actually test, if your infrastructure allows it) what happens during step 3's database-unreachable scenario under this misconfiguration: since `/health/live` still returns 200 even with the database down, the load balancer would continue routing traffic to an instance that cannot actually serve requests correctly, resulting in real user-facing errors that infrastructure automation had every opportunity to prevent by simply checking the correct endpoint. Explain why this specific misconfiguration — using the wrong one of two conceptually similar-sounding endpoints — is a genuinely easy, common mistake, and why explicitly naming and documenting the liveness/readiness distinction (not just implementing both endpoints) is part of correctly operationalizing this lesson's content.

## Transfer

If TARDOC or Mahall's deployment currently has a health check endpoint, audit it against this lesson's liveness/readiness distinction: does it check dependencies (making it more of a readiness check) or nothing beyond basic responsiveness (making it more of a liveness check), and is it currently being used for both restart decisions and traffic-routing decisions by whatever infrastructure you're running on (even a simple deployment script or systemd configuration counts) — if so, describe the specific risk this lesson's failure drill identified, applied to your actual deployment.

## Done when

You've implemented and directly demonstrated the distinction between liveness (process alive) and readiness (can currently serve traffic correctly) with real, observably different responses under a real dependency-failure scenario, and you can explain — using the failure drill's misconfiguration scenario — the concrete operational consequence of using the wrong check for the wrong infrastructure decision.
