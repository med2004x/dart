# Project 03 - OpenAPI Contract

## Goal

Describe one API operation in a machine-readable OpenAPI document before
implementation.

## Starter

`openapi.yaml` contains a health endpoint and a task schema placeholder.

## Checkpoints

1. define `POST /projects/{projectId}/tasks`.
2. define path parameter.
3. define request body.
4. mark required fields.
5. define 201 response.
6. define reusable 400, 401, 404, and 500 errors.
7. add examples for every response.
8. define `GET` collection behavior.

## Contract Rules

- request and response schemas are different when ownership differs
- timestamps state format and timezone semantics
- nullable and optional are not the same
- every non-2xx response has a stable body
- examples must satisfy schemas

## Failure Drill

Add a field to an example but not its schema. Then add an undocumented 409 from
the implementation. Explain why clients cannot rely on undocumented behavior.

## Done Means

The contract is sufficient to build a client mock without reading server code.

