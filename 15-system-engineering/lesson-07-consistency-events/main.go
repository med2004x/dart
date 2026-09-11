package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID    int
	Title string
}

type OutboxEvent struct {
	Task      Task
	Delivered bool
}

type Store struct {
	mu         sync.Mutex
	tasks      map[int]Task
	outbox     []OutboxEvent
	projection map[int]Task
	nextID     int
}

func NewStore() *Store {
	return &Store{
		tasks:      make(map[int]Task),
		projection: make(map[int]Task),
		nextID:     1,
	}
}

// TODO: implement CreateTask and ProcessPendingEvents atomically and
// idempotently.

func main() {
	store := NewStore()
	fmt.Println("primary tasks:", len(store.tasks))
	fmt.Println("search projection:", len(store.projection))
	time.Sleep(10 * time.Millisecond)
}
