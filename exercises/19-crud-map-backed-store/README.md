# Exercise 19 - CRUD Map-Backed Store

## Goal

Build CRUD using a map instead of a slice.

The earlier CRUD exercises made you scan slices. A map changes the lookup model:
you can go directly from ID to record.

## Data Model

Create an `Employee` type with:

| Field | Meaning |
|---|---|
| ID | stable employee identifier |
| Name | employee name |
| Department | department name |
| Active | whether employee is active |

## Storage

Use:

```text
map[int]Employee
```

The key is the employee ID. The value is the employee record.

## Required Functions

Use plain functions.

| Function | Job |
|---|---|
| `addEmployee` | add employee if ID does not already exist |
| `findEmployee` | look up employee by ID |
| `updateDepartment` | change department for one employee |
| `deactivateEmployee` | set active to false |
| `deleteEmployee` | remove employee from map |
| `printEmployees` | print all employees |

## How Map CRUD Works In Go

### Create

Create is assignment by key:

```text
employees[id] = employee
```

But first check whether the key already exists. Otherwise you may overwrite an
existing employee.

### Read

Map lookup can return two values:

```text
value, found = employees[id]
```

The boolean tells whether the key exists.

### Update

Map values that are structs need care.

Do not assume changing a copy updates the map. A safe beginner flow is:

```text
employee, found = employees[id]
if not found: report failure
change employee copy
assign employee back into employees[id]
```

The final assignment back to the map is required.

### Delete

Delete by key:

```text
delete(employees, id)
```

Check existence first if the caller needs to know whether anything was removed.

## Main Program Requirements

In `main`:

1. Create an empty employee map.
2. Add at least four employees.
3. Try to add a duplicate ID.
4. Find an existing employee.
5. Try to find a missing employee.
6. Update one existing department.
7. Try to update a missing employee.
8. Deactivate one existing employee.
9. Delete one existing employee.
10. Try to delete a missing employee.
11. Print final employees.

## Proof Cases

| Case | Expected |
|---|---|
| add new ID | employee appears |
| add duplicate ID | existing employee is not overwritten |
| find existing ID | found true |
| find missing ID | found false |
| update department | changed value is stored back |
| deactivate employee | active becomes false |
| delete existing ID | key removed |
| delete missing ID | no crash, failure reported |

## Common Mistakes To Catch

- overwriting duplicate IDs
- changing a map value copy and forgetting to assign it back
- assuming map print order is stable
- deleting without reporting whether the key existed

## Done When

You can explain the difference between slice search and map lookup.
