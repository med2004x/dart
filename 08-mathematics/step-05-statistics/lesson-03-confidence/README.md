# Lesson 3: Confidence Intervals and Hypothesis Testing

## Objective

Construct confidence intervals correctly, run a hypothesis test (specifically for the A/B testing case you're most likely to actually use), and understand exactly what a p-value does and does not mean — the single most commonly misinterpreted concept in applied statistics.

## Prerequisites

Lesson 2 (sampling distributions — confidence intervals are built directly from the sampling distribution's standard error).

## Learn

**Confidence interval.** A range constructed from sample data such that, if you repeated the sampling process many times and built an interval the same way each time, roughly `X%` of those intervals would contain the true population parameter. For a sample mean with known/estimated standard error `SE = s/√n`, a 95% CI is approximately `x̄ ± 1.96·SE` (1.96 comes from the normal distribution's 97.5th percentile, since 95% CI leaves 2.5% in each tail).

**The single most important, most commonly wrong interpretation to avoid:** a 95% CI does *not* mean "there's a 95% probability the true mean is in this specific interval." The true mean is a fixed (if unknown) number — it either is or isn't in your specific interval, with no probability involved once the interval is computed. The 95% refers to the *procedure's* long-run reliability across repeated sampling, not to this one interval. This distinction feels pedantic until you need to explain a result to someone and realize the sloppy version is actually a different, false claim.

**Hypothesis testing.** Set up a null hypothesis `H₀` (typically "no effect" / "no difference") and an alternative `H₁`. Compute a test statistic from your data, and a **p-value**: the probability of observing data at least as extreme as what you got, *if the null hypothesis were true*. A small p-value (conventionally < 0.05) is evidence against `H₀`, leading you to "reject the null."

**What a p-value is NOT, stated explicitly because misreading this is extremely common:**
- It is *not* the probability that `H₀` is true.
- It is *not* the probability your result is "due to chance" in some general sense.
- A p-value of 0.03 does not mean "97% confident there's a real effect."

It is specifically: `P(data this extreme or more extreme | H₀ is true)`. This is a conditional probability computed under an assumption, not a statement about which hypothesis is actually correct — confusing these two is called the "prosecutor's fallacy" outside of stats, and it's genuinely easy to fall into even when you know the formal definition, because the informal-sounding version is so close to the correct one.

**A/B testing, concretely.** Comparing conversion rates between variant A (`p_A` observed) and variant B (`p_B` observed), each with sample size `n_A`, `n_B`. The standard approach is a two-proportion z-test: compute the standard error of the difference `p_A - p_B` under the null (no true difference), compute a z-statistic, and derive a p-value from the normal distribution. Critically: this requires a pre-determined sample size before looking at results repeatedly — checking results continuously and stopping "when it looks significant" (a very common real mistake, sometimes called "peeking") inflates the false-positive rate well above the nominal 5%, because you're implicitly running many tests, not one.

## Attempt

1. For a sample of `n=100` with sample mean `x̄=52` and sample standard deviation `s=10`, compute a 95% confidence interval for the true population mean using `x̄ ± 1.96·(s/√n)`.

2. Run a one-sample hypothesis test: you claim a website's average session length is 3 minutes. You collect `n=50` sessions with sample mean `2.7` minutes and sample standard deviation `0.8` minutes. Compute the z-statistic `(x̄ - μ₀)/(s/√n)` and the corresponding two-tailed p-value (you can use a standard normal table, or compute it in code using an error-function-based normal CDF). State whether you'd reject `H₀: μ=3` at the 0.05 significance level.

3. Simulate an A/B test in Go or Python: variant A has a true conversion rate of 10%, variant B has a true conversion rate of 11% (a genuine, small real effect you've built into the simulation). Simulate `n=1,000` users per variant (random Bernoulli draws at each true rate), compute the observed conversion rates, and run a two-proportion z-test comparing them. Report the p-value and whether you'd conclude a significant difference at n=1,000.

4. Repeat the simulation from step 3 at `n=10,000` per variant instead of 1,000 (same true rates, 10% vs 11%). Compare the p-values and significance conclusions between the two sample sizes — this demonstrates directly that the same true effect can be statistically significant or not purely depending on sample size, which connects back to why underpowered A/B tests (too few users) can miss real, meaningful effects.

## Verify

Report the actual p-values from steps 3 and 4 side by side, and explicitly state whether each was below 0.05 — a real 1% absolute difference (10% vs 11%) is a common size of effect that's genuinely hard to detect reliably at n=1,000 but easier at n=10,000; if your specific random simulation run doesn't show this pattern, rerun it a few times and note the variability, since a single simulation run is itself just one sample.

## Failure drill

Take your step 3 simulation (n=1,000) and run it repeatedly (e.g. 20 independent simulation runs, each a fresh random draw) even though the *true* underlying conversion rates never change between runs. Record how many of the 20 runs produce a p-value below 0.05 by chance alone in either direction, including runs where you set both variants to the *same* true rate (a true null case) — you should see roughly 1 in 20 runs falsely reject the null even when there's genuinely no difference, matching the 5% significance level's actual meaning: a 5% false-positive rate under repeated true-null testing, not a 5% chance of "getting it wrong" for any single test as commonly misstated.

## Transfer

If you've run or plan to run any kind of A/B comparison in Mahall or TARDOC (two pricing pages, two email subject lines, two onboarding flows), state what sample size you'd need to reliably detect an effect of a given size, using this lesson's n=1,000 vs n=10,000 contrast as a rough calibration point, and state explicitly why checking results daily and stopping as soon as p<0.05 first appears (rather than committing to a sample size in advance) would inflate your real false-positive rate above the nominal 5%.

## Done when

You can construct a confidence interval and correctly state what it does and does not claim, you can run a hypothesis test and compute a p-value from raw data, and you can explain — using your own failure-drill results, not just the abstract warning — why repeatedly testing until you get p<0.05 is a real statistical error, not just poor practice.
