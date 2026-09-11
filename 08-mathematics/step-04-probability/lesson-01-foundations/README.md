# Lesson 1: Probability Foundations

## Objective

Build probability from its axioms (not just intuition), master conditional probability and Bayes' theorem, and connect Bayes' theorem to real applications you'll actually use (spam/fraud detection, A/B test interpretation).

## Prerequisites

Lesson 3 of discrete math (combinatorics — counting is how you compute probabilities for equally-likely outcomes).

## Learn

**Axioms (Kolmogorov).** A probability space assigns `P(A) ∈ [0,1]` to events `A` in a sample space, with `P(sample space) = 1`, and for mutually exclusive events, `P(A∪B) = P(A) + P(B)`. Everything else in probability — every formula you'll ever use — is derivable from these three simple rules.

**Conditional probability.** `P(A|B) = P(A∩B) / P(B)` (for `P(B) > 0`) — the probability of A, given that B has already happened, restricted to the world where B is true.

**Independence.** Events `A` and `B` are independent if `P(A∩B) = P(A)·P(B)`, equivalently `P(A|B) = P(A)` (knowing B happened tells you nothing about A). Independence is an assumption you make about a model, not something you can prove from data alone without care — a very common real-world error is assuming independence when events are actually correlated.

**Bayes' theorem.** `P(A|B) = P(B|A)·P(A) / P(B)`. This single formula lets you flip a conditional probability around — critical because the probability you can measure (e.g. `P(test positive | has disease)`, from clinical trial data) is often not the one you actually want (`P(has disease | test positive)`, what a patient actually needs to know).

Worked example (the classic, worth internalizing fully): a disease affects 1% of a population. A test is 99% accurate (99% true positive rate, 99% true negative rate). Given a positive test, what's the probability of actually having the disease?

```
P(disease) = 0.01, P(no disease) = 0.99
P(positive | disease) = 0.99
P(positive | no disease) = 0.01  (false positive rate)

P(positive) = P(positive|disease)P(disease) + P(positive|no disease)P(no disease)
            = 0.99·0.01 + 0.01·0.99
            = 0.0099 + 0.0099 = 0.0198

P(disease | positive) = P(positive|disease)P(disease) / P(positive)
                       = 0.0099 / 0.0198 = 0.5
```

Despite a "99% accurate" test, a positive result only means a 50% chance of actually having the disease — because the disease is rare, false positives from the large healthy population roughly match true positives from the small sick population. This is the single most important intuition-correcting result in basic probability, and it directly explains why rare-event detection systems (fraud, spam, intrusion detection) need very low false-positive rates or a large prior, not just "high accuracy," to be actually useful.

## Attempt

1. A fair six-sided die is rolled. Compute `P(roll is even)`, `P(roll is even | roll > 3)`, and determine whether "roll is even" and "roll > 3" are independent (compute `P(even)·P(>3)` and compare to `P(even ∩ >3)`).

2. Redo the disease-testing worked example with a rarer disease (0.1% prevalence instead of 1%) and the same 99% test accuracy. Compute `P(disease | positive)` and note how much lower it drops — this demonstrates that rarity of the event being detected matters enormously, independent of how "accurate" the test sounds.

3. A spam filter classifies emails: 20% of all email is spam. The filter flags 95% of actual spam as spam (true positive rate), and incorrectly flags 5% of legitimate email as spam (false positive rate). Given an email flagged as spam, compute `P(actually spam | flagged)` using Bayes' theorem, following the disease-testing structure exactly.

4. Implement a small Bayes' theorem calculator in Go or Python that takes `P(A)`, `P(B|A)`, `P(B|¬A)` as inputs and returns `P(A|B)`. Verify it against your hand-computed answers from steps 2 and 3.

## Verify

Your calculator's output for steps 2 and 3 must match your hand-computed values exactly (within floating-point rounding). Report both hand and code results side by side.

## Failure drill

Compute `P(disease | positive)` for the original 1% prevalence example but using a badly wrong "intuitive" shortcut: just report the test's stated accuracy, 99%, as if it were the answer. Compare that wrong 99% to the correct 50% from the worked example. State explicitly, in your own words, what's being conflated — the wrong shortcut mistakes `P(positive|disease)` (the test's sensitivity, a property of the test) for `P(disease|positive)` (what you actually want, which also depends on the disease's prevalence) — these are different conditional probabilities and Bayes' theorem is precisely the tool that keeps them from being confused.

## Transfer

If Lead Sourcer's lead-scoring pipeline (mentioned in your project history — pain point detection, outreach angle scoring) makes any kind of "is this a good lead" classification, describe how you'd frame it in Bayes' terms: what would play the role of the "prior" (`P(good lead)` before seeing any signal), and what would play the role of the "test" (some measurable signal correlated with lead quality) — and whether a rare positive class (few actually-good leads among many scored) creates the same false-positive-swamping effect as the disease example.

## Done when

You can compute conditional probabilities and check independence from a joint probability table, you can apply Bayes' theorem correctly to flip a conditional probability, and you can explain from first principles (not by memorized rule) why a "99% accurate" test can still be a coin flip's worth of informative when the underlying condition is rare.
