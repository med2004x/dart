# Exercise 02 - Auth Users

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise teaches how to search a collection:

- a **slice** stores an ordered sequence of values of one type
- `range` visits each value
- a condition decides whether the current value matches
- multiple return values report both the result and whether it exists

The login scenario is only a learning model. Plain-text passwords are never
acceptable authentication storage in a real system.

## Beginner Bridge: What Search Really Means

Search means: look through a collection one item at a time until you find the
item that matches a rule.

For this exercise, the collection is:

```go
users := []User{...}
```

The rule is:

```text
username matches AND password matches
```

The algorithm is not magic:

```text
check first user
if it matches, return it
otherwise check second user
if it matches, return it
continue until there are no users left
then report failure
```

That final "then report failure" is important. You cannot know the user is
missing until every user has been checked.

Similar programs:

| Program | Collection | Match rule |
|---|---|---|
| Product lookup | `[]Product` | product code equals wanted code |
| Order lookup | `[]Order` | order ID equals wanted ID |
| Student lookup | `[]Student` | email equals wanted email |

The pattern is always: loop, compare, return success early, return failure after
the loop.

## From One User To Many

One value:

```go
type User struct {
    Username string
    Password string
    Role     string
}

oneUser := User{Username: "sara", Password: "demo", Role: "admin"}
```

Many values:

```go
users := []User{
    {Username: "sara", Password: "demo", Role: "admin"},
    {Username: "adam", Password: "test", Role: "reader"},
}
```

Read `[]User` as "a slice of `User` values." The slice has length 2. Valid
indexes are 0 and 1 because indexing starts at zero.

## What `range` Actually Produces

```go
for index, currentUser := range users {
    fmt.Println(index, currentUser.Username)
}
```

For the slice above:

| Iteration | `index` | `currentUser.Username` |
|---:|---:|---|
| 1 | 0 | `sara` |
| 2 | 1 | `adam` |

If the index is not needed:

```go
for _, currentUser := range users {
    fmt.Println(currentUser.Username)
}
```

`_` means "this value is intentionally unused." It is not a normal variable.

A common mistake is:

```go
for currentUser := range users {
    fmt.Println(currentUser)
}
```

With one variable, `range` produces indexes, so this prints `0`, then `1`.

## Equality And Boolean Conditions

Authentication succeeds only when both fields match:

```go
currentUser.Username == username &&
    currentUser.Password == password
```

Each `==` expression produces a boolean:

```text
username comparison -> true or false
password comparison -> true or false
&& combines them     -> true only when both are true
```

Strings themselves are not conditions. This is invalid:

```go
if currentUser.Username && currentUser.Password {
}
```

Go requires the condition after `if` to have type `bool`.

## Why Return `(User, bool)`

Search has two outcomes:

```text
found     -> return the matching user and true
not found -> return an empty user and false
```

The zero value of a struct contains the zero value of every field:

```go
User{} // Username "", Password "", Role ""
```

Returning only `User{}` would be ambiguous. The `bool` makes the outcome
explicit.

Call syntax:

```go
matchedUser, found := authenticate(users, "sara", "demo")
if found {
    fmt.Println("role:", matchedUser.Role)
}
```

The left side has two variables because the function returns two values.

## Pseudocode First

```text
FUNCTION authenticate(users, wanted username, wanted password)
    FOR each current user in users
        IF current username equals wanted username
           AND current password equals wanted password
            RETURN current user and true
    RETURN empty user and false
```

The final return must be after the loop. Returning failure inside the loop would
stop after checking only the first user.

## Worked Example: Find A Product

This uses the same search pattern without solving the login task:

```go
package main

import "fmt"

type Product struct {
    Code  string
    Price int
}

func findProduct(products []Product, wantedCode string) (Product, bool) {
    for _, product := range products {
        if product.Code == wantedCode {
            return product, true
        }
    }
    return Product{}, false
}

func main() {
    products := []Product{
        {Code: "PEN", Price: 200},
        {Code: "BOOK", Price: 900},
    }

    product, found := findProduct(products, "BOOK")
    if !found {
        fmt.Println("product not found")
        return
    }

    fmt.Println(product.Code, product.Price)
}
```

Trace for `"BOOK"`:

| Iteration | Current code | Match? | Action |
|---:|---|---|---|
| 1 | `PEN` | false | continue |
| 2 | `BOOK` | true | return product, true |

Transfer:

```text
products     -> users
wantedCode   -> username and password
Product,bool -> User,bool
```

## Your Program

Build:

1. A user type containing username, password, and role.
2. A slice containing at least three users.
3. `authenticate` receiving the slice, username, and password.
4. A `(User, bool)` return.
5. One successful and one failed login in `main`.

Do not use a map or `errors` yet. The goal is to understand linear search and
explicit success reporting.

## Build In Checkpoints

1. Create the type and print one user.
2. Create the slice and print every username.
3. Write a search that checks username only.
4. Return the user and boolean.
5. Add the password condition.
6. Test success, wrong password, unknown user, and empty slice.

```bash
gofmt -w .\main.go
go run .
```

## Test Table

| Username | Password | Expected |
|---|---|---|
| existing | correct | matched user, `true` |
| existing | wrong | empty user, `false` |
| unknown | any | empty user, `false` |
| empty | empty | normally false |
| any | any, empty user slice | false |

## Security Boundary

This exercise demonstrates search logic, not real authentication. Production
authentication requires:

- salted password hashing with Argon2, bcrypt, or scrypt
- constant-time comparison where appropriate
- rate limiting
- secure sessions or signed tokens
- no passwords in logs

Do not copy this model into a deployed service.

## Failure Experiments

1. Put `return User{}, false` inside the loop and search for the second user.
   Explain why the second user is never checked.
2. Use one range variable and print it. Confirm that it is an index.
3. Return a string `"true"` instead of a boolean. Explain why strings prevent
   normal `if found` usage.

## You Understand This Exercise When

You can trace every loop iteration, explain why failure is returned after the
loop, and adapt the pattern to finding an order, task, or product.

References:

- [Slices](https://go.dev/tour/moretypes/7)
- [`range`](https://go.dev/tour/moretypes/16)
- [`if`](https://go.dev/tour/flowcontrol/5)
