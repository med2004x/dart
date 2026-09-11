# Lesson 4: Limit Theorems

## Objective

Understand the Law of Large Numbers and the Central Limit Theorem — the two results that justify essentially all of statistical inference (Lesson 05 step) and explain why sample means behave predictably even when individual observations don't.

## Prerequisites

Lesson 2 (random variables, mean and variance), Lesson 4 of the analysis track (sequences and series — these are convergence results, same formal machinery).

## Learn

**Law of Large Numbers (LLN).** As the number of independent, identically distributed (i.i.d.) samples `n` grows, the sample mean `X̄ₙ = (X₁+...+Xₙ)/n` converges to the true population mean `μ = E[X]`. This is why Lesson 2's simulation (10,000 flips converging closer to the theoretical mean than 10 flips) worked the way it did — it's not a coincidence of that specific example, it's a general theorem. This is the formal justification for "more data gives a more reliable average," and it's why A/B tests need a minimum sample size before their results are trustworthy — with too few samples, the sample mean can be far from the true underlying rate purely by chance.

**Central Limit Theorem (CLT).** This is the more surprising and more useful result. Regardless of the shape of the original distribution `X` comes from (as long as it has finite mean and variance — doesn't need to be normal itself, doesn't even need to be symmetric), the distribution of the sample mean `X̄ₙ` approaches a **normal distribution** as `n` grows, specifically `X̄ₙ ≈ Normal(μ, σ²/n)`. This means the sample mean's variance shrinks proportionally to `1/n` — to halve your margin of error, you need 4x the sample size, not 2x, because standard deviation shrinks as `√n`, not `n`.

Why this is genuinely remarkable, not just a technical fact to memorize: it means you can construct confidence intervals and hypothesis tests using the normal distribution's well-understood properties, *even when the underlying data is nothing like normal* — heavily skewed conversion data, exponential wait times, whatever the raw distribution actually is — as long as you're looking at a sample mean (or sum) over a reasonably large sample. This is the mathematical reason the entire framework of Lesson 03/statistics (confidence intervals, hypothesis tests) works as broadly as it does.

## Attempt

1. Simulate the Law of Large Numbers directly: in Go or Python, roll a fair die repeatedly (or simulate any distribution with a known mean, e.g. `Uniform(1,6)`, true mean 3.5), and compute the running sample mean after every 1, 10, 100, 1000, 10000, 100000 rolls. Plot or print these running means and confirm they converge toward 3.5, with increasingly small fluctuations as n grows — not monotonically approaching, but bounded ever more tightly.

2. Simulate the Central Limit Theorem directly, using a deliberately *non-normal* starting distribution: generate samples from a skewed distribution (e.g. an exponential distribution, or simply `X = (uniform random number)²`, which is skewed). Take many samples (e.g. 10,000 repetitions) of "the mean of 30 draws" from this skewed distribution, and plot a histogram of those 10,000 sample means. Compare its shape to a histogram of the raw skewed distribution itself (single draws, not means).

3. From the same simulation in step 2, compute the standard deviation of your 10,000 sample means, and compare it to `σ/√30` where `σ` is the standard deviation of the original (single-draw) distribution — these should be close, confirming the CLT's variance-shrinkage prediction quantitatively, not just the shape claim.

4. Repeat step 2/3 with sample means of only 2 draws instead of 30 (still averaging, just over a much smaller sample), and observe the resulting histogram is noticeably less normal-shaped (closer to the original skewed shape) than the n=30 case — the CLT's normal approximation improves with larger n, and n=2 is a genuine stress case where the approximation is still weak.

## Verify

Report the actual computed standard deviation of your sample means from step 3, alongside the theoretical prediction `σ/√30`, and state the percent difference between them (should be small, single-digit percent, with 10,000 repetitions).

## Failure drill

Reduce your per-mean sample size in step 2 to `n=1` (i.e., you're not averaging anything, just looking at raw individual draws from the skewed distribution). Confirm the resulting histogram looks exactly like the original skewed distribution — no bell-curve shape at all. This is the boundary case that proves the bell shape in step 2 was genuinely produced by the averaging process (the CLT), not some property of your random number generator or plotting code.

## Transfer

If you've ever looked at aggregate metrics in TARDOC or Mahall (average transcription time, average order value) and treated them as roughly normally distributed for the purposes of eyeballing "is this unusual," state explicitly what sample size you were implicitly relying on for that normality to actually hold via the CLT, and whether individual raw events (a single transcription time, a single order) would show the same bell-shaped behavior — they generally would not, since the CLT applies to the sample mean, not to individual draws.

## Done when

You have directly simulated and observed both LLN convergence and CLT's shape-transformation-toward-normal, your step 3 comparison between simulated and theoretical standard deviation of the sample mean is quantitatively close, and you can explain in your own words why averaging (not the raw data itself) is what produces the normal shape — using your own n=1 failure-drill result as the concrete counter-example.
