# Lesson 3: Git as an Engineering Tool

## Objective

Use Git's history-inspection tools (bisect, blame, log) as debugging instruments, not just version control for saving work — and produce commit history that's actually useful to read later, including by yourself in six months.

## Prerequisites

None, though this lesson assumes basic add/commit/branch familiarity as a given, and focuses on the tools most self-taught developers skip.

## Learn

**Git bisect: binary search over history.** When a regression exists somewhere between a known-good commit and the current broken state, `git bisect` performs binary search over the commit range: you mark a good and a bad commit, Git checks out the midpoint, you test and report `git bisect good` or `git bisect bad`, and it narrows the range by half each time — finding the exact offending commit in `O(log n)` steps instead of manually checking commits one at a time. This directly applies the binary search algorithm (already covered in the algorithms track) to a completely different domain: searching not sorted data, but a sequence of commits ordered by "still broken" as a monotonic property (assuming the bug wasn't fixed and then reintroduced, which would break the binary search assumption — worth stating explicitly, since bisect's correctness depends on it).

**Git blame: attributing every line to its origin commit.** `git blame <file>` shows, line by line, which commit last touched each line and who wrote it. This is most useful not for "who do I blame" but for context: `git show <commit>` on the blamed commit often reveals *why* a seemingly odd line exists — a comment in the commit message, a linked issue, or sibling changes in the same commit that explain the reasoning invisible from the code alone.

**Reading history before changing unfamiliar code.** `git log --follow <file>` (follows renames) and `git log -p <file>` (shows the actual diffs, not just commit messages) let you reconstruct how a piece of code evolved — often revealing that an apparently strange pattern exists because of a bug fix, a specific edge case, or a deliberate tradeoff that isn't visible from reading the current state alone. Changing code without this context risks reverting a fix you don't know happened.

**Clean commits, and why they're a debugging tool, not just tidiness.** A commit that bundles an unrelated refactor with a bug fix makes `git bisect` and `git blame` both less useful — bisect will flag the combined commit as "bad" without telling you which part actually caused the regression, and blame will attribute unrelated lines to a commit message about something else entirely. Small, focused commits with accurate messages are what make the tools in this lesson actually work well later.

## Attempt

1. In a test repository (create a throwaway one, don't practice on real project history), make a series of 8-10 commits to a small script, where one commit in the middle deliberately introduces a bug (e.g. an off-by-one error, or a logic inversion) while the surrounding commits are unrelated, valid changes. Note the exact commit hash of the bug-introducing commit for later verification, then forget it deliberately (don't reference it directly in the next step).

2. Use `git bisect start`, mark the current (broken) commit as bad and the first commit as good, and work through the bisect process — for each checked-out commit, run your test/script and report `good` or `bad` — until Git identifies the offending commit. Confirm it matches the hash you noted in step 1.

3. Pick a file with real history (from any of your own real projects — TARDOC, Mahall, or Lead Sourcer) and use `git log -p` on one specific function or section to reconstruct why it looks the way it does. Find at least one commit message that explains a *reason* for a specific line or pattern, not just "what changed."

4. Use `git blame` on that same file/section, identify the commit that introduced a specific line you're curious about, and use `git show <hash>` to see the full commit (message, diff, and any related context) — confirm this gives you more insight than the blame output alone (which only shows the line-to-commit mapping, not the reasoning).

## Verify

For step 2, report the number of bisect steps it took to find the bug and compare it to `log2(number of commits)` — confirm it's close to that theoretical bound, connecting directly back to the binary search complexity analysis from the algorithms/discrete math tracks.

## Failure drill

Deliberately violate the bisect monotonicity assumption: in your test repo, introduce the bug, then later "fix" it partially, then reintroduce a *different* bug that also fails the same test. Run bisect again with the same good/bad range. Observe that bisect can now identify an incorrect or misleading commit as the culprit, because the "badness" of commits along the range is no longer monotonic (some commits in the middle pass the test, some later ones fail again for an unrelated reason). Explain in your own words why bisect's binary search correctness genuinely depends on the "broken-ness" being monotonic across the searched range, exactly like binary search over data depends on the data being sorted — bisect on non-monotonic history is the git-history equivalent of running binary search on unsorted data and getting a plausible-looking but wrong answer.

## Transfer

Look at your own commit history in Mahall, TARDOC, or Lead Sourcer and find one commit that bundles multiple unrelated changes together (a common pattern when moving fast on solo projects). State specifically how that bundled commit would degrade `git bisect`'s usefulness if a bug were later found to originate somewhere inside it, and what splitting it into focused commits would have bought you.

## Done when

You've used `git bisect` to correctly locate a deliberately introduced bug and confirmed the step count matches the expected binary-search bound, you've used `git log -p` and `git blame` together (not blame alone) to recover the actual reasoning behind an existing piece of code, and you can explain, using your own failure-drill result, why bisect can silently give a wrong answer when history isn't monotonically good-then-bad.
