# Project 13 - Migrations

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Add task descriptions without breaking old application queries or preventing
application rollback.

## Files

```text
migrations/
    001_add_description.up.sql
    001_add_description.down.sql
compatibility-matrix.md
```

## Checkpoints

1. add nullable description.
2. verify old selects/inserts.
3. verify new selects/inserts.
4. backfill in bounded batches.
5. run old code after new data exists.
6. add stronger constraints only when all rows and writers are compatible.
7. record applied migration version.

Run:

```powershell
Get-Content -Raw .\13-migrations\migrations\001_add_description.up.sql |
    docker exec -i postgres-course `
    psql -X -v ON_ERROR_STOP=1 --single-transaction `
    -U student -d go_course
```

## Failure Drills

1. rename/drop a column used by old code.
2. add not-null before backfill.
3. run without error-stop.
4. hold a conflicting lock and use `lock_timeout`.

## Done Means

The expanded schema supports old and new code, and application rollback is
proven after migration.

