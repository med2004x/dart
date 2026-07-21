# Functions And Methods Quick Reference

Read this before the numbered exercise when the syntax is unfamiliar. The
examples use different names so they do not solve an exercise for you.

## Function Forms

```go
func printLine(text string) {
	fmt.Println(text)
}

func add(left, right int) int {
	return left + right
}

func lookup(values []string, wanted string) (int, bool) {
	for index, value := range values {
		if value == wanted {
			return index, true
		}
	}
	return 0, false
}
```

The caller decides what to do with returned values. A function cannot change a
caller variable just because it received a copy of that value.

## Slices: Return Updated State

```go
func addTag(tags []string, tag string) []string {
	return append(tags, tag)
}

tags = addTag(tags, "urgent")
```

When a function returns a new slice, assign the result. For an element update,
use an index rather than the range value copy.

## Methods And Receivers

```go
type Meter struct {
	Reading int
}

func (meter Meter) Current() int {
	return meter.Reading
}

func (meter *Meter) Increase(amount int) {
	meter.Reading += amount
}
```

Use a value receiver for read-only behavior when copying the value is fine. Use
a pointer receiver when the method changes the original value or the struct is
large. The receiver is the value before the method dot: `meter.Increase(2)`.

## Functions As Values

```go
func apply(value int, transform func(int) int) int {
	return transform(value)
}

result := apply(5, func(value int) int {
	return value * value
})
```

Start with named functions. Use anonymous functions or higher-order functions
only when the repeated behavior is genuinely useful.

## Errors

```go
func parsePositive(value int) (int, error) {
	if value <= 0 {
		return 0, errors.New("value must be positive")
	}
	return value, nil
}
```

Check `err` immediately. Do not return a plausible value together with an
ignored failure.

## Best Practices

- Name functions with a clear action: `findUser`, `calculateTotal`.
- Keep a function small enough that its inputs, state changes, and result are
  easy to trace.
- Prefer returning updated state to hiding mutation.
- Choose methods when behavior belongs to a type; choose functions when it does
  not need a receiver.
- Keep receiver style consistent for one type.

## Official Resources

- [Tour: Functions](https://go.dev/tour/basics/4)
- [Tour: Methods](https://go.dev/tour/methods/1)
- [Tour: Function values](https://go.dev/tour/moretypes/24)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Effective Go: Methods](https://go.dev/doc/effective_go#methods)
