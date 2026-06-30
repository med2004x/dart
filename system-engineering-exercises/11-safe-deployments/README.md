# Project 11 - Safe Deployments And Migrations

## Goal

Add task due dates while old and new application versions overlap and remain
rollback-compatible.

## Artifacts

```text
migrations/
    001_add_due_at.up.sql
    001_add_due_at.down.sql
compatibility-matrix.md
deployment-plan.md
```

## Checkpoints

1. Add nullable `due_at`.
2. Run old application queries.
3. Run new application queries.
4. Write new rows with and without due dates.
5. Roll application code back.
6. Backfill existing rows in bounded batches.
7. Add stronger constraints only after old code is gone.

Apply through the PostgreSQL course container:

```powershell
Get-Content -Raw .\migrations\001_add_due_at.up.sql |
    docker exec -i postgres-course `
    psql -X -v ON_ERROR_STOP=1 --single-transaction `
    -U student -d go_course
```

## Failure Drills

1. Rename or drop a column used by old code.
2. Hold a conflicting transaction and apply with a two-second lock timeout.
3. Let one migration statement fail without `ON_ERROR_STOP`.
4. Roll back application code after writing new-format data.

## Done Means

Forward schema deployment and application rollback are both proven.

