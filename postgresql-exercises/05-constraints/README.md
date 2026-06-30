# Project 05 - Constraints

## Goal

Make PostgreSQL reject invalid state regardless of which caller writes it.

## Required Constraints

- unique user email
- nonblank email and display name
- nonblank project name
- nonblank task title
- task status in `open`, `blocked`, `done`
- priority from 1 through 5
- project owner references user
- task project references project
- deleting a user with projects is restricted
- deleting a project cascades to its tasks

## Checkpoints

1. ensure current data satisfies every rule.
2. add named constraints in `constraints.sql`.
3. run every invalid case inside a transaction.
4. record the exact constraint name from each error.
5. verify valid inserts still work.

## Failure Drills

Run cases for duplicate email, blank title, priority 8, unknown status,
nonexistent project, and deleting a referenced owner.

## Design Question

Defend `RESTRICT` versus `CASCADE` for each relationship. Cascade is not a
default; it is a deletion policy.

## Done Means

Direct SQL that bypasses application validation still cannot create invalid
state.

