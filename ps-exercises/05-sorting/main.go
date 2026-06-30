package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Priority  int
	CreatedAt time.Time
}

func sortedTasks(tasks []Task) []Task {
	// TODO: copy tasks, then sort with all tie-breakers.
	return nil
}

func main() {
	now := time.Now()
	tasks := []Task{
		{ID: 2, Priority: 1, CreatedAt: now},
		{ID: 1, Priority: 1, CreatedAt: now},
	}
	fmt.Println(sortedTasks(tasks))
}
