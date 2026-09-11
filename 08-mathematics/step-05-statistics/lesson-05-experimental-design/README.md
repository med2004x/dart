# Lesson 5: Experimental Design

## Objective

Understand the principles that make an experiment (or A/B test) valid — randomization, control groups, statistical power, and confounding variables — closing the statistics track by tying together sampling, hypothesis testing, and regression into a coherent framework for actually running a trustworthy experiment.

## Prerequisites

Lesson 3 (hypothesis testing — experimental design is largely about ensuring the assumptions behind a hypothesis test actually hold in practice).

## Learn

**Confounding variables.** A confounder is a variable that affects both the treatment assignment and the outcome, creating a spurious apparent relationship between them even with no real causal effect. Classic example: ice cream sales and drowning deaths are correlated — not because ice cream causes drowning, but because both increase with hot weather (the confounder). Correlation from Lesson 03 of probability said nothing about causation; experimental design's entire purpose is to isolate causal effects despite this.

**Randomization** is the primary defense against confounding. Randomly assigning subjects to treatment and control groups ensures that, in expectation, confounders (known and unknown) are balanced between groups — any remaining difference in outcome is attributable to the treatment, not to some pre-existing difference between who ended up in which group. This is precisely why a randomized A/B test is more trustworthy than comparing "clinics that adopted TARDOC" against "clinics that didn't," where adoption itself likely correlates with confounders (tech-savviness, dissatisfaction with current tools) that also affect the outcome you're trying to measure.

**Control group.** A baseline group that doesn't receive the treatment, used for comparison. Without one, you can't distinguish "the treatment worked" from "things would have improved anyway" (regression to the mean, seasonal effects, general trends).

**Statistical power.** The probability of correctly detecting a real effect, if one exists — formally, `1 - P(Type II error)`, where a Type II error is failing to reject a false null hypothesis (missing a real effect). Power depends on sample size, effect size, and significance threshold. Lesson 3's step 3 vs step 4 comparison (n=1,000 vs n=10,000 for the same 1% true effect) was directly a power comparison: the larger sample had higher power to detect the same real effect. An underpowered study risks a false negative — concluding "no effect" when a real, meaningful effect existed but the sample was too small to reliably detect it.

**Type I vs Type II errors, stated together since confusing them is common:**
- *Type I error* (false positive): rejecting `H₀` when it's actually true — concluding there's an effect when there isn't. Controlled by your significance level (α, typically 0.05).
- *Type II error* (false negative): failing to reject `H₀` when it's actually false — missing a real effect. Controlled by power (1-β), which depends on sample size and effect size.

There's an inherent tradeoff: for a fixed sample size, reducing your Type I error rate (a stricter significance threshold) generally increases your Type II error rate, and vice versa — you cannot minimize both simultaneously without increasing sample size.

## Attempt

1. Design (on paper, no code needed) a randomized experiment to test whether a new onboarding email sequence increases user activation rate for Mahall. State explicitly: what's the treatment, what's the control, how would subjects be randomly assigned, what's the outcome metric, and what confounders randomization is meant to neutralize (e.g. time-of-signup effects, marketing-channel differences between users).

2. Using the two-proportion z-test setup from Lesson 3, estimate (via simulation, since exact power formulas require more machinery than covered here) the statistical power to detect a true 10% → 12% conversion rate improvement at `n=500` per group, by simulating this scenario 200 times (200 independent draws of two Bernoulli-sampled groups at the true rates) and recording what fraction of those 200 simulated experiments produce p < 0.05. This simulated fraction is your estimated power at that sample size.

3. Repeat step 2 at `n=2,000` and `n=5,000` per group, and confirm power increases with sample size for the same true effect — report all three estimated power values side by side.

4. Construct a deliberately confounded (non-randomized) comparison: simulate a scenario where "users who opted into a feature" have a different baseline behavior than "users who didn't," independent of the feature's real effect — e.g., give opted-in users a slightly higher baseline conversion rate purely due to self-selection (unrelated to the feature itself), then also apply the feature's true small effect only to the opted-in group, and show that a naive comparison overstates the feature's real effect because it's conflating the self-selection difference with the treatment effect.

## Verify

Report your three power estimates from steps 2-3 (should show a clear increasing trend with n, e.g. something like 15-30% power at n=500, 50-70% at n=2,000, 85%+ at n=5,000, though your exact numbers will vary with the random simulation — the trend direction is what matters, not hitting exact figures) and report the overstated effect size from step 4's confounded comparison versus the true effect size you built into the simulation.

## Failure drill

Take step 4's confounded scenario and now compare it to a properly randomized version of the same underlying situation (same true feature effect, but this time assign the baseline-boost randomly rather than correlating it with treatment assignment). Confirm the randomized version's estimated effect is much closer to the true effect than the confounded version's. This is the direct, simulated proof that randomization — not just a larger sample, not just a "cleaner-looking" comparison — is what removes the confound's distorting influence.

## Transfer

If you've ever compared outcomes between two groups of users, clinics, or sellers in TARDOC, Mahall, or Lead Sourcer without explicit random assignment (e.g. "power users vs. casual users," "early signups vs. late signups"), identify at least one plausible confounder that could be driving an apparent difference you observed, independent of whatever causal story you might have assumed, and state what a properly randomized experiment on the same question would need to look like.

## Done when

You can design a randomized experiment with an explicit treatment, control, and randomization scheme for a real question, you've simulated and observed statistical power increasing with sample size for a fixed effect, and you've demonstrated with your own simulation that a confounded comparison can produce a systematically wrong effect estimate that randomization corrects.
