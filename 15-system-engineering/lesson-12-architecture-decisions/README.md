# Lesson 12: Architecture Decisions and Tradeoff Analysis

## Objective

Write a real Architecture Decision Record (ADR) for a genuine design choice — documenting not just what was decided, but the alternatives considered and specifically why they were rejected — building the discipline of making tradeoffs explicit rather than leaving them implicit in code that no one can later reconstruct the reasoning behind.

## Prerequisites

Lessons 1-11 of this track (this lesson is about documenting and communicating exactly the kind of tradeoff reasoning those lessons have been building throughout).

## Learn

**Why undocumented architecture decisions are a real, recurring cost.** Six months after a decision was made (perhaps by you, perhaps by a collaborator, perhaps by your own past self), the code reflects *what* was decided but rarely *why* — and without that context, a future change either repeats the same reasoning process from scratch (wasted effort re-deriving something already figured out once) or, worse, undoes a deliberate decision without realizing it was deliberate, reintroducing a problem the original decision was specifically designed to avoid.

**What an ADR actually contains, at minimum.** Context (what problem or requirement prompted this decision), the decision itself (stated clearly and unambiguously), alternatives considered (genuinely considered, not straw-man options included only to make the chosen one look obviously superior), and consequences (what tradeoffs the chosen decision accepts — every real architectural decision has downsides, and naming them explicitly is part of the record's value, not something to omit to make the decision look cleaner than it was).

**Why "alternatives considered" is the most valuable, and most often skipped, section.** A decision record that only states what was chosen, with no record of what else was considered and why it was rejected, provides much less value to a future reader trying to understand whether circumstances have changed enough to warrant reconsidering — if the ADR explicitly says "we rejected microservices here because our team size and deployment complexity didn't justify the coordination cost," a future reader facing a genuinely different context (a much larger team, a real need for independent scaling) has exactly the information needed to recognize the original reasoning may no longer apply, rather than having to guess whether the monolith was a deliberate choice or simply the default nobody questioned.

**ADRs as a communication tool, not just a personal memory aid.** Even working solo, writing an ADR forces the discipline of articulating tradeoffs precisely enough to defend them to a skeptical reader — a genuinely useful exercise even with an audience of one, since it's easy to convince yourself a decision is obviously correct without ever stating the alternatives clearly enough to notice a weak point in your own reasoning. For a decision that will eventually involve other people (a collaborator, a future hire, or even future-you needing to explain a past choice), an ADR is the artifact that makes the reasoning transferable rather than trapped in whichever person's head made the original call.

## Attempt

1. Identify a real, genuine architectural decision you've already made in TARDOC, Mahall, or Lead Sourcer (per your project history — e.g. the choice to migrate to Groq-hosted Whisper with multi-key rotation, or the decision to restructure TARDOC's repo, or Mahall's choice of path-based seller URLs over subdomain architecture per your competitive analysis of Converty) and write a complete ADR for it: context, decision, at least 2 genuinely-considered alternatives with specific reasons for rejection, and explicit consequences/tradeoffs accepted.

2. For the alternatives section specifically, be honest about tradeoffs — if an alternative had a genuine advantage the chosen option lacks, state it explicitly rather than only listing the alternative's downsides; a one-sided "why the alternative was obviously worse" writeup is less useful (and less honest) than one that acknowledges genuine tradeoffs even in a decision you're confident was correct.

3. Write a second ADR for a decision you *haven't* made yet but are currently facing (a genuine, open decision in your current work) — going through the same structure prospectively rather than retrospectively, and using the act of writing it to actually help you reach a more considered decision, not just to document one already made.

4. Share (or imagine sharing, if working solo) your step 3 ADR with a critical reader's mindset: reread it as if you were someone skeptical of your chosen option, and identify the strongest counter-argument for one of the alternatives you rejected — if you find your original reasoning doesn't actually hold up well against this scrutiny, revise the ADR's decision or reasoning accordingly, rather than defending the original choice for its own sake.

## Verify

Present both completed ADRs (steps 1 and 3), and for step 4, show explicitly what counter-argument you identified and whether it changed your final decision or reasoning, or whether it strengthened your confidence in the original choice by surviving the scrutiny — either outcome is a legitimate result of the exercise, but the counter-argument needs to have been genuinely engaged with, not dismissed without real consideration.

## Failure drill

Take your step 1 retrospective ADR and deliberately write a version with the "alternatives considered" section removed or reduced to a single dismissive sentence per alternative ("we considered X but it was worse"). Compare the two versions side by side and identify specifically what information a future reader (including future-you) would lose from the stripped-down version — could they, from the stripped version alone, determine whether changed circumstances (e.g. team growth, new requirements) might warrant reconsidering the decision? Explain why the stripped version, despite still technically "documenting a decision," fails at the actual value an ADR is supposed to provide, and connect this concretely to a real situation (if one exists) where you've previously had to re-derive reasoning behind an old decision because it wasn't adequately recorded the first time.

## Transfer

If Lead Sourcer or TARDOC's future roadmap includes a decision you know is coming (e.g., at what point moving from a single VPS to a multi-instance setup would be justified, echoing system-engineering Lesson 6's scaling discussion, or when Mahall might need to reconsider its current architecture given growth), write a brief, prospective ADR sketch for that anticipated future decision now, while the relevant context and constraints are fresh, rather than waiting until the decision is urgent and the reasoning has to be reconstructed under time pressure.

## Done when

You've written a complete, honest retrospective ADR for a real past decision with genuinely-considered alternatives and explicit accepted tradeoffs, you've written a prospective ADR for a real, currently-open decision and subjected it to genuine critical scrutiny that either changed or strengthened your reasoning, and you can articulate specifically what value the "alternatives considered" section provides that a bare decision statement would not.
