# Go Quick Reference

Use this page when an exercise mentions a Go feature you have not used yet.
It is a syntax reference, not a solution to any exercise.

## Run And Check Code

```powershell
gofmt -w .\main.go
go run .
go test ./...
go vet ./...
```

Go starts a runnable program at `main` in `package main`. Every imported
package must be used, and every variable must be used.

## Values, Functions, And Results

```go
name := "Mina"
count := 3

func double(value int) int {
	return value * 2
}

result := double(count)
```

Use `:=` for a new local variable. Use `var` when you need an explicit type or
zero value. A function declares parameters and return values after its name.

Multiple returns are common for a result and a failure signal:

```go
func divide(total, parts int) (int, error) {
	if parts == 0 {
		return 0, errors.New("parts cannot be zero")
	}
	return total / parts, nil
}

value, err := divide(10, 2)
if err != nil {
	// Handle the failure before using value.
}
```

## Control Flow

```go
if count > 0 {
	fmt.Println("has values")
} else {
	fmt.Println("empty")
}

for index := 0; index < count; index++ {
	fmt.Println(index)
}

for _, value := range values {
	fmt.Println(value)
}
```

`range` gives copies of slice elements. Use an index when you need to change
the element in the original slice:

```go
for index := range values {
	values[index]++
}
```

## Structs And Methods

```go
type Counter struct {
	Value int
}

func (counter *Counter) Add(amount int) {
	counter.Value += amount
}

counter := Counter{}
counter.Add(2)
```

A method has a receiver. Use a pointer receiver when the method must mutate
the original struct. A value receiver reads a copy. Keep a receiver type's
methods consistent unless there is a deliberate reason not to.

## Slices And Maps

```go
names := []string{"Ana", "Bo"}
names = append(names, "Cy")

scores := map[string]int{"Ana": 10}
score, found := scores["Ana"]
if !found {
	// The key is absent.
}
scores["Bo"] = 8
delete(scores, "Ana")
```

`append` returns the slice value to keep. A map lookup returns a value and an
optional boolean. Map values are copied out, so update a struct value by
reading it, changing it, and assigning it back.

## Errors And Boundaries

Return errors from operations that can fail. Check them immediately at the
boundary where they occur. Add context with `fmt.Errorf("read note: %w", err)`
when the caller needs to know which operation failed.

Do not use `panic` for normal input errors. Validate external input before it
reaches business logic. Keep printing and HTTP response formatting in the
caller; lower-level functions should return values and errors.

## JSON And HTTP Shape

```go
type Message struct {
	Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var input Message
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(input)
}
```

The handler owns HTTP parsing and response formatting. It should not contain
storage or complex business rules. Return after writing an error response.

## Tests

```go
func TestDouble(t *testing.T) {
	got := double(4)
	want := 8
	if got != want {
		t.Fatalf("double(4) = %d, want %d", got, want)
	}
}
```

Test behavior with normal, boundary, invalid, and missing-input cases. A test
name should make the failed behavior obvious.

## Short Best-Practice List

- Format with `gofmt` before reviewing code.
- Give functions one clear responsibility.
- Return errors instead of hiding failures.
- Do not mutate a value accidentally through a range copy.
- Keep validation, business rules, storage, and transport in separate owners as
  the program grows.
- Use `context.Context` for cancellation and deadlines on external work.

## Official Resources

- [Tour of Go](https://go.dev/tour/)
- [Go specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/) for small runnable examples
- [Go standard library documentation](https://pkg.go.dev/std)
- [`testing`](https://pkg.go.dev/testing), [`net/http`](https://pkg.go.dev/net/http), and [`encoding/json`](https://pkg.go.dev/encoding/json)
