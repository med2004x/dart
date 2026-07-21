# Project 01 - API Requirements

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Define the user and system problem before choosing endpoints.

## Scenario

Design an API for teams managing projects and tasks.

Users need to:

- create projects
- invite project members
- create and update tasks
- list tasks by status
- assign tasks to members

## Complete The Brief

Use `api-brief.md`.

Define:

- clients and trust level
- core workflows
- resource ownership
- peak request rate
- maximum request/response size
- latency target
- availability target
- consistency requirements
- data classification
- retrying clients

## Failure Questions

- What happens if the client repeats a create?
- What happens if two users update one task?
- Can users enumerate another team's tasks?
- What happens when a dependency times out?
- Which operations may complete asynchronously?

## Done Means

Every later endpoint can be traced to a written workflow and quality
requirement.

