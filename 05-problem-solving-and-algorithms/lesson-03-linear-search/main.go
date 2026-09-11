package main

import "fmt"

type Task struct {
	ID    int
	Title string
}

func findTask(tasks []Task, wantedID int) (Task, bool) {
	// TODO: implement linear search.
	return Task{}, false
}

func main() {
	tasks := []Task{{ID: 4}, {ID: 9}, {ID: 2}}
	fmt.Println(findTask(tasks, 2))
}
