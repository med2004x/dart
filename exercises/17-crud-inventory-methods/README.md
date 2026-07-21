# Exercise 17 - CRUD Inventory With Methods

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Build CRUD again, but this time wrap the slice inside a struct and use methods.

Exercise 16 used plain functions. This exercise shows when methods make sense:
the operations belong to one owner type, `Inventory`.

## Data Model

Create an `Item` type with:

| Field | Meaning |
|---|---|
| ID | stable item identifier |
| Name | item name |
| Quantity | amount available |

Create an `Inventory` type that owns the slice:

```text
Inventory contains []Item
```

You choose the exact field names.

## Required Methods

Create methods on `Inventory`.

| Method | Receiver kind | Job |
|---|---|---|
| `AddItem` | pointer receiver | add one item |
| `FindItemByID` | value or pointer receiver | return item plus found boolean |
| `UpdateQuantity` | pointer receiver | change quantity for one item |
| `DeleteItem` | pointer receiver | remove one item |
| `PrintItems` | value or pointer receiver | print all items |

## Why Pointer Receivers Here

Methods that change the inventory should use a pointer receiver because the
method must change the original inventory, not a copy.

Mutating methods:

- add
- update
- delete

Read-only methods:

- find
- print

Read-only methods can use value receivers, but pointer receivers are also fine
if you keep the style consistent.

## How The Go Flow Works

### Method Calls

Plain function shape:

```text
updated = addItem(updated, item)
```

Method shape:

```text
inventory.AddItem(item)
```

The receiver is the value before the dot. In this exercise, the receiver is the
inventory.

### Mutating State

If `Inventory` contains a slice, adding an item changes the inventory's slice.
That is why `AddItem`, `UpdateQuantity`, and `DeleteItem` need access to the
original inventory value.

### Return Booleans

Update and delete should still report whether they found the item.

A method can mutate state and still return a boolean:

```text
success = inventory.UpdateQuantity(id, quantity)
```

## Main Program Requirements

In `main`:

1. Create an empty inventory.
2. Add at least four items.
3. Print all items.
4. Find one existing item.
5. Try to find one missing item.
6. Update quantity for one existing item.
7. Try to update quantity for one missing item.
8. Delete one existing item.
9. Try to delete one missing item.
10. Print final inventory.

## Proof Cases

| Case | Expected |
|---|---|
| add item | inventory grows |
| find existing ID | found true |
| find missing ID | found false |
| update existing ID | quantity changes |
| update missing ID | inventory unchanged |
| delete existing ID | item removed |
| delete missing ID | inventory unchanged |

## Common Mistakes To Catch

- using a value receiver for a method that must replace the slice
- putting all logic directly in `main`
- not returning a boolean from update/delete
- changing every item instead of only the matching item

## Done When

You can explain which methods mutate state and why their receiver choice matters.
