# Project 06 - Relationships And Joins

## Goal

Reconstruct related data while predicting result cardinality.

## Add Many-To-Many Tags

Create:

```text
tags
task_tags(task_id, tag_id)
```

Use a composite primary key to prevent duplicate task/tag pairs.

## Required Queries

1. every task with project and owner names
2. every task owned by one user
3. every project including empty projects
4. every tag for one task
5. every task using one tag
6. users with no projects

## Join Rule

Before running a join, write:

```text
left input rows:
right matches per row:
expected output rows:
```

An inner join removes unmatched rows. A left join keeps all left rows and uses
`NULL` for missing right values.

## Failure Drills

1. omit a join condition and observe the Cartesian product.
2. use inner join when empty parents must remain.
3. count `*` after a many-to-many join.
4. insert the same task/tag pair twice.

## Done Means

You can predict output row count and explain every repeated parent value.

