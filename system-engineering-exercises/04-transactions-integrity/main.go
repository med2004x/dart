package main

import (
	"errors"
	"fmt"
)

var errDiskFull = errors.New("disk full")

type Task struct {
	ID    int
	Title string
}

type TaskWriter interface {
	Save([]Task) error
}

type FailingWriter struct{}

func (FailingWriter) Save([]Task) error {
	return errDiskFull
}

// TODO: implement an atomic Repository using candidate state.

func main() {
	fmt.Println("Implement atomic create and prove failure leaves no task.")
}
