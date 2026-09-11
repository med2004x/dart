# Lesson 3: Joint Distributions

## Objective

Extend probability to multiple random variables simultaneously (joint, marginal, conditional distributions), understand covariance and correlation, and connect this directly to A/B testing pitfalls and feature correlation in data analysis.

## Prerequisites

Lesson 2 (single random variables — joint distributions are the multivariable extension, same relationship as multivariable calculus to single-variable).

## Learn

**Joint distribution.** For two random variables `X` and `Y`, the joint PMF `P(X=x, Y=y)` (discrete) or joint PDF `f(x,y)` (continuous) describes their behavior together, capturing any dependence between them.

**Marginal distribution.** Recovering the distribution of `X` alone from the joint: `P(X=x) = Σ_y P(X=x, Y=y)` (sum out `Y`) — this is exactly analogous to integrating out a variable in multivariable calculus.

**Conditional distribution.** `P(Y=y | X=x) = P(X=x,Y=y)/P(X=x)` — the distribution of `Y` restricted to the world where `X=x`, directly generalizing Lesson 1's conditional probability from events to random variables.

**Independence of random variables.** `X` and `Y` are independent if `P(X=x,Y=y) = P(X=x)·P(Y=y)` for all `x,y` — knowing one tells you nothing about the other.

**Covariance.** `Cov(X,Y) = E[(X-E[X])(Y-E[Y])] = E[XY] - E[X]E[Y]`. Positive covariance means X and Y tend to move together (both above or both below their means at the same time); negative means they move oppositely; zero (for independent variables, though the converse isn't generally true) means no linear relationship.

**Correlation.** `ρ(X,Y) = Cov(X,Y) / (σ_X·σ_Y)`, normalized to `[-1, 1]` — this normalization is what makes correlation comparable across variables with different units and scales, unlike raw covariance.

**A critical warning worth internalizing, not just noting:** correlation measures only *linear* association. Two variables can be strongly, deterministically related (e.g. `Y = X²`) and still have correlation near zero, because that relationship isn't linear. "No correlation" is not the same claim as "no relationship" — checking correlation alone can badly mislead you about whether two variables are actually connected.

## Attempt

1. Roll two fair dice, `X` = first die, `Y` = sum of both dice. Construct the joint PMF table `P(X=x, Y=y)` for all valid `(x,y)` pairs (X ∈ {1..6}, Y ∈ {2..12}). From this table, compute the marginal `P(Y=y)` for each y by summing out X, and confirm your marginal matches the standard "sum of two dice" distribution you may already know (e.g. `P(Y=7) = 6/36`).

2. Determine whether `X` (first die) and `Y` (sum) are independent, using your joint table from step 1 — pick one specific `(x,y)` pair and check whether `P(X=x,Y=y) = P(X=x)·P(Y=y)` holds; it should not, since knowing the first die clearly constrains the possible sums.

3. Compute `Cov(X,Y)` for the dice example from step 1, using the joint PMF to compute `E[XY]`, and `E[X]`, `E[Y]` from the marginals. Then compute the correlation `ρ(X,Y)`.

4. Implement a synthetic demonstration in Go or Python: generate data where `Y = X² + small_noise` for `X` uniformly distributed over `[-5, 5]` (at least 1000 samples), compute the Pearson correlation coefficient between X and Y, and observe it comes out close to 0 despite `X` and `Y` being deterministically related (up to noise). This is a direct, hands-on demonstration of the "zero correlation ≠ no relationship" warning above.

## Verify

For step 4, report the actual computed correlation coefficient (should be close to 0, likely in the range -0.1 to 0.1 depending on your noise and sampling) alongside a simple scatter description or plot showing the clear parabolic relationship — the contrast between "correlation says ~0" and "the data is visibly, strongly related" is the point to demonstrate concretely.

## Failure drill

Take the same synthetic data from step 4 but restrict `X` to only positive values, e.g. `X` uniform over `[0, 5]` instead of `[-5, 5]`, keeping `Y = X² + noise`. Recompute the correlation. It should now be strongly positive (close to 1), because on this restricted range the relationship is monotonic (as X increases, Y increases), even though the underlying function `Y=X²` hasn't changed at all — only the domain you're sampling from. Explain in your own words why restricting the domain changed the correlation so drastically despite the true relationship being identical, and what this implies about correlation being sensitive to the range of your data, not just the underlying relationship.

## Transfer

If you've ever looked at two metrics in TARDOC or Mahall (e.g. clinic size and monthly billing volume, or product price and conversion rate) and eyeballed whether they seemed related, describe how you'd now check that more rigorously with correlation — and, given this lesson's warning, what you'd do differently if the correlation came back near zero but you still suspected some (possibly nonlinear) relationship existed.

## Done when

You can construct a joint distribution table and correctly derive marginals and conditionals from it, you can compute covariance and correlation from a joint distribution or from data, and you can explain — using your own step 4/failure-drill results as evidence, not just the abstract warning — why near-zero correlation does not imply no relationship.
