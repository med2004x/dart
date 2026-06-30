package main

type TaskRepository interface {
	Create(Task) (Task, error)
	FindByID(int) (Task, error)
	List() ([]Task, error)
}

type MemoryTaskRepository struct {
	tasks  []Task
	nextID int
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{nextID: 1}
}

// TODO: implement Create, FindByID, and List.
