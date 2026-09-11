# Lesson 5: Multivariable Calculus

## Objective

Extend derivatives to functions of several variables (partial derivatives, gradients), and implement multivariable gradient descent — the actual algorithm behind training any real machine learning model with more than one parameter.

## Prerequisites

Lesson 2 (single-variable derivatives), Lesson 2 of linear algebra (vectors — the gradient is a vector).

## Learn

**Partial derivatives.** For `f(x, y)`, the partial derivative with respect to `x`, written `∂f/∂x`, treats `y` as a constant and differentiates with respect to `x` alone, using the ordinary single-variable rules from Lesson 2. Symmetrically for `∂f/∂y`.

Worked example: for `f(x, y) = x²y + 3y²`,

`∂f/∂x = 2xy` (treat `y` as constant; power rule on `x²` gives `2x`, multiplied by the constant `y`)

`∂f/∂y = x² + 6y` (treat `x` as constant; `x²y` differentiates to `x²`, and `3y²` differentiates to `6y`)

**The gradient.** The gradient `∇f = (∂f/∂x, ∂f/∂y)` is a vector pointing in the direction of steepest ascent of `f` at a given point, with magnitude equal to the rate of increase in that direction. This is the direct multivariable generalization of the single-variable derivative from Lesson 2, and it's the exact quantity gradient descent uses when there's more than one parameter to optimize (which is every real machine learning model — a model with a million parameters has a million-dimensional gradient).

**Multivariable gradient descent.** Same idea as Lesson 2's single-variable version, but now the update is a vector operation: `x_new = x_old - α·∇f(x_old)`, moving the entire parameter vector in the direction opposite the gradient (steepest descent), scaled by learning rate `α`.

**Chain rule for multivariable functions** (used constantly in backpropagation): if `f` depends on `u` and `v`, and both `u` and `v` depend on `t`, then `df/dt = (∂f/∂u)(du/dt) + (∂f/∂v)(dv/dt)`. This is the rule that lets gradients "flow backward" through a computational graph with many intermediate variables — the entire mechanical content of backpropagation is this rule applied repeatedly through a chain of composed functions.

## Attempt

1. Compute `∂f/∂x` and `∂f/∂y` for `f(x, y) = 3x²y³ - 2x + y`. Show your work for both partials separately.

2. Compute the gradient `∇f` for `f(x, y) = x² + y²` (a paraboloid, minimum at the origin) at the points `(1, 1)`, `(2, 0)`, and `(0, -3)`. Verify each gradient vector points radially away from the origin (since this function increases fastest moving away from the minimum) — check this by confirming each gradient is a positive scalar multiple of the position vector at that point.

3. Implement multivariable gradient descent in Go or Python to minimize `f(x, y) = (x-3)² + (y+2)²` (minimum at `(3, -2)`, value 0). Gradient is `∇f = (2(x-3), 2(y+2))`. Start at `(0, 0)`, use `α = 0.1`, iterate at least 50 steps, and print `(x, y)` and `f(x,y)` every 10 iterations.

4. Modify step 3's function to `f(x, y) = (x-3)² + 10(y+2)²` — the same minimum, but the `y` direction is scaled up 10x, making the "bowl" much steeper in `y` than `x`. Run gradient descent with the same fixed learning rate and observe the behavior — it likely oscillates in `y` while converging slowly in `x`, or diverges in `y` if `α` is too large for that direction.

## Verify

Report the final `(x, y)` and `f(x,y)` from step 3 (should converge close to `(3, -2)`, `0`), and report what specifically goes wrong or converges slowly in step 4 — actual printed trajectory values, not just a description.

## Failure drill

For step 4's ill-conditioned function, try a much smaller learning rate (e.g. `α = 0.01`) and observe that it now converges in `y` without oscillating, but takes far more iterations to converge in `x`. Explain in 2-3 sentences why one fixed learning rate is fundamentally a compromise when different directions have very different curvature (second-derivative magnitude) — this is the actual motivation for adaptive learning rate methods (Adam, RMSprop) used in real ML training, which you don't need to implement here, just understand the problem they solve.

## Transfer

Take the single-variable Newton's method from Lesson 2's transfer task and describe (you don't need to implement it) how it generalizes to multiple dimensions: Newton's method in multiple variables uses the Hessian (matrix of second partial derivatives, covered conceptually here even though the full Hessian machinery belongs to a more advanced course) instead of a single second derivative, to account for curvature in every direction simultaneously rather than assuming one learning rate fits all directions equally, unlike plain gradient descent's failure mode in step 4.

## Done when

You can compute partial derivatives and assemble them into a gradient for a new function without help, your multivariable gradient descent converges correctly on a well-conditioned function, and you can explain — using your own step 4 results as evidence — why a single fixed learning rate struggles when a function's curvature differs sharply across dimensions.
