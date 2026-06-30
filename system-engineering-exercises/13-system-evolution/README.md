# Project 13 - Evolve A System In Stages

## Goal

Build four stages of one task system. Add complexity only when a new requirement
proves the previous stage insufficient.

## Stages

```text
stage1-cli-file/
stage2-http-postgres/
stage3-multi-instance/
stage4-outbox-worker/
```

## Required Proof

Stage 1:

```text
one local user
tasks survive restart
corrupt file fails safely
```

Stage 2:

```text
concurrent HTTP clients
shared PostgreSQL state
constraints and transactions
```

Stage 3:

```text
two application processes on different ports
both see the same tasks
readiness controls traffic
pool budget covers both
```

Stage 4:

```text
notification provider is outside create-task latency
outbox survives worker restart
duplicate delivery is safe
```

## For Every Stage

Complete one copy of `stage-template.md`:

- triggering requirement
- diagram
- commands
- success evidence
- failure evidence
- new operational burden
- migration path
- rollback path

## Rule

Do not start a later stage until the current stage's proof is executable.

## Done Means

Every added process or storage system solves one requirement the previous stage
could not meet.

