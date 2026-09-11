package main

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

// TODO: add CreateTask, GetTask, and ListTasks.
