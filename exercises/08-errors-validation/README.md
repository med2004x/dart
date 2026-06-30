# Exercise 08 - Errors And Validation

## What You Are Learning

An error is a value describing why a requested operation could not be completed.
The function that detects the problem returns it. The caller decides how to
present or transport it.

This separation matters:

```text
business function -> reports rule violation
terminal program   -> prints a message
HTTP handler       -> sends a status code and JSON
background job     -> logs and retries or stops
```

The same business function can be reused in all three environments.

## Beginner Bridge: Errors Are Part Of The Function Result

An error is not a crash by default. In Go, an error is a normal returned value
that says: "the requested operation did not complete, and here is why."

Success:

```text
updated users, nil error
```

Failure:

```text
original users, non-nil error
```

That contract matters because the caller needs to know whether it is safe to use
the returned data.

Similar validation programs:

| Program | Invalid case | Error meaning |
|---|---|---|
| Product catalog | blank SKU | product cannot be saved |
| Signup form | weak password | user cannot be registered |
| Expense tracker | negative amount | expense is invalid |

Do not print inside validation. Validation detects facts. The caller decides
whether to print, return JSON, log, retry, or stop.

## Function Contracts Include Failure

This signature:

```go
func registerUser(users []User, candidate User) ([]User, error)
```

says:

- input: current users and a candidate
- success: updated users and `nil`
- failure: an error and a slice whose meaning must be documented

For this exercise, failure should return the original users unchanged.

Typical call:

```go
updatedUsers, err := registerUser(users, candidate)
if err != nil {
    fmt.Println("registration failed:", err)
    return
}
users = updatedUsers
```

Never assign the result before deciding what failure returns.

## Validation And Business Rules

Different checks answer different questions:

- **shape validation:** Is username empty? Is password long enough?
- **business rule:** Is this username already registered?

Keep a single-user validator focused:

```text
FUNCTION validate user(candidate)
    IF trimmed username is empty
        RETURN username-required error
    IF password has fewer than 8 characters
        RETURN password-too-short error
    RETURN no error
```

Registration coordinates validation and collection rules:

```text
FUNCTION register user(users, candidate)
    validate candidate
    IF validation failed
        RETURN original users and error

    FOR each existing user
        IF usernames match
            RETURN original users and duplicate error

    append candidate
    RETURN updated users and no error
```

Validate before mutating state. A failed operation must not leave a partial
change.

## Creating And Wrapping Errors

Simple error:

```go
errors.New("username is required")
```

Error with context:

```go
fmt.Errorf("register %q: %w", candidate.Username, err)
```

`%w` wraps the original error. It retains machine-checkable identity while
adding information for humans.

Do not inspect errors by comparing their message strings. Messages are for
people and may change.

## Errors Are Not Logging

Wrong responsibility:

```go
func validateUser(user User) error {
    if user.Username == "" {
        fmt.Println("bad username")
        return errors.New("bad username")
    }
    return nil
}
```

Now every caller gets unwanted terminal output, and a server may log the same
failure twice.

The validator should only return the error.

## Worked Example: Add A Product

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

type Product struct {
    SKU   string
    Name  string
    Price int
}

func validateProduct(product Product) error {
    if strings.TrimSpace(product.SKU) == "" {
        return errors.New("SKU is required")
    }
    if strings.TrimSpace(product.Name) == "" {
        return errors.New("name is required")
    }
    if product.Price <= 0 {
        return errors.New("price must be positive")
    }
    return nil
}

func addProduct(products []Product, candidate Product) ([]Product, error) {
    if err := validateProduct(candidate); err != nil {
        return products, fmt.Errorf("validate product: %w", err)
    }

    for _, product := range products {
        if product.SKU == candidate.SKU {
            return products, errors.New("SKU already exists")
        }
    }

    return append(products, candidate), nil
}
```

This is the same design as registration: validate one value, check collection
rules, mutate only after every check passes.

## Your Program

Build:

1. A user type with ID, username, and password.
2. `validateUser` that rejects blank usernames and passwords shorter than eight
   characters.
3. `registerUser` that rejects duplicate usernames.
4. A `main` function demonstrating valid and invalid attempts.

Passwords remain plain text only because this is an in-memory validation
exercise. Never store them this way in a deployed application.

## Build In Checkpoints

1. Validate one valid user.
2. Test blank and whitespace-only usernames.
3. Test password lengths 7 and 8.
4. Call validation from registration.
5. Add duplicate detection.
6. Prove the input slice stays unchanged after every failure.

```powershell
gofmt -w .\main.go
go run .
go vet .
```

## Test Table

| Candidate | Existing users | Expected |
|---|---|---|
| valid | empty | appended, nil error |
| blank username | any | unchanged, required error |
| whitespace username | any | unchanged, required error |
| 7-character password | any | unchanged, length error |
| 8-character password | empty | success |
| duplicate username | matching user | unchanged, duplicate error |

Define whether usernames are case-sensitive. If `"Sara"` and `"sara"` should be
the same, normalize before duplicate comparison.

## Failure Experiments

1. Append before duplicate checking. Show how failure leaves bad state.
2. Print inside validation, then call it from two different callers. Observe the
   unwanted coupling.
3. Ignore the returned error and print success. Explain why ignoring errors makes
   program output untrustworthy.
4. Check only `username == ""`, then test spaces.

## You Understand This Exercise When

You can identify who detects, reports, and presents each failure; prove failed
operations do not mutate state; and reuse the same validator outside a terminal
program.

References:

- [`errors`](https://pkg.go.dev/errors)
- [Error handling in Go](https://go.dev/blog/error-handling-and-go)
- [`fmt.Errorf`](https://pkg.go.dev/fmt#Errorf)
