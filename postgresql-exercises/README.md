# PostgreSQL And SQL Project Track

This folder teaches relational databases through one evolving project/task
database. It is cumulative: each numbered project assumes the previous project
is complete.

## What You Build

```text
users
`-- projects
    `-- tasks
        `-- task_tags -- tags
```

Later projects add accounts, transactions, concurrency tests, indexes,
migrations, restricted roles, backup/restore, and a Go repository.

## Local PostgreSQL

The track uses the official PostgreSQL 18 container image pinned in
`docker-compose.yml`.

Prerequisites:

```powershell
docker version
docker compose version
```

Start:

```powershell
Set-Location C:\Users\pc\Documents\dart\postgresql-exercises
docker compose up -d
docker compose logs postgres
```

Connect:

```powershell
.\scripts\connect.ps1
```

Run an SQL file and stop on the first error:

```powershell
.\scripts\run-sql.ps1 .\02-schema-types\schema.sql
```

Stop without deleting data:

```powershell
docker compose stop
```

The named volume persists data. Do not delete it unless you deliberately want a
fresh course database.

## Project Map

| Project | Main skill |
|---|---|
| 01 Server And psql | connect and inspect |
| 02 Schema And Types | model tables |
| 03 Insert And Select | create/read rows |
| 04 Update And Delete | mutate safely |
| 05 Constraints | reject invalid state |
| 06 Relationships And Joins | reconstruct related data |
| 07 Aggregations | build reports |
| 08 Advanced Queries | NULL, CTEs, windows |
| 09 Transactions | make writes atomic |
| 10 Concurrency And Locks | observe overlapping work |
| 11 Indexes And EXPLAIN | optimize measured queries |
| 12 Data Modeling | normalize real domains |
| 13 Migrations | evolve without breaking code |
| 14 Roles And Security | least privilege and parameters |
| 15 Backup And Restore | prove recovery |
| 16 Go Repository | use `database/sql` correctly |
| 17 Task API Capstone | complete PostgreSQL backend |

## Exercise Method

Every project includes:

- focused `README.md`
- starter SQL or Go artifacts
- verification queries
- deliberate failure cases
- a completion standard

Before every mutation:

1. predict affected rows
2. select the target rows
3. run inside a transaction while learning
4. inspect returned/affected rows
5. commit only after verification

## Security Boundary

The included credentials are local course credentials. They are not production
examples. Production systems require secret management, TLS, restricted
networking, monitored backups, and role separation.

