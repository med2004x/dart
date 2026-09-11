# Project 02 - Schema, Tables, And Types

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Model users, projects, and tasks with types matching their real meaning.

## Required Shape

Users:

```text
id, email, display_name, created_at
```

Projects:

```text
id, owner_id, name, created_at
```

Tasks:

```text
id, project_id, title, status, priority, due_at, created_at, updated_at
```

Use:

- identity `bigint` primary keys
- `text` for names and titles
- `smallint` for priority
- `timestamptz` for instants
- nullable `due_at`
- server defaults for timestamps, status, and priority

Constraints beyond primary key and not-null come in project 05.

## Checkpoints

1. Create schema `app`.
2. create tables in dependency order.
3. inspect every table with `\d`.
4. run `verify.sql`.
5. explain every type, nullability choice, and default.

```bash
.\scripts\run-sql.ps1 .\02-schema-types\schema.sql
.\scripts\run-sql.ps1 .\02-schema-types\verify.sql
```

## Failure Drills

1. Run `schema.sql` twice.
2. insert text into priority.
3. insert null into a required column.
4. store a date as free-form text and explain what validation is lost.

## Done Means

The verification query returns all expected columns with intentional types and
nullability.

