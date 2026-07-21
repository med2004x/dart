# Project 14 - Roles, Privileges, And SQL Injection

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Run the application with only required data privileges and use parameterized
queries for all untrusted values.

## Roles

The local `student` role owns course objects. Create `task_app` for runtime use.

Required runtime privileges:

- connect to `go_course`
- use schema `app`
- select/insert/update/delete application tables
- use identity sequences

Forbidden:

- drop tables
- create roles
- change schema
- read unrelated schemas

## Checkpoints

1. complete `roles.sql`.
2. connect as `task_app`.
3. run permitted operations.
4. run every denied operation in `verify.sql`.
5. configure default privileges for future objects.
6. write one parameterized Go query.
7. validate dynamic sort fields through a fixed allowlist.

## Failure Drills

1. build SQL by concatenating a title containing a quote.
2. grant all privileges.
3. log a connection URL containing the password.
4. trust an unsafe writable schema in `search_path`.

## Done Means

Compromising the application credential does not grant database administration,
and user input cannot become SQL structure.

