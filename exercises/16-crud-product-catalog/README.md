# Exercise 16 - CRUD Product Catalog

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Build another CRUD program using a slice, but this time make the operation
requirements more explicit than the contact book.

The point is repetition: same CRUD pattern, different data, clearer Go
responsibilities.

## Data Model

Create a `Product` type with:

| Field | Meaning |
|---|---|
| ID | stable product identifier |
| Name | product name |
| PriceCents | price stored as integer cents |
| InStock | whether the product is available |

Use integer cents instead of floats. For example, 1299 means 12.99.

## Storage

Use:

```go
products := []Product{}
```

Do not use a map yet. This exercise is for repeating slice CRUD until the loop
shape is normal.

## Required Functions

Create plain functions, not methods.

| Function | Job |
|---|---|
| `addProduct` | receive current slice and one product, return updated slice |
| `findProductByID` | receive slice and ID, return product plus found boolean |
| `updateProductPrice` | receive slice, ID, new price, return updated slice plus success boolean |
| `updateProductStock` | receive slice, ID, new stock value, return updated slice plus success boolean |
| `deleteProduct` | receive slice and ID, return updated slice plus success boolean |
| `printProducts` | print all products clearly |

## How The Go Flow Works

### Create

Create uses `append`.

The important Go rule: `append` returns the updated slice. The caller must store
that returned slice.

Shape:

```go
products = addProduct(products, product)
```

Inside the function, the result of `append` must be returned.

### Read

Read uses a loop.

The find function should check products one by one:

```text
for each product:
    if product ID matches wanted ID:
        return product and true
after loop:
    return empty product and false
```

The `false` return belongs after the loop. If you return `false` inside the
first failed iteration, you only checked one product.

### Update

Update must modify the actual slice element.

Use index-based looping:

```text
for each index in products:
    if products[index] has wanted ID:
        change the field on products[index]
        return updated products and true
after loop:
    return original products and false
```

Do not update the range value copy.

### Delete

Use the beginner-friendly delete style:

```text
make an empty result slice
for each product:
    if product ID is the wanted ID:
        mark found and skip it
    otherwise:
        append product to result
if not found:
    return original products and false
return result and true
```

This makes it easy to see exactly which products remain.

## Main Program Requirements

In `main`:

1. Start with an empty product slice.
2. Add at least four products.
3. Print the products.
4. Find one existing product.
5. Try to find one missing product.
6. Update the price of one existing product.
7. Try to update the price of a missing product.
8. Update stock for one existing product.
9. Delete one existing product.
10. Try to delete one missing product.
11. Print the final products.

## Proof Cases

| Case | Expected |
|---|---|
| add product | slice length increases |
| find existing ID | product returned, found true |
| find missing ID | empty product, found false |
| update price existing ID | only price changes |
| update stock existing ID | only stock changes |
| update missing ID | no product changes |
| delete existing ID | exactly one product removed |
| delete missing ID | slice unchanged |

## Common Mistakes To Catch

- returning `false` before the loop finishes
- updating the range copy instead of `products[index]`
- forgetting to assign the returned slice in `main`
- accidentally changing ID during update
- deleting by position instead of ID

## Done When

You can implement the same function list for another type without reading this
README.
