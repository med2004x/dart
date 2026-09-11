# Lesson 15: Architecture Review

## Objective

Conduct a rigorous architecture review of someone else's design (or your own, treated with genuine critical distance) — the specific skill of finding real weaknesses in a design before it's built, which is different from and complements the design-production skill Lesson 14 focused on.

## Prerequisites

Lesson 14 (design capstone — this lesson reviews a design of that same shape, either your own or an external one), Lesson 12 (architecture decisions — a good review specifically interrogates whether the "alternatives considered" reasoning actually holds up).

## Learn

**Why reviewing is a genuinely different skill from designing, not just "the same thing done by someone else."** When you design something, you're inside your own reasoning — the assumptions and tradeoffs that led to your conclusion feel natural and easy to defend, precisely because you made them and lived through the reasoning process. Reviewing requires deliberately stepping outside that — actively looking for what the designer might have underweighted, assumed without justification, or failed to consider, which is a different cognitive stance than generating the design in the first place. This is exactly why "review your own work" is harder than reviewing someone else's, and why the failure-drill pattern used throughout this curriculum (deliberately trying to break your own designs) is a specific technique for approximating that outside perspective when no other reviewer is available.

**What a substantive architecture review actually checks, beyond "does this look reasonable."** Does the stated capacity/requirements reasoning (Lesson 2) use real numbers or arbitrary placeholders? Are the module boundaries (Lesson 3) genuinely justified, or is the decomposition arbitrary? Does the consistency model (Lesson 7) correctly classify which data needs strong consistency, or is there a plausible scenario (per that lesson's failure drill) where the stated tradeoff produces a real user-visible problem? Are the reliability targets (Lesson 8) grounded in actual business impact reasoning, or asserted without justification? Do the ADRs (Lesson 12) genuinely engage with alternatives, or are they one-sided justifications of a decision already made? Each of these questions maps directly to a specific lesson in this track — an architecture review is, in a real sense, applying every earlier lesson's critical-thinking discipline to someone else's (or your own) finished design.

**Giving critical feedback usefully — a genuinely separate skill from finding the flaw.** Identifying a weakness is only half the value; communicating it in a way that's specific, actionable, and doesn't just register as generic criticism is what makes a review actually useful to the person receiving it. "The consistency model seems risky" is much less useful than "the consistency model treats X as eventually-consistent, but per the scenario I traced through, a user could observe Y within Z seconds of the primary action, and given your stated requirement that users immediately see confirmation, this seems like a real gap, not just a theoretical one" — specific, traceable, and tied to a concrete scenario rather than a vague impression.

## Attempt

1. Obtain a design document to review — ideally your own Lesson 14 capstone (reviewed with as much critical distance as you can manage, perhaps after a meaningful time gap), or, if available, a real design document from someone else's work, or a plausible design document you construct specifically to contain several deliberate, realistic flaws for practice purposes.

2. Systematically review it against every checklist item implied by Learn's second paragraph — for each of Lessons 2, 3, 7, 8, 12 specifically, write down whether that section of the document holds up to scrutiny, and if not, exactly what's missing or unjustified.

3. For at least 2 identified weaknesses, go beyond "this seems weak" and construct a specific, traceable scenario (following the pattern from Learn's third paragraph) demonstrating the actual concrete consequence of the gap — not just asserting a problem exists, but showing, step by step, what would actually go wrong and under what conditions.

4. Write the review as a document intended for the design's author (even if that's you) — organized, specific, and including both genuine strengths (a review that's 100% criticism, with no acknowledgment of what's actually solid, is both less useful and less accurate than a balanced one) and the concrete weaknesses from step 3, each with a specific, actionable suggestion for what would address it.

## Verify

Present your completed review document, and for the 2+ weaknesses identified in step 3, show the specific scenario you traced through demonstrating the concrete consequence, not just the abstract concern.

## Failure drill

Take your own step 4 review and, before finalizing it, deliberately try to find a counter-argument for your strongest identified weakness — is there a reason the original design's choice might actually be defensible that you didn't initially consider (echoing Lesson 12's own critical-reader exercise, but now applied to your review itself, one level removed)? If you find a genuine counter-argument, revise your review to acknowledge it rather than presenting your critique as more certain than it actually is. Explain why a review that itself withstands this kind of scrutiny — one that's been tested against its own strongest counter-argument — is more valuable and more trustworthy to its recipient than one presented with unexamined confidence, and connect this to why the discipline of actively seeking disconfirming evidence for your own conclusions (not just for the thing you're reviewing) is the more general skill this whole exercise is building.

## Transfer

If you have access to (or can imagine) a real design decision from TARDOC, Mahall, or Lead Sourcer that was made without a formal review process, apply this lesson's checklist retroactively to it, and identify honestly whether a review at the time — using this same systematic approach — would likely have caught anything the actual, unreviewed decision missed. If the answer is genuinely "no, it holds up fine," that's a legitimate and useful finding too, not a failure of the exercise.

## Done when

You've completed a systematic review of a real or realistic design document against this track's full checklist of concerns, you've constructed at least 2 specific, traceable scenarios demonstrating concrete consequences of identified weaknesses rather than vague criticism, and you've subjected your own review to a genuine counter-argument check, revising it if warranted — demonstrating the same critical-distance discipline toward your own critique that the review itself was meant to apply to the original design.
