package main

import (
	"errors"
	"fmt"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrForbidden    = errors.New("forbidden")
)

type Task struct {
	ID      int
	OwnerID int
	Title   string
}

func getTaskForActor(tasks []Task, actorID, taskID int) (Task, error) {
	// TODO: find by task ID, then authorize with the stored OwnerID.
	return Task{}, ErrTaskNotFound
}

func main() {
	tasks := []Task{{ID: 1, OwnerID: 1, Title: "private task"}}
	task, err := getTaskForActor(tasks, 1, 1)
	fmt.Println(task, err)
}
