# Project 02 - Requirements And Capacity

## Goal

Turn vague product language into measurable traffic, storage, latency, and
recovery requirements before choosing architecture.

## Starter

Run:

```powershell
.\capacity.ps1
```

The script calculates average requests, peak requests, and yearly storage from
explicit assumptions.

## Checkpoints

1. Complete `requirements-template.md` for a task service.
2. Model normal traffic.
3. Model a marketing peak.
4. Model ten-times user growth.
5. Add read/write ratio and largest response size.
6. Add database connection and worker-capacity budgets.
7. Identify the first likely bottleneck in each scenario.

## Required Scenarios

```text
normal:
    25,000 daily users
    30 requests per user
    12x peak multiplier

campaign:
    same users
    40x peak multiplier

growth:
    250,000 daily users
    30 requests per user
    12x peak multiplier
```

## Failure Drill

Design capacity from average RPS only. Compare it with campaign peak RPS and
explain the expected overload.

## Done Means

Every capacity claim has units, arithmetic, assumptions, and a measurement that
could confirm or reject it.

