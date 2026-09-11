# Capstone 4: Open-Source Graduation

## Objective

Contribute a genuinely non-trivial fix or feature to a real open-source project — the fourth and final terminal capstone, and the curriculum's actual closing test: applying the reproduce-patch-review-submit pipeline from the open-source track to something substantial enough that it draws on real technical depth from across this curriculum, not just process discipline.

## Prerequisites

Open-source-and-real-code track (all 4 lessons — this capstone is that track's full pipeline, applied at greater scope), and, implicitly, whatever technical tracks are relevant to the specific project and issue you choose (if you contribute to a database project, database-internals; if a networking library, the networking track; and so on).

## Learn

There is no new material. This is deliberately the curriculum's final lesson, and its purpose is different in kind from every other capstone: where Capstones 1-3 test whether you can *build* integrated systems, this one tests whether you can bring everything you've learned to bear on someone *else's* real system, under real external evaluation, on a problem substantial enough to matter.

**Why "non-trivial" is the operative word distinguishing this from open-source Lesson 2-4's own exercise.** A documentation fix or a trivial one-line bug fix (legitimate, useful contributions, and a reasonable outcome if that's genuinely what's appropriate) don't require much of this curriculum's actual technical depth. This capstone specifically asks you to find and complete something that *does* — a bug whose root cause requires real systems reasoning to understand (a concurrency issue, a subtle protocol violation, a performance regression with a real, diagnosable cause), or a genuinely useful, moderately-scoped feature a project has expressed interest in but hasn't yet had someone implement.

**Why this is the right way to end the curriculum, not just a formality.** Every technical lesson in this curriculum has been built and verified in a controlled, self-directed context — you chose the scope, you decided when it was "done," and even the failure drills were designed by this curriculum's own authors to be instructive rather than genuinely adversarial. A real open-source contribution has none of these comforts: the project's existing code has real constraints you didn't choose, a real maintainer will judge your work by their own standards, and the problem itself might turn out to be harder, or different, than it first appeared. This is deliberately closer to real engineering work than anything else in this curriculum, and finishing it well is the most honest evidence available that the preceding 20+ tracks actually built real, applicable capability.

## Attempt

1. **Select a suitable issue** (open-source Lesson 1-2's methodology, at greater scope): find a real, open, unresolved issue in a real project — ideally one connected to a technical area this curriculum covered in depth (a concurrency bug if you completed OS/Rust concurrency lessons, a query-planning issue if you completed database-internals, a protocol-handling bug if you completed networking) — and confirm, before committing significant time, that it's both genuinely unresolved and a reasonable scope for you to actually complete.

2. **Reproduce, patch, test, and review** (open-source Lessons 2-3's full pipeline): minimize a reliable reproduction, identify the actual root cause (not just a symptom-level fix), implement the smallest correct patch with proper regression coverage, and review your own diff with genuine maintainer-level scrutiny before submission.

3. **Ship the change upstream where accepted** (open-source Lesson 4's submission-and-feedback cycle): submit your contribution, engage genuinely with reviewer feedback, and see it through to a real resolution — merged, or a documented, honest account of why it wasn't (a legitimate, real possible outcome of open-source contribution that's worth documenting honestly rather than avoiding by picking only trivially-safe issues).

## Verify

Present the complete, real trail of evidence: your reproduction (open-source Lesson 2's methodology), your root-cause analysis and patch (open-source Lesson 3's methodology), your actual submitted pull request or patch, the actual reviewer feedback received, and the actual final outcome. This capstone's verification is inherently external — a real maintainer's assessment, not a self-graded checklist — which is precisely the point.

## Failure drill

If your contribution is rejected, requires substantially more rework than expected, or the project simply doesn't respond, resist the temptation to treat this as a failed exercise and quietly substitute an easier target. Instead, apply system-engineering Lesson 12's alternatives-considered discipline explicitly: write out what you'd do differently if attempting a similar contribution again, what specifically about this project's standards or your own approach caused the outcome, and whether the technical substance of your fix was sound even if the *process* outcome wasn't what you'd hoped. A rejected-but-technically-sound contribution, honestly analyzed this way, demonstrates more real capability than an artificially easy, guaranteed-to-succeed contribution chosen specifically to avoid this risk.

## Transfer

Reflect, in writing, on the complete arc from this curriculum's Lesson 00 (orientation) through this final capstone: which specific track's material did you draw on most directly for this final contribution, which track, in retrospect, do you feel least confident in and would want to revisit, and what's the most significant way your actual approach to reading, debugging, and reasoning about unfamiliar code has changed between when you started this curriculum and now, evidenced concretely by how you approached this specific open-source contribution compared to how you'd likely have approached it before.

## Done when

You've selected and completed a genuinely non-trivial open-source contribution — reproduced, root-cause-analyzed, patched, tested, self-reviewed, and submitted for real external review — and seen it through to a real, honestly-documented outcome, whatever that outcome was, with an honest reflection connecting this final piece of work back to the specific technical depth this curriculum built across its full 21 tracks.
