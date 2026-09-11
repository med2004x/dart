# Lesson 2: Random Variables and Distributions

## Objective

Formalize random variables, distinguish discrete from continuous distributions, and master expectation and variance — the two numbers that summarize a distribution's center and spread, used everywhere from latency SLOs to A/B test analysis.

## Prerequisites

Lesson 1 (probability foundations), Lesson 3 of the analysis track (integrals — continuous distributions require integrating a density function).

## Learn

**Random variable.** A function that assigns a number to each outcome of a random process. `X = number of heads in 3 coin flips` is a discrete random variable (values in a countable set). `X = time until a server request completes` is a continuous random variable (values in a continuous range).

**Discrete distributions** are described by a probability mass function (PMF), `P(X=k)`. Key examples:
- **Bernoulli**: single trial, success probability `p`. `P(X=1)=p, P(X=0)=1-p`.
- **Binomial**: number of successes in `n` independent Bernoulli trials. `P(X=k) = C(n,k)pᵏ(1-p)ⁿ⁻ᵏ` (note the combinatorics from discrete math Lesson 3 appearing directly in the formula).
- **Poisson**: number of events in a fixed interval, given a known average rate `λ`. `P(X=k) = e^(-λ)λᵏ/k!`. This is the standard model for "requests arriving per second" or "errors occurring per hour" — independent, rare, memoryless events.

**Continuous distributions** are described by a probability density function (PDF), `f(x)`, where `P(a ≤ X ≤ b) = ∫[a to b] f(x)dx` (directly using the integral machinery from the analysis track). Unlike discrete PMFs, `f(x)` itself is not a probability — only the integral over an interval is; `P(X = exact single value) = 0` for any continuous distribution.
- **Uniform**: equally likely over an interval `[a,b]`.
- **Exponential**: models waiting time until the next event in a Poisson process (e.g. time between server requests). `f(x) = λe^(-λx)` for `x ≥ 0`. Memoryless: the distribution of remaining wait time doesn't depend on how long you've already waited — a property worth checking your intuition against, since it's not how most physical waiting feels.
- **Normal (Gaussian)**: the bell curve, `f(x) = (1/√(2πσ²))e^(-(x-μ)²/(2σ²))`. Central to the Central Limit Theorem (Lesson 4).

**Expectation (mean).** `E[X] = Σ x·P(X=x)` (discrete) or `∫x·f(x)dx` (continuous) — the long-run average value.

**Variance.** `Var(X) = E[(X-E[X])²] = E[X²] - (E[X])²` — average squared deviation from the mean, quantifying spread. Standard deviation `σ = √Var(X)` is in the same units as `X` itself, which is why it's usually more interpretable than variance directly.

## Attempt

1. For `X = number of heads in 3 fair coin flips` (Binomial with n=3, p=0.5): write out the full PMF (`P(X=0), P(X=1), P(X=2), P(X=3)`), and compute `E[X]` directly from the definition (sum of `x·P(X=x)`). Confirm it matches the general binomial mean formula `E[X] = np`.

2. Compute `Var(X)` for the same distribution using `Var(X) = E[X²] - (E[X])²` (you'll need `E[X²] = Σx²P(X=x)`). Confirm it matches the general binomial variance formula `Var(X) = np(1-p)`.

3. A server receives requests at an average rate of 5 per second (Poisson, λ=5). Compute `P(exactly 3 requests in one second)` and `P(0 requests in one second)` using the Poisson PMF.

4. Implement a simulation in Go or Python: generate 100,000 samples of "3 coin flips, count heads" using a random number generator, compute the empirical mean and variance from the samples, and compare against your hand-computed theoretical `E[X]` and `Var(X)` from steps 1-2. They should be close (not identical — this is the Law of Large Numbers in action, previewing Lesson 4).

## Verify

Report your simulated mean/variance from step 4 alongside the theoretical values from steps 1-2, and confirm they agree to within a reasonable margin (with 100,000 samples, agreement to 2 decimal places is a reasonable expectation — if it's far off, check whether your random number generation or counting logic has a bug).

## Failure drill

Rerun the simulation from step 4 with only 10 samples instead of 100,000. Observe that the empirical mean/variance can be quite far from the theoretical values with so few samples. Run it several times with 10 samples and note how much the empirical mean varies run to run, versus how stable it becomes at 100,000. This is a direct, hands-on preview of why sample size matters for statistical estimation (formalized properly in the statistics track).

## Transfer

If TARDOC's Celery beat scheduler processes tasks (the hourly subscription sweep, or the nightly audio retention purge, mentioned in your project history) at some rate, describe how you'd model the arrival or processing of related events (e.g. incoming transcription jobs) using a Poisson process, and what estimating `λ` from real observed data would let you predict (e.g. `P(more than N jobs arrive in the next hour)`, relevant for capacity planning).

## Done when

You can compute the PMF, mean, and variance of a binomial distribution from the raw definitions (not just plugging into the shortcut formulas, though you should confirm they match), you can compute a Poisson probability for a given rate, and your simulation-vs-theory comparison in step 4 actually converges as you demonstrated in the failure drill.
