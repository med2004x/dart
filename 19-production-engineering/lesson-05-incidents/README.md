# Lesson 5: Incident Response

## Objective

Run a realistic incident-response simulation end to end — diagnosing under time pressure using only production Lesson 2's telemetry, maintaining a clear timeline while it's happening, and producing a genuine postmortem with concrete, trackable corrective actions afterward.

## Prerequisites

Production Lesson 2 (observability — the diagnostic tools this lesson's incident response depends on), production Lesson 3 (SLOs/error budgets — the framework for assessing an incident's actual severity/impact), system-engineering Lesson 8 (reliability/recovery mechanisms — what an incident response might invoke).

## Learn

**Why incident response is a distinct skill from debugging in a calm, unhurried context.** Ordinary debugging (Linux tools Lesson 4, performance track's whole toolkit) assumes you have time to explore, hypothesize, and test carefully. An active incident has real, ongoing user impact accumulating with every minute — this creates genuine pressure to act quickly, which is precisely when mistakes (a rushed, unverified fix that makes things worse, a wrong diagnosis acted on prematurely) are most likely. Incident response as a discipline is specifically about managing this tension: acting with appropriate urgency without sacrificing the rigor that prevents making things worse.

**Maintaining a timeline while the incident is happening, not reconstructing it afterward from memory.** A live, timestamped log of what was observed, what was tried, and what happened as a result — updated *during* the incident, not written up afterward — serves two purposes: it keeps the responder's own reasoning organized under pressure (reducing the risk of repeating an already-tried, already-failed action), and it provides an accurate record for the postmortem, since human memory of a stressful, fast-moving event is notoriously unreliable after the fact, even a short time later.

**The postmortem: what makes it valuable versus merely a formality.** A genuinely useful postmortem is blameless (focused on what happened and why the system/process allowed it, not on assigning fault to a person — a person acting on the information and tools available at the time rarely "should have known better" in a way that blame usefully addresses) and produces *specific, trackable* corrective actions (not vague statements like "improve monitoring," but "add a p99 latency alert on endpoint X with threshold Y, owned by [you], due by [date]") — the difference between a postmortem that actually reduces future risk and one that's a ritual with no lasting effect.

**Corrective actions need a mechanism to actually get done, not just be listed.** A postmortem's corrective-action list, if not tracked and followed up on with the same rigor as any other planned work, tends to be forgotten once the immediate pressure of the incident has passed — treating corrective actions as real, tracked work items (in whatever task-tracking system you actually use) with an owner and a deadline is what closes the loop between "we learned something from this incident" and "the system is actually more resilient as a result."

## Attempt

1. Design and inject a realistic, synthetic failure into a test service (reuse production Lesson 2's simulated service) — something you don't fully script the exact symptoms of in advance, to genuinely simulate not-yet-knowing-the-cause, even though you're the one triggering it.

2. Respond to the incident in real time: maintain a live, timestamped log as you investigate (what you observed, what you checked, what you tried), using only production Lesson 2's telemetry (logs, metrics, traces) to diagnose — no looking at the actual failure-injection code you wrote, treating it as genuinely unknown until your telemetry-based investigation reveals it.

3. Once diagnosed, apply a fix (or a mitigation — distinguish explicitly between a full fix and a temporary mitigation that restores service while a proper fix is developed later, a real and common incident-response distinction) and confirm, via your telemetry, that the fix/mitigation actually resolved the symptom, not just that you believe it should have.

4. Write a complete, blameless postmortem: a factual timeline (derived from your step 2 live log, not reconstructed from memory afterward), a root-cause analysis, an assessment of actual impact (using production Lesson 3's SLO/error-budget framework — how much budget did this incident consume), and at least 3 specific, trackable corrective actions each with a concrete description of what would need to be done.

## Verify

Present your step 2 live incident log (with real timestamps showing the actual investigation sequence, including any dead ends or wrong hypotheses you pursued before finding the real cause — an honest log, not a cleaned-up version that only shows the correct path), your step 3 confirmation that the fix actually worked (via telemetry, not assumption), and your complete step 4 postmortem document.

## Failure drill

Review your own step 2 live log and identify at least one point where you pursued an incorrect hypothesis before finding the actual root cause — describe explicitly what telemetry evidence *should* have ruled out that incorrect hypothesis sooner, had you noticed it, versus what evidence actually led you to eventually abandon it. Explain why this honest self-review — including the wrong turns, not just the successful path — is itself valuable postmortem material: if a specific piece of available telemetry could have ruled out a wrong hypothesis faster but wasn't obviously salient in the moment, that's a real finding about how your telemetry is organized/presented, potentially worth its own corrective action (e.g. a dashboard reorganization, or an alert that would have pointed more directly at the actual cause) distinct from whatever fixed the underlying incident itself.

## Transfer

If TARDOC has experienced a real incident (the silent billing bug, per your project history, is a documented example, even if it wasn't formally postmortem'd at the time), write a retrospective postmortem for it now, using this lesson's complete framework — timeline (reconstructed as accurately as you can from whatever records exist), root cause, impact assessment, and concrete corrective actions — and note explicitly whether any of the corrective actions you'd now propose have already been addressed by subsequent work, or remain genuinely open.

## Done when

You've responded to a genuinely blind (to you, in the moment) synthetic incident using only telemetry-based diagnosis, maintained a real-time, honest log including wrong turns, confirmed your fix's effectiveness via evidence rather than assumption, and produced a complete, blameless postmortem with specific, trackable corrective actions — plus honestly identified, via the failure drill, a point where better telemetry organization could have shortened your own diagnostic path.
