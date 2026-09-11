# Lesson 3: SLOs and Error Budgets

## Objective

Define a real Service Level Objective (SLO) with a measurable target, calculate error budget consumption from real or realistic data, and write a concrete policy connecting error-budget status to actual engineering decisions (feature velocity versus reliability investment).

## Prerequisites

System-engineering Lesson 8 (reliability targets — this lesson formalizes that lesson's availability-target discussion into the specific SLO/error-budget framework), statistics track (percentiles — SLOs are commonly defined in terms of percentile latency or success-rate thresholds).

## Learn

**SLI, SLO, SLA: three related but distinct terms, worth being precise about.** A Service Level *Indicator* (SLI) is the actual measured metric (e.g. "percentage of requests completing in under 200ms," or "percentage of requests returning a non-5xx status"). A Service Level *Objective* (SLO) is a target value for that indicator (e.g. "99.5% of requests complete in under 200ms, measured over a rolling 30-day window"). A Service Level *Agreement* (SLA) is an external, often contractual commitment (frequently with financial or other consequences for missing it) — typically set looser than your internal SLO, so you have margin to notice and respond to a problem internally before it becomes an SLA breach visible to (and potentially costly with) external parties.

**Error budget: the mathematically precise flip side of an SLO.** If your SLO is 99.5% success over 30 days, your error budget is the remaining 0.5% — the amount of "acceptable failure" you have to spend before violating your own stated objective. Concretely, for a service handling 1 million requests over 30 days, a 99.5% SLO means an error budget of 5,000 failed/slow requests for that period — a genuinely useful, concrete number, since it converts an abstract reliability target into a specific, trackable quantity you can measure consumption against in real time.

**Why error budgets are a genuinely useful engineering-decision tool, not just a reporting metric.** The core idea: as long as you have error budget remaining, you have room to take calculated risks (ship a new feature quickly, run an aggressive experiment) — but once the budget is exhausted (or trending toward exhaustion faster than the period allows), that's a signal to shift focus toward reliability work and slow down on risky changes, until the budget recovers. This turns "should we prioritize this reliability fix or this new feature" from a subjective argument into a data-driven decision with an explicit, pre-agreed trigger — the kind of policy system-engineering Lesson 12's ADR practice would document explicitly, stating in advance what happens at various budget-consumption thresholds, rather than deciding reactively and inconsistently each time.

**Choosing the right SLI and threshold: a real design decision, not an arbitrary number.** An SLO should reflect what actually matters to users/business (system-engineering Lesson 8's business-impact reasoning, applied specifically here) — a latency SLO on p99, not just average, per statistics track Lesson 1's point about percentiles mattering more than averages for user-perceived experience; a success-rate SLO scoped to the specific operations that matter most, not a vague, unscoped "overall uptime" number that could mask a critical path being broken while non-critical paths compensate in an aggregate average.

## Attempt

1. For a real service (TARDOC or Mahall), define a concrete SLO: pick a specific SLI (e.g. "percentage of API requests returning within 500ms" or "percentage of transcription jobs completing successfully within 10 minutes"), a target percentage, and a measurement window (e.g. rolling 30 days) — justify your chosen target using system-engineering Lesson 8's business-impact reasoning, not an arbitrary round number.

2. Using real or realistically estimated traffic volume for your service, calculate the actual error budget in concrete units (e.g. "X failed/slow requests allowed per 30-day window") for your step 1 SLO.

3. Using real (if available) or synthetic/simulated latency and error data for a representative period, calculate actual error budget consumption — what fraction of your calculated budget has been "spent" by the actual (or simulated) failure/slow-request rate over that period, and whether the service is currently trending toward exhausting its budget before the window resets.

4. Write an explicit error-budget policy: state, in concrete terms, what happens at specific consumption thresholds (e.g. "at 50% budget consumed with 10 days remaining in the window, pause non-critical feature releases and prioritize reliability work until consumption trend improves; at 90% consumed, halt all non-emergency releases entirely") — a real, actionable policy someone could actually follow, not a vague statement of intent.

## Verify

Present your step 1 SLO definition with justification, your step 2 calculated error budget in concrete units, your step 3 actual/simulated consumption calculation with the resulting trend assessment, and your step 4 explicit policy with concrete thresholds and actions.

## Failure drill

Take your step 3 consumption calculation and construct a scenario where the budget is being consumed at a rate that, if it continued linearly, would exhaust the entire 30-day budget by day 15 — explicitly calculate this projected exhaustion date, and then apply your step 4 policy to determine what action it actually specifies for this situation. If your policy doesn't have a clear, specific trigger for "budget is on pace to be exhausted early, not just currently low," revise it to include one — this is a genuinely common gap in error-budget policies that only specify total-consumption thresholds without accounting for *rate* of consumption, which matters because a budget that's 50% consumed with 25 days remaining is a very different situation from 50% consumed with 3 days remaining, even though the raw consumption percentage is identical.

## Transfer

If TARDOC's actual infrastructure (a single Contabo VPS, per your project history) makes hitting a genuinely ambitious SLO (say, 99.9%) unrealistic without infrastructure investment you haven't yet made, describe honestly what SLO is actually achievable given current infrastructure, and connect this back to system-engineering Lesson 8's transfer task about matching aspirational targets to actual infrastructure capability — an SLO you cannot realistically meet given current infrastructure isn't a useful target, it's just an ongoing, unaddressed budget violation with no real decision-making value.

## Done when

You've defined a real, business-justified SLO with a specific SLI and threshold, calculated the corresponding error budget in concrete units, computed actual or realistic budget consumption and its trend, and written an explicit, actionable policy with specific thresholds — including, per the failure drill, accounting for consumption *rate* and not just total consumed percentage.
