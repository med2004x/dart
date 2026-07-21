# Exercise 09 - Payment Interfaces

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

An interface is a set of required methods. It lets a caller depend on behavior
without depending on one concrete type.

Use an interface only after at least two concrete types need to be used through
the same operation. Starting with an interface before concrete behavior exists
usually creates abstraction without evidence.

## Beginner Bridge: Interfaces Describe Needed Behavior

An interface does not say what a value is. It says what a value can do.

For checkout, the code does not need to know whether payment comes from a card
or a wallet. It needs one behavior:

```text
can this thing pay this amount?
```

That becomes:

```go
type paymentMethod interface {
    pay(amount int) bool
}
```

Similar programs:

| Caller | Interface behavior | Concrete examples |
|---|---|---|
| Checkout | `pay(amount)` | card, wallet |
| Notification service | `send(message)` | email, console |
| Report exporter | `export(data)` | CSV, JSON |

If there is only one concrete type, an interface may be premature. If the caller
needs ten methods, the interface is probably too large. The right interface is
the smallest behavior the caller actually uses.

## Concrete Types First

Suppose two payment types store different data:

```text
Card payment   -> card number, available balance
Wallet payment -> username, available balance
```

They can still share one behavior:

```go
pay(amount int) bool
```

The interface describes that common requirement:

```go
type paymentMethod interface {
    pay(amount int) bool
}
```

Read it as: "Any value with a method named `pay` accepting an `int` and
returning a `bool` may be used as a `paymentMethod`."

Go has no `implements` declaration. Satisfaction is implicit and checked by the
compiler.

## Method Sets Must Match Exactly

These do not satisfy the interface:

```go
Pay(amount int) bool        // uppercase name is different
pay(amount float64) bool    // parameter type is different
pay(amount int) error       // return type is different
```

Method names, parameters, and results must match.

## The Caller Owns The Small Interface

`checkout` only needs one operation:

```go
func checkout(method paymentMethod, amount int) bool {
    return method.pay(amount)
}
```

It should not know card fields, wallet fields, or use a string type switch:

```go
if paymentType == "card" {
    // ...
}
```

The concrete type owns how payment works. The consumer states the minimum
behavior it needs.

## State Changes Need Pointer Receivers

A payment generally reduces the original balance. A value receiver would change
only a copy:

```go
func (card CardPayment) pay(amount int) bool
```

A pointer receiver can mutate the original:

```go
func (card *CardPayment) pay(amount int) bool
```

Then `*CardPayment` satisfies the interface. A plain `CardPayment` value does
not have the pointer-only method in its method set.

Pseudocode:

```text
METHOD pay(amount)
    IF amount is not positive
        RETURN false
    IF amount exceeds balance
        RETURN false
    subtract amount from balance
    RETURN true
```

## Worked Example: Notification Senders

```go
package main

import "fmt"

type notifier interface {
    send(message string) error
}

type EmailNotifier struct {
    Address string
}

func (email EmailNotifier) send(message string) error {
    fmt.Println("email to", email.Address+":", message)
    return nil
}

type ConsoleNotifier struct{}

func (ConsoleNotifier) send(message string) error {
    fmt.Println("console:", message)
    return nil
}

func announce(destination notifier, message string) error {
    return destination.send(message)
}

func main() {
    email := EmailNotifier{Address: "student@example.com"}
    console := ConsoleNotifier{}

    if err := announce(email, "build completed"); err != nil {
        fmt.Println("email failed:", err)
    }
    if err := announce(console, "build completed"); err != nil {
        fmt.Println("console failed:", err)
    }
}
```

`announce` is closed to notifier-specific details but accepts both concrete
types. A test could supply a fake notifier implementing the same one method.

## Interface Value Mental Model

An interface value conceptually contains:

```text
dynamic concrete type + dynamic concrete value
```

For `announce(email, ...)`:

```text
interface type: notifier
dynamic type:   EmailNotifier
dynamic value:  EmailNotifier{Address: ...}
```

The method call dispatches to `EmailNotifier.send`.

Avoid storing a nil pointer inside a non-nil interface until you understand the
"typed nil" problem. It can make `value != nil` while the contained pointer is
nil.

## Your Program

Build:

1. A payment interface with `pay(amount int) bool`.
2. Card and wallet concrete types.
3. A pay method on each type.
4. A checkout function that receives only the interface.
5. Calls using both concrete payment types.

Decide whether successful payment reduces balance. If it does, use pointer
receivers and verify the balance after checkout.

## Build In Checkpoints

1. Implement card payment without an interface.
2. Implement wallet payment without an interface.
3. Confirm both methods have identical signatures.
4. Define the minimal interface.
5. Write checkout against the interface.
6. Test invalid, insufficient, and valid amounts for both types.

```powershell
gofmt -w .\main.go
go run .
go vet .
```

## Test Table

Assume balance 100:

| Amount | Expected result | Expected balance |
|---:|---|---:|
| -1 | false | 100 |
| 0 | false | 100 |
| 60 | true | 40 |
| 100 | true | 0 |
| 101 | false | 100 |

Use a fresh payment value for each case.

## Failure Experiments

1. Change one method parameter to `float64` and pass it to checkout. Read the
   compiler's missing-method explanation.
2. Use a value receiver, subtract balance, then print the original balance.
3. Add card-number logic inside checkout. Explain how that defeats the interface.
4. Create a ten-method payment interface when checkout needs one. Explain the
   unnecessary burden on every implementation.

## You Understand This Exercise When

You can state why the interface exists, identify which concrete values satisfy
it, explain pointer method sets, and build the same design for repositories,
loggers, or notification senders.

References:

- [Interfaces](https://go.dev/tour/methods/9)
- [Interface values](https://go.dev/tour/methods/11)
- [Methods and pointer indirection](https://go.dev/tour/methods/6)
