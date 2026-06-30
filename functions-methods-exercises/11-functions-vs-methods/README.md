# Exercise 11 - Functions Vs Methods

## Goal

Practice deciding whether behavior belongs as a function or a method.

Use a method when the behavior clearly belongs to one type. Use a function when
the behavior combines multiple independent inputs or does not belong to one
receiver.

## Required Data

Create:

- `Temperature` type with Celsius value
- `WeatherReport` type with city and temperature

## Required Behavior

| Behavior | Use | Reason |
|---|---|---|
| Celsius to Fahrenheit | method on `Temperature` | depends on one temperature |
| Freezing check | method on `Temperature` | asks about one temperature |
| Build report line | function | combines city/report formatting |
| Compare two temperatures | function | uses two independent values |

## Main Requirements

In `main`:

1. Create at least three temperatures.
2. Create at least two weather reports.
3. Call both methods.
4. Call both functions.

## Constraints

- Do not make everything a method.
- Do not make everything a function.
- Be able to explain why each behavior is placed where it is.

## Prove It Works

Your output should show conversions, freezing status, report lines, and a
comparison result.
