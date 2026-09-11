# Lesson 11: Safe Deployments

## Objective

Implement a deployment strategy that lets you release changes with a bounded blast radius and a fast, reliable rollback path — rather than a single "replace everything at once and hope" deployment, where any bug affects 100% of traffic immediately with no easy way back.

## Prerequisites

System-engineering Lesson 9 (observability — safe deployment strategies depend on being able to detect a bad deployment quickly, which requires the metrics/logging this lesson assumes exist), networking Lesson 6 (load balancing — several safe-deployment patterns rely directly on routing control).

## Learn

**Why "deploy and hope" is a real, common, and avoidable risk.** A deployment that replaces every running instance simultaneously with new code means: if the new code has a bug that only manifests under real production conditions (not caught by testing), 100% of your traffic is immediately affected, and rolling back means re-deploying the previous version — itself another deployment, taking real time during which users are actively affected. The core idea this lesson introduces: control *how much* traffic sees a new deployment, and *how fast* you can detect and reverse a problem, rather than accepting an all-or-nothing model by default.

**Rolling deployments: replacing instances gradually, not all at once.** Instead of stopping all instances and starting new ones simultaneously, replace them incrementally (e.g. one at a time, or in small batches), with health checks (API engineering Lesson 14's readiness endpoint) confirming each new instance is actually healthy before proceeding to the next — if a new instance fails its readiness check, the rollout can halt automatically before affecting more of the fleet, rather than blindly continuing to replace everything.

**Blue-green deployments: two complete environments, instant traffic switch.** Maintain two full, independent production environments ("blue" currently live, "green" idle or being prepared with the new version). Deploy and fully test the new version in green while blue continues serving all production traffic, then switch the load balancer to route to green — and critically, keep blue running and ready, so if a problem is discovered post-switch, rolling back is just switching the load balancer back, a fast, low-risk operation, rather than a fresh deployment under pressure.

**Canary deployments: route a small percentage of real traffic to the new version first.** Rather than an all-or-nothing switch, route (e.g.) 5% of real production traffic to the new version while 95% continues on the old, stable version — and actively monitor the canary's metrics (error rate, latency, per system-engineering Lesson 9's observability) against the stable version's baseline before gradually increasing the percentage. This directly bounds blast radius: a bug in the new version affects only the small canary fraction of real users, not everyone, and is detected via real production traffic and conditions (which testing, however thorough, can never perfectly replicate) before broader rollout.

**The common thread across all three patterns.** Each is a different way of achieving the same underlying goal: limit how much of your system is exposed to a new, unproven version at once, and maintain a fast path back to the known-good state if something goes wrong — the specific mechanism (gradual instance replacement, environment switching, traffic percentage) differs, but the goal (bounded blast radius, fast rollback) is identical across all of them.

## Attempt

1. For a real or simulated service (reuse your system-engineering Lesson 9 observable service, ideally), design a rolling deployment process: define the batch size (how many instances replaced at once), the health-check gate (per API engineering Lesson 14) each new instance must pass before proceeding, and the halt condition (what causes the rollout to stop automatically rather than continuing).

2. Simulate a rolling deployment with a deliberately introduced bug in the "new version" that only manifests under specific conditions (e.g. fails its readiness check after a few seconds of running, simulating a delayed startup issue not caught by a naive immediate health check). Confirm your rollout process correctly halts after detecting the failing instance, rather than continuing to replace the remaining healthy instances with the same broken version.

3. Implement a basic canary deployment: route a configurable percentage of simulated traffic to a "new version" (which you can make deliberately buggy for testing, e.g. an elevated error rate) while the rest continues to the stable version, and monitor both versions' error rates (system-engineering Lesson 9's metrics) separately. Confirm you can detect the canary's elevated error rate specifically, distinguishing it from the stable version's baseline, using only your metrics — not by manually inspecting individual requests.

4. Implement the rollback path for your canary setup: upon detecting the canary's elevated error rate (step 3), route 100% of traffic back to the stable version, and measure how quickly this rollback actually takes effect (from detection to full traffic reversion) — report the actual measured time, which is your concrete "how fast can we recover from a bad deployment" number for this specific mechanism.

## Verify

For step 2, show the actual rollout log/trace demonstrating the halt occurring after the specific failing instance was detected, with the remaining, not-yet-replaced instances confirmed to still be running the previous, known-good version. For step 4, report the actual measured rollback time from detection to full traffic reversion.

## Failure drill

Simulate a canary deployment where the bug is *not* elevated error rate (which your step 3 monitoring correctly catches) but instead a subtle correctness bug that doesn't manifest as an HTTP error at all (e.g. the new version silently computes a wrong value but still returns 200 OK) — confirm your error-rate-based canary monitoring from step 3 does *not* detect this specific class of bug, since it's not looking for it. Explain why this is a genuine, important limitation of canary deployments as commonly implemented: they're effective against bugs that manifest as elevated error rates or latency (things your metrics are actually watching for), but silent, "successful-looking but wrong" bugs require different detection mechanisms (business-metric monitoring, data validation checks, or more thorough pre-deployment testing) — canary deployment reduces blast radius for a *class* of bugs, it doesn't catch every possible kind of deployment problem, and conflating "we do canary deployments" with "we're safe from all deployment risk" would be a real, dangerous overconfidence.

## Transfer

If TARDOC's or Mahall's current deployment process (given the single-VPS setup mentioned in your project history) is closer to "stop and replace everything at once" than any of this lesson's patterns, describe which pattern (rolling, blue-green, or canary) would be most practical to adopt first given your actual infrastructure constraints (a single VPS may make true blue-green, requiring two full parallel environments, harder to justify than a simpler rolling or canary approach), and what the minimum viable version of that pattern would look like for your specific setup.

## Done when

You've designed and tested a rolling deployment that correctly halts on a detected failure rather than continuing blindly, you've implemented and tested a canary deployment that detects and enables fast rollback from an elevated-error-rate bug, with a real measured rollback time, and you've directly demonstrated — via the failure drill — a specific class of bug your canary monitoring does not catch, understanding this as a genuine scope limitation rather than a gap in your implementation specifically.
