# Lesson 2: Observability (Production Context)

## Objective

Take system-engineering Lesson 9's observability implementation (logs, metrics, traces) and apply it to diagnosing a realistic, synthetic production incident end to end — defining what signals actually matter before an incident happens, not scrambling to add instrumentation during one.

## Prerequisites

System-engineering Lesson 9 (observability fundamentals — this lesson is that lesson's discipline applied under simulated real incident pressure), performance track Lesson 6 (system-level diagnosis tools, used here alongside application-level telemetry).

## Learn

**Why "we'll add logging when something breaks" is backwards.** By definition, an incident is discovered *after* it starts — if the telemetry needed to diagnose it isn't already flowing when the incident begins, that diagnostic window is permanently lost; you cannot retroactively generate logs for requests that already happened without adequate instrumentation. This is why system-engineering Lesson 9's observability implementation needs to exist *before* it's needed, covering the signals you'd actually want during an incident you haven't had yet.

**Defining useful signals in advance: the specific question to ask for each metric/log/trace.** Not "what could we measure" (nearly unlimited) but "if this specific thing breaks, what would we need to see to diagnose it quickly" — for each critical dependency or operation, working backward from a plausible failure scenario to the specific telemetry that would make that scenario diagnosable. A metric that exists but wasn't designed with any specific diagnostic scenario in mind often turns out, when actually needed, to be missing exactly the dimension/label that would have made it useful (e.g. an error-rate metric with no breakdown by endpoint or error type, when the actual incident needs exactly that breakdown to localize the problem).

**The specific skill this lesson builds: diagnosing from telemetry alone, under time pressure, without access to live debugging.** A real production incident often cannot be debugged interactively (attaching a debugger to a production process, per Linux tools Lesson 4, is often impractical or actively risky) — the diagnosis has to come from what was already captured: logs, metrics, traces. This is a genuinely different skill from live debugging, and it's specifically what system-engineering Lesson 9's structured logging plus tracing was building toward, now tested under realistic incident conditions rather than a controlled lesson exercise.

**Golden signals: a useful starting checklist, not an exhaustive one.** Latency, traffic (request rate), errors, and saturation (how "full" a resource is — CPU, memory, queue depth) are commonly cited as the four signals worth having for nearly any service, specifically because they cover the most common failure categories (slow, overloaded, failing, or resource-exhausted) — a reasonable starting checklist when deciding what to instrument for a new component, though not a substitute for the scenario-specific thinking in Learn's second point.

## Attempt

1. For a real or realistic service (reuse or extend system-engineering Lesson 9's observable service), explicitly enumerate the golden signals (latency, traffic, errors, saturation) it currently exposes, and identify any gaps — signals you'd want but that aren't currently captured.

2. Before running any simulation, write down at least 3 plausible failure scenarios for this service (e.g. "database connection pool exhausted," "a specific downstream dependency becomes slow," "a memory leak causes gradual degradation") and, for each, state explicitly what telemetry you'd need to diagnose it — this is the "design signals from failure scenarios backward" discipline from Learn, done deliberately before the next step.

3. Actually simulate one of your step 2 scenarios (inject the failure into your test service — e.g. artificially exhaust a connection pool, or introduce a slow downstream dependency) without telling yourself in advance exactly when or how it will manifest in the telemetry, then attempt to diagnose it purely from your logs/metrics/traces, as if investigating a real incident you didn't cause yourself moments ago.

4. Write a timeline reconstruction of your step 3 incident based purely on the telemetry evidence (timestamps, specific log lines, specific metric values) — not from your memory of what you actually did to cause it — and compare this reconstructed timeline against what you know actually happened, checking specifically for any gap where the telemetry didn't clearly capture something that mattered.

## Verify

Present your step 2 signal-scenario mapping (failure scenario → needed telemetry), and your step 4 timeline reconstruction alongside the actual sequence of events, with an honest accounting of any gap between what the telemetry showed and what you know actually happened.

## Failure drill

Choose one of your step 2 scenarios that you did *not* actually simulate, and deliberately check whether your currently-implemented telemetry (from step 1) would actually capture what you said, in step 2, it would need to. If you find a gap — telemetry you assumed existed based on your step 2 reasoning, but that isn't actually implemented — this is a genuinely realistic, valuable finding: it demonstrates exactly how observability gaps go unnoticed until (or unless) a specific incident actually occurs and reveals them. Explain why this specific exercise — checking your stated intentions against actual implementation, without waiting for a real incident to force the check — is a proactive discipline worth doing periodically, not just something you're forced into by an actual outage.

## Transfer

If TARDOC has experienced any real production issue (per your project history — the silent billing bug defaulting to Geneva canton tax rates is a documented example), describe, using this lesson's golden-signals and scenario-backward-design framework, what specific telemetry (if it had existed at the time) would have surfaced that issue faster than however it was actually discovered, and whether that specific telemetry gap has since been closed or remains a real, currently-unaddressed blind spot.

## Done when

You've mapped concrete failure scenarios to the specific telemetry needed to diagnose each, before simulating any of them, you've diagnosed a real, self-induced (but blindly investigated) synthetic incident purely from telemetry evidence and produced an accurate timeline reconstruction, and you've found — via the failure drill — at least one real gap between telemetry you assumed existed and what's actually implemented, understanding this as the realistic, proactive way such gaps are normally only discovered.
