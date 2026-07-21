# Project 15 - Backup And Restore

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Create a logical backup and restore it into a separate database. A backup is not
proven until restore and application verification pass.

## Run

```powershell
.\15-backup-restore\backup.ps1
.\15-backup-restore\restore.ps1
.\scripts\run-sql.ps1 .\15-backup-restore\source-verification.sql
```

The provided `run-sql.ps1` targets `go_course`. Run the same verification
directly against `go_course_restore` after restore.

## Checkpoints

1. record source row counts and totals.
2. create a custom-format dump.
3. copy it outside the container.
4. create a separate restore database.
5. restore with `pg_restore`.
6. compare schema, constraints, indexes, counts, and representative queries.
7. measure restore duration.
8. record RPO and RTO.

## Failure Drills

1. truncate a copied dump and restore it.
2. restore with a required role missing.
3. restore into conflicting objects without a policy.
4. verify only row count and miss a broken application query.

## Done Means

A separate restored database serves required queries within the stated recovery
time and data-loss targets.

