# Lesson 4: Upstream Contribution

## Objective

Navigate a real open-source project's contribution and review process — submitting a genuine contribution where appropriate, responding professionally to reviewer feedback, and reflecting on what the review process itself taught you, closing the loop on this track's full pipeline from reading an unfamiliar codebase (Lesson 1) to a merged (or at least seriously reviewed) real-world change.

## Prerequisites

Lesson 3 (patch and review — this lesson takes that lesson's self-reviewed patch and submits it for genuine external review), Linux tools Lesson 3 (Git workflow — pull requests build directly on the Git fundamentals covered there).

## Learn

**Why submitting to a real project, when appropriate, is qualitatively different from any self-directed exercise in this curriculum.** Every other lesson in this curriculum, however rigorous its failure drills and verification steps, is ultimately self-graded — you decide when you've met the "done when" criteria. A real upstream contribution is reviewed by someone with no obligation to be generous, who has their own standards for what's acceptable in their project, and who may reject, request changes to, or simply never respond to your contribution — this genuine external evaluation is precisely the value of this final lesson, and precisely why it can't be fully simulated by any exercise design.

**"When appropriate" — a real, important qualifier.** Not every bug fixed in Lessons 2-3 should necessarily be submitted upstream: check the project's contribution guidelines (many projects have explicit ones — `CONTRIBUTING.md` is a common convention), check whether the specific bug is already being addressed elsewhere (a duplicate, unwanted contribution wastes a maintainer's time), and honestly assess whether your fix is actually ready for external review (per Lesson 3's self-review rigor) rather than submitting something you know has gaps just to complete this exercise. If your Lesson 2-3 work doesn't turn out to be a good candidate for upstream submission (a very real, legitimate outcome), documenting why not, and instead finding a different, smaller, genuinely-appropriate contribution to a project (even something like a documentation fix or a well-scoped test addition) is a reasonable adjustment.

**Responding to reviewer feedback: a real communication skill, not just a technical one.** Reviewer feedback can range from a straightforward requested change to pushback on your fundamental approach, or a request for context you hadn't provided. Responding well means engaging genuinely with the specific concern raised (not defensively dismissing it, and not silently capitulating without understanding whether the feedback is actually correct), asking clarifying questions when the feedback itself is unclear, and being willing to revise your approach if the reviewer's concern is legitimate — even if it means more work than your original submission anticipated.

**What review feedback actually teaches, beyond the specific fix.** A maintainer's feedback on your specific patch often reveals something more general — a convention you didn't know about (worth updating your Lesson 1 mental model of the project), a subtlety about the codebase's design you'd missed, or a different way of thinking about the problem than your own approach. Explicitly reflecting on and recording this (not just fixing what was asked and moving on) is what turns one specific contribution into a genuinely transferable lesson about working in this specific codebase, and about code review generally.

## Attempt

1. Review your Lesson 2-3 bug fix against your chosen project's actual contribution guidelines (if any exist) and assess honestly whether it's ready and appropriate for submission — if it's not (the bug is already fixed elsewhere, the project isn't accepting contributions in that area, your fix doesn't meet the project's standards on reflection), document why, and find an alternative, genuinely appropriate contribution to the same or a different project instead.

2. Submit your contribution following the project's actual process (a pull request, typically) — including a clear description (per Lesson 3's patch-description guidance) and confirming you've followed any project-specific requirements (a contributor license agreement, specific commit message format, and so on).

3. Respond to whatever review feedback you receive, engaging genuinely with each specific point raised — if a change is requested, make it and explain what you changed and why; if you disagree with a piece of feedback, explain your reasoning respectfully rather than either capitulating without genuine agreement or dismissing the reviewer's concern without real consideration.

4. Regardless of the final outcome (merged, rejected, or still pending — all are legitimate, real outcomes of the actual open-source contribution process, not something fully within your control), write a reflection: what did you learn about the project specifically, what did you learn about code review generally, and what would you do differently in your next contribution based on this experience.

## Verify

Present your actual submitted contribution (the pull request or equivalent), the actual reviewer feedback you received (if any — a genuinely valuable outcome even if the review is minimal or slow, and worth documenting honestly if the project simply hasn't responded, which is itself a realistic aspect of open-source contribution worth experiencing), your responses to that feedback, and your final reflection.

## Failure drill

If your contribution receives push-back or a rejection you initially feel is unwarranted, deliberately apply system-engineering Lesson 12's "seek the strongest counter-argument for your own position" discipline before responding: articulate, in writing, the strongest possible case *for* the reviewer's position, even if you still ultimately disagree after doing so. If, after genuinely attempting this, you still believe your original approach was correct, that's a legitimate outcome — but if the exercise reveals the reviewer had a point you'd initially dismissed too quickly, revise your response (and your understanding) accordingly. Explain why this discipline — actively steelmanning disagreement before responding, rather than either capitulating reflexively or defending your position reflexively — is a genuinely valuable communication skill this lesson is specifically testing, beyond the purely technical content of the contribution itself.

## Transfer

Reflect on how this lesson's full pipeline (Lesson 1's repository reading, Lesson 2's bug reproduction, Lesson 3's patch-and-self-review, this lesson's submission-and-feedback cycle) compares to how you currently work on TARDOC, Mahall, or Lead Sourcer as a solo developer with no external code review — what specific value did the external review step (even if it was minimal, slow, or ultimately a rejection) provide that your own self-review process, however rigorous, structurally cannot replicate? Consider whether finding a trusted collaborator or a code-review service for your own projects would provide some of this same value going forward, now that you've directly experienced what it adds.

## Done when

You've submitted a genuine contribution to a real open-source project (or made and documented a considered decision that your specific Lesson 2-3 work wasn't appropriate for submission, substituting a different genuine contribution instead), engaged authentically with whatever review feedback resulted, and written an honest reflection connecting what you learned to both the specific project and to code review as a general discipline — including, via the failure drill, having genuinely steelmanned any disagreement with reviewer feedback before responding to it.
