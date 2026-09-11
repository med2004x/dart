# Exercise 01 - Bank Account

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise teaches four foundations:

- a **type** describes what kind of value something is
- a **struct** groups values that belong to one concept
- a **method** gives a type behavior
- a **receiver** is the value on which a method operates

The real lesson is state: what changes, what does not change, and why.

## Beginner Bridge: From Variables To Objects With Behavior

At the start, a program is just named values:

```go
owner := "Sara"
balance := 70
```

That works until the values belong together. A bank account is not two random
variables. It is one thing with two pieces of information. That is what a struct
is for: it gives related data one name.

```go
type Account struct {
    Owner   string
    Balance int
}
```

Methods answer the next question: "What can this thing do?" For an account,
useful actions are deposit, withdraw, and print a summary.

Think of the exercise like this:

```text
data:      Account{Owner, Balance}
rules:     reject invalid deposits and withdrawals
behavior:  methods that apply those rules
proof:     print the balance before and after each operation
```

Similar programs use the same shape:

| Program | Struct | State field | Behavior |
|---|---|---|---|
| Inventory item | `Item` | `Stock` | add stock, remove stock |
| Game player | `Player` | `Health` | heal, take damage |
| Thermostat | `Thermostat` | `Temperature` | warm up, cool down |

The important question is not "how do I write a method?" It is "which state is
allowed to change, and under which rule?"

## Start With The Data

A bank account has related values:

```text
owner: "Sara"
balance: 70
```

Separate variables work for one account:

```go
owner := "Sara"
balance := 70
```

They become fragile with many accounts because nothing tells Go that each owner
belongs to a specific balance. A struct creates that relationship:

```go
type Account struct {
    Owner   string
    Balance int
}
```

Read it aloud: "`Account` is a new type containing an `Owner` string and a
`Balance` integer."

Creating a value:

```go
account := Account{
    Owner:   "Sara",
    Balance: 70,
}
```

After this statement, memory conceptually contains:

```text
account
|-- Owner   = "Sara"
`-- Balance = 70
```

## Methods And Receivers

A normal function receives arguments:

```go
func printName(name string) {
    fmt.Println(name)
}
```

A method adds a receiver before its name:

```go
func (account Account) printSummary() {
    fmt.Println(account.Owner, account.Balance)
}
```

Call:

```go
account.printSummary()
```

Execution:

1. Go evaluates the value on the left side of the dot.
2. That value becomes the receiver named `account`.
3. The method body reads fields from that receiver.

The receiver name is local to the method. It could be `a`, but `account` is
clearer for a beginner.

## Value Receiver: Why Assignment Is Needed

This method uses a value receiver:

```go
func (account Account) increasedBalance(amount int) int {
    return account.Balance + amount
}
```

`account` inside the method is a copy. The method returns a number; it does not
replace the original field automatically.

```go
account.increasedBalance(40)                   // result is discarded
account.Balance = account.increasedBalance(40) // result is stored
```

Trace the second line:

| Step | Value |
|---|---:|
| original `account.Balance` | 70 |
| argument `amount` | 40 |
| returned expression | 110 |
| assigned `account.Balance` | 110 |

Later, a pointer receiver such as `func (account *Account) deposit(...)` can
change the original struct. This exercise can use the return-and-assign pattern
so the copy behavior remains visible.

## Design The Rules Before The Syntax

A valid deposit must be positive. A valid withdrawal must be positive and no
larger than the current balance.

Pseudocode:

```text
FUNCTION deposit(account, amount)
    IF amount is zero or negative
        RETURN the unchanged balance
    RETURN balance plus amount

FUNCTION withdraw(account, amount)
    IF amount is zero or negative
        RETURN the unchanged balance
    IF amount is greater than balance
        RETURN the unchanged balance
    RETURN balance minus amount
```

Notice that every path returns a value. Ask what should happen before writing an
`if`; do not add branches while guessing.

## Worked Example: Thermostat

This example uses the same idea in a different program:

```go
package main

import "fmt"

type Thermostat struct {
    Room        string
    Temperature int
}

func (thermostat Thermostat) warmerBy(degrees int) int {
    if degrees <= 0 {
        return thermostat.Temperature
    }
    return thermostat.Temperature + degrees
}

func (thermostat Thermostat) printStatus() {
    fmt.Println("Room:", thermostat.Room)
    fmt.Println("Temperature:", thermostat.Temperature)
}

func main() {
    thermostat := Thermostat{Room: "office", Temperature: 19}
    thermostat.Temperature = thermostat.warmerBy(2)
    thermostat.printStatus()
}
```

Expected output:

```text
Room: office
Temperature: 21
```

Transfer the pattern, not the names:

```text
Thermostat     -> Account
Temperature    -> Balance
warmerBy       -> deposit
invalid degree -> invalid amount
```

## Your Program

Build:

1. `Account` with owner and balance fields.
2. A deposit method that rejects zero and negative amounts.
3. A withdrawal method that also rejects insufficient funds.
4. A summary method.
5. A `main` function that creates an account, deposits, withdraws, and prints.

Use integers for this exercise. In real financial software the integer would
normally represent the smallest currency unit, such as cents.

## Build In Checkpoints

1. Define the struct and create one value.
2. Print both fields directly.
3. Add the summary method and call it.
4. Add deposit logic and test one valid amount.
5. Test zero and a negative deposit.
6. Add withdrawal logic.
7. Test valid, excessive, zero, and negative withdrawals.

Run after every checkpoint:

```bash
gofmt -w .\main.go
go run .
```

## Test Table

Assume the opening balance is 80:

| Operation | Amount | Expected balance | Reason |
|---|---:|---:|---|
| deposit | 40 | 120 | valid positive amount |
| deposit | 0 | 80 | zero is rejected |
| deposit | -10 | 80 | negative is rejected |
| withdraw | 20 | 60 | enough funds |
| withdraw | 100 | 80 | insufficient funds |
| withdraw | 0 | 80 | zero is rejected |

Test each case from a fresh account so one test does not change the starting
state of the next test.

## Failure Experiments

Try these deliberately:

1. Remove the assignment from the deposit call. Explain why the balance stays
   unchanged.
2. Change `Balance` to `string`. Read the compiler error produced by arithmetic.
3. Remove one return path. Read how Go proves that a result might be missing.

Then restore the correct code.

## You Understand This Exercise When

You can explain:

- why the struct is better than unrelated variables
- what value becomes the receiver
- why a value receiver works on a copy
- why returning a balance is not the same as storing it
- which inputs leave the account unchanged
- how you would adapt the same pattern to inventory stock

References:

- [Structs](https://go.dev/tour/moretypes/2)
- [Methods](https://go.dev/tour/methods/1)
- [Pointer receivers](https://go.dev/tour/methods/4)
