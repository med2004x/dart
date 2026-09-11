# Lesson 2: Derivatives

## Objective

Derive the definition of the derivative from limits, compute derivatives from first principles and from rules, and connect derivatives to gradient-based optimization used in real software (gradient descent, Newton's method).

## Prerequisites

Lesson 1 (limits — the derivative is defined as a limit).

## Learn

**Definition.** The derivative of `f` at `x` is `f'(x) = lim(h→0) [f(x+h) - f(x)] / h`, when this limit exists. Geometrically: the slope of the tangent line to `f` at `x`. Physically: the instantaneous rate of change.

Worked example: derive `f'(x)` for `f(x) = x²` from first principles.

```
f'(x) = lim(h→0) [(x+h)² - x²] / h
      = lim(h→0) [x² + 2xh + h² - x²] / h
      = lim(h→0) [2xh + h²] / h
      = lim(h→0) (2x + h)      [valid for h ≠ 0, cancel h]
      = 2x
```

This matches the power rule `d/dx[xⁿ] = n·xⁿ⁻¹` for n=2.

**Rules (proved from the definition, used in practice):**
- Power rule: `d/dx[xⁿ] = n·xⁿ⁻¹`
- Sum rule: `(f+g)' = f' + g'`
- Product rule: `(fg)' = f'g + fg'`
- Chain rule: `(f∘g)'(x) = f'(g(x))·g'(x)` — this is the one that matters most for the transfer task below.

**Why the chain rule specifically matters to you:** backpropagation in neural networks is repeated application of the chain rule through a computational graph — every "gradient" a training loop computes is a chain-rule derivative. You don't need to implement a neural net for this lesson, but recognizing that backprop is not a separate algorithm from calculus, it's calculus applied mechanically to a composed function, changes how you read ML code.

**Gradient descent, concretely.** To minimize a function `f(x)`, move in the direction that decreases `f` fastest, which is the negative direction of the derivative (for multivariable functions, the negative gradient — covered in Lesson 5). Update rule: `x_new = x_old - α·f'(x_old)`, where `α` is a step size (learning rate). This is the entire idea behind training a model — the loss function is `f`, the model parameters are `x`, and gradient descent iteratively nudges parameters downhill.

## Attempt

1. Derive `f'(x)` from first principles (the limit definition, not the power rule shortcut) for `f(x) = x³`. Show every algebraic step, matching the worked example's structure.

2. Compute derivatives using rules (state which rule you used for each):
   - `f(x) = 3x⁴ - 2x² + 5`
   - `f(x) = x²·sin(x)` (product rule)
   - `f(x) = sin(x²)` (chain rule — do not confuse with the previous one)

3. Implement gradient descent in Go or Python to minimize `f(x) = (x - 3)² + 1` (minimum at x=3, value 1). Start at `x = 0`, use `f'(x) = 2(x-3)`, pick a learning rate (start with `α = 0.1`), and iterate the update rule for at least 50 steps. Print `x` and `f(x)` every 10 iterations.

4. Rerun step 3 with a learning rate that's too large (e.g. `α = 1.1`) and observe divergence instead of convergence. Then try one that's very small (e.g. `α = 0.001`) and observe slow convergence. Record what happens in both cases.

## Verify

Your gradient descent implementation from step 3 should converge to `x ≈ 3.0` and `f(x) ≈ 1.0` within your iteration budget at a reasonable learning rate. Report the actual final values and how many iterations it took to get within 0.01 of the true minimum.

## Failure drill

In step 4's diverging case (α too large), print `x` at every single iteration (not every 10th) and observe it oscillating with increasing amplitude rather than settling. Explain in one or two sentences, using the update rule algebraically, why a learning rate above a certain threshold causes overshoot that gets worse each step rather than better, for this specific quadratic.

## Transfer

Take Newton's method for root-finding — `x_new = x_old - f(x_old)/f'(x_old)` — and use it to find a root of `f(x) = x² - 2` (i.e., approximate √2) starting from `x_0 = 1`. Implement it and iterate until the value stabilizes to several decimal places. Compare its convergence speed (iterations needed) to gradient descent's from step 3, and note that Newton's method uses the derivative differently (to jump toward a root, not just downhill).

## Done when

You can derive a polynomial derivative from the limit definition without help, correctly apply product/chain rules to a function combining both, and have a working gradient descent implementation whose divergence/convergence behavior you can explain in terms of the learning rate and the derivative's magnitude.
