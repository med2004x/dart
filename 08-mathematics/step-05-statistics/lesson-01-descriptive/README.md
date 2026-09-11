# Lesson 1: Descriptive Statistics

## Objective

Master the core descriptive statistics (measures of center and spread) beyond just "mean and standard deviation," understand when each is appropriate, and recognize how outliers distort naive summaries.

## Prerequisites

Lesson 2 of probability (mean and variance as properties of a random variable — this lesson is their empirical, data-driven counterpart).

## Learn

**Measures of center.**
- **Mean** `x̄ = (Σxᵢ)/n` — sensitive to outliers; a single extreme value can shift it substantially.
- **Median** — the middle value when sorted (or average of the two middle values for even n) — robust to outliers, since it only depends on rank order, not magnitude.
- **Mode** — the most frequent value; most useful for categorical or discrete data.

**Measures of spread.**
- **Range** = max - min — extremely sensitive to outliers (a single extreme point sets it entirely).
- **Variance** `s² = Σ(xᵢ-x̄)²/(n-1)` — note the `n-1` (not `n`) for a *sample* variance; this is Bessel's correction, which makes the sample variance an unbiased estimator of the true population variance (using `n` would systematically underestimate it, since the sample mean is itself fit to the data, reducing the apparent spread around it).
- **Standard deviation** `s = √s²` — same units as the data, generally more interpretable than variance.
- **Interquartile range (IQR)** = Q3 - Q1 (75th percentile minus 25th percentile) — robust to outliers, since it ignores the extreme tails entirely.

**Why mean vs. median matters practically, not just as trivia:** income and latency data are both classically right-skewed (most values clustered low, with a long tail of high values). Reporting "average response time" when the distribution is skewed can be badly misleading — a small number of very slow requests drags the mean up even if most requests are fast — which is exactly why real-world latency metrics are reported as percentiles (p50, p95, p99) rather than a single mean. p50 is the median; p95/p99 explicitly report the tail behavior the mean obscures.

**Outlier sensitivity, concretely.** Consider response times (ms): `[10, 12, 11, 13, 9, 500]`. Mean ≈ 92.5 ms — dominated by the single 500ms outlier, and arguably describes none of the six actual requests well. Median = 11.5 ms — much more representative of "typical" performance, completely unaffected by how extreme that one outlier is.

## Attempt

1. For the dataset `[4, 8, 15, 16, 23, 42]`, compute the mean, median, sample variance (with `n-1`), sample standard deviation, and range by hand.

2. For the response-time example in Learn (`[10, 12, 11, 13, 9, 500]`), compute mean and median by hand, and explicitly compare them to the "typical" value a human would intuitively describe this data as having — argue in one sentence for which measure better represents this specific dataset and why.

3. Compute the IQR for the sorted dataset `[2, 4, 4, 4, 5, 5, 7, 9, 12, 21]` (find Q1 and Q3 using whichever standard method you choose — e.g. the median of the lower/upper half — and state which method you used, since there are multiple valid conventions).

4. Implement a small statistics library in Go or Python (`mean`, `median`, `variance`, `stddev`, `percentile(data, p)`) from scratch (no external stats library), and use it to compute p50, p95, and p99 of a synthetic "latency" dataset you generate: 950 values around 10-15ms plus 50 outlier values around 200-500ms (simulating a realistic latency distribution with a long tail). Report all three percentiles plus the mean, and compare how much the mean is dragged up by the tail versus how stable p50 remains.

## Verify

Your from-scratch functions' outputs for step 1 must exactly match your hand-computed values. For step 4, explicitly report: mean, p50, p95, p99 — and confirm p50 is close to the bulk of your "normal" 10-15ms data while the mean is measurably higher due to the tail, demonstrating the point from Learn with your own generated numbers rather than the abstract example.

## Failure drill

Compute the sample variance for a small dataset (e.g. `[1, 2, 3, 4, 5]`) using `n` in the denominator instead of the correct `n-1`, and compare to the correct value. The difference is small for large n but proportionally larger for small n (as in this 5-element example) — compute both explicitly, note the percentage difference, and explain in one sentence why systematically using `n` would make your estimate of population variance slightly too small on average, connecting to why Bessel's correction exists rather than just being an arbitrary convention.

## Transfer

If you monitor response times, processing durations, or queue wait times anywhere in TARDOC or Mahall (or plan to), state which summary statistic (mean, median, p95, p99) is actually appropriate for whatever SLA or performance goal you care about, and why reporting only a mean would risk hiding exactly the kind of tail behavior that matters most for user-facing latency.

## Done when

You can compute all the core descriptive statistics by hand and verify your code implementation matches, and you can explain — using your own step 4 numbers as evidence, not just citing the general principle — why percentiles are preferred over the mean for skewed, real-world latency-like data.
