# Lesson 2: Sampling

## Objective

Understand sampling distributions, sampling bias, and the difference between population and sample — the conceptual foundation that makes confidence intervals (Lesson 3) and hypothesis tests meaningful rather than just formula-plugging.

## Prerequisites

Lesson 4 of probability (Central Limit Theorem — the sampling distribution of the mean is a direct CLT application), Lesson 1 (descriptive statistics — sample statistics are computed the same way, now with sampling uncertainty layered on top).

## Learn

**Population vs. sample.** The population is the entire group you actually care about (e.g. all Swiss outpatient clinics). A sample is a subset you actually observe/measure. Nearly all real statistics is about using a sample to infer something about a population you cannot fully observe.

**Sampling distribution.** If you repeatedly drew samples of size `n` from the same population and computed the sample mean each time, those sample means themselves form a distribution — the *sampling distribution of the mean*. Lesson 4 of probability (CLT) tells you this distribution is approximately normal, centered at the true population mean, with standard deviation `σ/√n` (called the *standard error*). This is the object confidence intervals in Lesson 3 are built from directly.

**Sampling bias.** A sample that isn't representative of the population you're trying to draw conclusions about, no matter how large. Common real forms:
- *Selection bias*: your sampling method systematically favors certain outcomes (e.g. surveying only people who respond to emails — response itself correlates with the thing you're measuring).
- *Survivorship bias*: only observing units that "survived" some filtering process (e.g. only analyzing customers who didn't churn, when churn itself is what you wanted to study).
- *Self-selection bias*: participants choose whether to be in the sample, and that choice correlates with the outcome (e.g. only clinics that proactively reach out to try TARDOC may be systematically different — more tech-forward, more frustrated with existing tools — from the broader clinic population).

**Why bias is a fundamentally different problem than "small sample size."** More data does not fix bias — the Law of Large Numbers guarantees convergence to the *sampled* population's true mean, not the population you actually wanted to study. A biased sample of 1,000,000 is not better than an unbiased sample of 100 for answering the real question; it converges confidently to the wrong answer. This is a genuinely important distinction, since "get more data" is the reflexive fix for noise but does nothing for bias.

**Random sampling** (every population member has a known, typically equal, probability of selection) is the standard defense against selection bias, though it's often harder to achieve in practice than the textbook description suggests — you can only randomly sample from people/things that are actually reachable by your sampling method in the first place.

## Attempt

1. Simulate sampling distribution directly: in Go or Python, define a population as a large array of, say, 100,000 values drawn from some distribution with a known mean (e.g. `Normal(50, 10)` or any distribution you can generate). Draw 1,000 independent samples of size `n=30` from this population, compute the sample mean of each, and plot/examine the resulting distribution of those 1,000 sample means. Confirm it's centered near the true population mean and its spread is close to `σ/√30`.

2. Repeat step 1 with `n=5` and `n=100` (keeping 1,000 repetitions each time). Compare the spread (standard deviation) of the resulting sampling distributions across the three sample sizes, and confirm the `1/√n` shrinkage pattern holds — going from n=5 to n=100 (20x) should shrink the standard error by a factor of `√20 ≈ 4.47`, not 20x.

3. Construct a deliberately biased sampling scenario: from your population of 100,000 values, instead of sampling uniformly at random, sample only from the top 20% of values (a stand-in for "only clinics that proactively respond to outreach," if you like). Compute the mean of this biased sample of size 30, repeated 1,000 times, and compare its distribution to the correctly-centered one from step 1 — it should be clearly and consistently shifted, not just noisier.

4. Increase the biased sample's size dramatically (e.g. n=10,000 drawn the same biased way) and confirm the bias does *not* shrink or disappear with more data — the biased sampling distribution should still be centered at the wrong value, just with a tighter spread around that wrong value. This is the direct, hands-on demonstration of "bias isn't fixed by sample size."

## Verify

Report the actual means of your three sampling distributions from steps 1-2 (should all be close to the true population mean) and the actual mean of your biased sampling distribution from steps 3-4 (should be clearly, persistently offset, and should stay offset even as n grows in step 4).

## Failure drill

Take your step 4 large biased sample and compute a naive 95% confidence interval around it (you can use the standard formula from Lesson 3, or simply note that a tight, confident-looking interval will result, since bias reduces spread just like more unbiased data would). Confirm this confidence interval does *not* contain the true population mean, despite your having 10,000 "data points" and a seemingly tight, confident-looking result. Explain in your own words why a large biased sample can produce a highly confident-looking but wrong answer — this is a genuinely dangerous failure mode because it doesn't look uncertain, it looks precise.

## Transfer

If TARDOC's or Mahall's early user base (the clinics or sellers who signed up first) might differ systematically from the broader target population (e.g., early adopters tend to be more tech-forward or more actively dissatisfied with their current tools than the median clinic/seller), state explicitly what kind of bias this is (selection, self-selection, or survivorship) and what conclusion you should be cautious about generalizing from early-user feedback or metrics to the full target market.

## Done when

You've directly simulated a sampling distribution and confirmed the `1/√n` standard error shrinkage numerically, and you've demonstrated for yourself (not just read) that a biased sample stays wrong regardless of size, and you can identify a plausible source of sampling bias in your own actual business context.
