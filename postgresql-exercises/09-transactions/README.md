# Project 09 - Transactions

## Goal

Transfer value between accounts so all required writes commit or none do.

## Setup

Run `setup.sql`, then verify total money.

## Checkpoints

1. write a successful transfer in `transfer.sql`.
2. debit only when sufficient balance exists.
3. require exactly one source row.
4. require exactly one destination row.
5. commit only after both updates.
6. verify total money is unchanged.
7. force an error after debit and prove rollback.

## Required Cases

- successful transfer
- insufficient balance
- missing source
- missing destination
- same source/destination policy
- forced second-statement failure

## Failure Drill

Run the debit and credit with autocommit, then force credit failure. Observe the
partial state and restore course data deliberately.

## Done Means

No failed transfer changes one account without the other.

