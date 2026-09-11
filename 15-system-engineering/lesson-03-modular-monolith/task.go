package main

import "errors"

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTitleRequired = errors.New("title is required")
)

type Task struct {
	ID    int
	Title string
	Done  bool
}
