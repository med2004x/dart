# Project 01 - PostgreSQL Server And psql

## Goal

Start PostgreSQL, connect with `psql`, and distinguish server, database, schema,
role, and client.

## Run

```powershell
Set-Location C:\Users\pc\Documents\dart\postgresql-exercises
docker compose up -d
docker compose ps
docker compose logs postgres
.\scripts\connect.ps1
```

Inside `psql`:

```text
\conninfo
\l
\dn
\dt
\du
SELECT current_database(), current_user, version();
\q
```

Backslash commands belong to `psql`; SQL statements go to PostgreSQL and end
with `;`.

## Verification

Run:

```powershell
.\scripts\run-sql.ps1 .\01-server-psql\verify.sql
```

## Failure Drills

1. Stop the container and connect.
2. use the wrong database name.
3. start another service on host port 5432.
4. omit a semicolon, then clear the query buffer with `\r`.

Classify connection refused, authentication failure, missing database, and port
allocation as different failures.

## Done Means

You can name every connection parameter and inspect the current session without
relying on hidden defaults.

