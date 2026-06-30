package main

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var ErrTaskNotFound = errors.New("task not found")

type Task struct {
	ID        int64
	ProjectID int64
	Title     string
	Status    string
	Priority  int16
	DueAt     sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostgresTaskRepository struct {
	database *sql.DB
}

func NewPostgresTaskRepository(database *sql.DB) PostgresTaskRepository {
	return PostgresTaskRepository{database: database}
}

func (repository PostgresTaskRepository) FindByID(
	ctx context.Context,
	id int64,
) (Task, error) {
	// TODO: use QueryRowContext with $1 and map sql.ErrNoRows.
	return Task{}, ErrTaskNotFound
}

// TODO: implement Create and ListByProject with parameterized SQL.
