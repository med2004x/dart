# Project 02 - Resource And URL Modeling

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Model stable resources and ownership before writing route code.

## Resources

```text
users
projects
project memberships
tasks
task comments
```

## Design Rules

- URLs identify resources, not implementation functions.
- HTTP methods describe operations.
- server-owned IDs never come from trusted client claims.
- nesting should express ownership only when useful.
- do not expose database table layout accidentally.

Candidate routes:

```text
POST   /projects
GET    /projects/{projectId}
GET    /projects/{projectId}/tasks
POST   /projects/{projectId}/tasks
GET    /tasks/{taskId}
PATCH  /tasks/{taskId}
DELETE /tasks/{taskId}
```

## Tasks

1. Complete `resource-model.md`.
2. decide whether task URL needs project nesting.
3. define create/update input fields.
4. identify server-owned fields.
5. define ownership lookup for every protected route.
6. reject action URLs such as `/doCreateTask`.

## Failure Drill

Design authorization using only `{projectId}` supplied in the URL. Show why the
server must still verify the stored task/project relationship.

## Done Means

Every route has one resource meaning, owner, and authorization source.

