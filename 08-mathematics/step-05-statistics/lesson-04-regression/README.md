# Lesson 4: Regression

## Objective

Extend the least-squares mechanics from the linear algebra track into a full statistical treatment: fit a regression, evaluate it properly (R², residual analysis), and understand its assumptions well enough to know when a fitted line is misleading.

## Prerequisites

Lesson 4 of linear algebra (least squares mechanics — this lesson adds the statistical interpretation and diagnostics on top of the computation already covered there).

## Learn

**Recap and reframe.** Linear algebra Lesson 4 showed how to compute the best-fit line `y = mx + c` by solving the normal equations. This lesson asks the next-level questions: how good is this fit, is a straight line even appropriate, and how much should you trust the fitted coefficients.

**R² (coefficient of determination).** `R² = 1 - (SS_res / SS_tot)`, where `SS_res = Σ(yᵢ - ŷᵢ)²` (sum of squared residuals — actual minus predicted) and `SS_tot = Σ(yᵢ - ȳ)²` (total variance around the mean, with no model at all). `R²` ranges from 0 to 1 (can be negative for a genuinely bad fit) and represents the proportion of variance in `y` explained by the model. `R² = 1` is a perfect fit; `R² = 0` means the model explains nothing beyond just predicting the mean every time.

**Residual analysis.** The residuals `eᵢ = yᵢ - ŷᵢ` should, for a well-specified linear model, look like unstructured random noise when plotted against `x` or against `ŷ` — no visible pattern, no curve, no fanning-out shape (heteroscedasticity — variance that changes systematically with x). A visible curved pattern in residuals is strong evidence the true relationship isn't linear, even if `R²` looks decent; a fanning pattern (residuals growing wider as x or ŷ increases) violates the constant-variance assumption most standard regression inference relies on.

**A crucial warning, worth internalizing at least as much as the mechanics:** a high `R²` does not mean the model is correctly specified, and a low `R²` does not mean there's no relationship — both can be misled by nonlinearity, outliers, or the wrong functional form entirely. This is exactly the linear-algebra Lesson 4 failure drill (fitting a line to `y=x²` data) revisited with a proper diagnostic: the residual plot for that case would show an obvious curve, which is the correct way to catch the problem, rather than relying on `R²` alone.

**Multiple regression.** The same normal-equations machinery extends directly to multiple predictors: `y = β₀ + β₁x₁ + β₂x₂ + ... + βₖxₖ`, fit by the same `AᵀAx = Aᵀb` structure with `A`'s columns now being multiple features plus an intercept column. Interpreting individual coefficients requires care: `βᵢ` represents the effect of `xᵢ` holding all other predictors constant, which is only a meaningful, stable interpretation if the predictors aren't too correlated with each other (multicollinearity, briefly worth knowing the name of even without deriving its full treatment here).

## Attempt

1. Using the least-squares fit you already computed by hand in linear algebra Lesson 4 (the points `(0,1),(1,3),(2,4),(3,6)`), compute `R²` explicitly: find `ŷᵢ` for each `xᵢ` using your fitted line, compute `SS_res` and `SS_tot`, and derive `R²`.

2. Extend your linear-algebra least-squares code (Go or Python) to also compute and return `R²`, and verify it against your hand computation from step 1.

3. Plot or print the residuals (`yᵢ - ŷᵢ`) for the `y=x²`-data failure-drill case from linear algebra Lesson 4 (points sampled from `y=x²` over `x ∈ {-3,...,3}`, fit with a straight line). Confirm the residuals show a clear U-shaped or curved pattern when plotted against `x`, rather than looking like unstructured noise — this is the diagnostic that would have caught the model misspecification even without inspecting `R²` directly.

4. Fit a multiple regression with 2 predictors on synthetic data you generate: `y = 2x₁ + 3x₂ + 1 + small_noise`, with `x₁` and `x₂` drawn independently (so no multicollinearity concern here). Use at least 50 data points, build the design matrix with 3 columns (`x₁`, `x₂`, and a column of 1s for the intercept), solve the normal equations, and confirm your fitted coefficients come out close to the true values (2, 3, 1).

## Verify

Report your computed `R²` from steps 1-2 (should match between hand and code), and report your fitted 3-coefficient vector from step 4 alongside the true values (2, 3, 1) — they should be close, with the gap explained by your added noise level.

## Failure drill

Take the multiple regression from step 4, but this time generate `x₂` as a near-exact linear function of `x₁` (e.g. `x₂ = 2·x₁ + tiny_noise`) instead of independently — this creates multicollinearity. Refit the model. Observe that the individual fitted coefficients for `x₁` and `x₂` can become unstable, unreliable, or nonsensical (e.g. one very large positive and one very large negative, even though the true values were 2 and 3) even though the model's overall predictions (and R²) may still look fine. Explain in your own words why this happens: when two predictors are highly correlated, the model can't reliably attribute effect to one versus the other, since many different coefficient combinations fit the data almost equally well.

## Transfer

If TARDOC's or Mahall's pricing, billing, or scoring logic ever fits (formally or informally) a relationship between an input and an outcome (e.g. clinic size predicting billing volume, or ad spend predicting conversions), state what you'd check before trusting a simple linear fit: what would you look for in a residual plot, and what would make you suspect the true relationship isn't linear even if a naive R² looked acceptable.

## Done when

You can compute R² by hand and in code and correctly interpret what it does and doesn't tell you, you've produced a residual plot that visibly reveals model misspecification on the y=x² case, and you've demonstrated multicollinearity's effect on coefficient stability with your own generated data.
