# Exercise 07 - Return Updated State

## Goal

Practice functions that receive current state and return updated state.

This is useful before pointer receivers and methods. The caller can see exactly
which value changed because it assigns the returned value.

## Required Data

Create a `CartItem` type with:

- name
- quantity

## Required Functions

| Function | Responsibility |
|---|---|
| `addItem` | add one item to the cart slice |
| `increaseQuantity` | increase quantity for one matching item |
| `removeItem` | remove one matching item |
| `cartTotalQuantity` | return total quantity across all items |

## Main Requirements

In `main`:

1. Start with an empty cart.
2. Add at least three items.
3. Increase the quantity of one existing item.
4. Try to increase a missing item.
5. Remove one existing item.
6. Try to remove a missing item.
7. Print the final cart and total quantity.

## Constraints

- Use functions, not methods.
- Store returned updated state in `main`.
- Return a boolean for operations that may not find an item.
- Do not use a map.

## Prove It Works

Show from output that missing-item operations do not change the cart.
