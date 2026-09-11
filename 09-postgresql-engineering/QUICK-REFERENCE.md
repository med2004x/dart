# PostgreSQL Quick Reference

Use this page for the recurring SQL syntax. Run every statement against the
course database and inspect the result before changing more data.

## Connect And Inspect

```sql
\dt
\d projects
SELECT current_database();
```

`\dt` lists tables and `\d` describes a table in `psql`.

## CRUD Shape

```sql
INSERT INTO projects (name) VALUES ('Roadmap') RETURNING id, name;

SELECT id, name FROM projects WHERE id = 1;

UPDATE projects
SET name = 'Updated roadmap'
WHERE id = 1
RETURNING id, name;

DELETE FROM projects
WHERE id = 1
RETURNING id;
```

Always write the `WHERE` clause before executing `UPDATE` or `DELETE`. Use
`RETURNING` while learning so the database shows what changed.

## Tables And Constraints

```sql
CREATE TABLE labels (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT now()
);
```

Use database constraints for invariants that must hold regardless of which
program writes the data. Typical tools are `NOT NULL`, `UNIQUE`, `CHECK`,
`PRIMARY KEY`, and `FOREIGN KEY`.

## Joins And Aggregation

```sql
SELECT projects.name, tasks.title
FROM projects
JOIN tasks ON tasks.project_id = projects.id
WHERE tasks.completed = false
ORDER BY projects.name, tasks.title;
```

```sql
SELECT project_id, count(*) AS task_count
FROM tasks
GROUP BY project_id
HAVING count(*) > 2;
```

The `ON` clause explains how rows relate. `WHERE` filters rows before grouping;
`HAVING` filters groups after aggregation.

## Transactions

```sql
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id = 1;
UPDATE accounts SET balance = balance + 10 WHERE id = 2;

-- Inspect affected rows, then choose one:
COMMIT;
-- ROLLBACK;
```

Use a transaction when several statements must succeed or fail together. Keep
transactions short and never leave one open while waiting for user input.

## Parameters And Go

Do not build SQL by concatenating user input. Use parameters through
`database/sql`:

```go
row := db.QueryRowContext(ctx,
	"SELECT id, name FROM projects WHERE id = $1", projectID)
```

The placeholder is data, not SQL text. The driver sends the value separately.

## Measuring Queries

```sql
EXPLAIN (ANALYZE, BUFFERS)
SELECT id FROM tasks WHERE project_id = 1;
```

Measure before adding an index. Compare the plan, rows examined, and execution
time after the change.

## Best Practices

- Treat `UPDATE` and `DELETE` without a verified `WHERE` as dangerous.
- Use constraints for data integrity, not only application checks.
- Use explicit columns instead of `SELECT *` in application queries.
- Use transactions for multi-step writes.
- Parameterize all external values.
- Back up and restore as separate tested operations.

## Official Resources

- [PostgreSQL SQL language](https://www.postgresql.org/docs/current/sql.html)
- [SQL syntax](https://www.postgresql.org/docs/current/sql-syntax.html)
- [Transactions](https://www.postgresql.org/docs/current/tutorial-transactions.html)
- [Indexes and `EXPLAIN`](https://www.postgresql.org/docs/current/using-explain.html)
- [`database/sql` package](https://pkg.go.dev/database/sql)
