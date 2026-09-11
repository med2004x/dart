package main

import "fmt"

func validBrackets(input string) bool {
	// TODO: use a stack of opening brackets.
	return false
}

type Queue[T any] struct {
	values []T
	head   int
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.values = append(queue.values, value)
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	// TODO: return values in first-in-first-out order.
	var zero T
	return zero, false
}

func main() {
	fmt.Println(validBrackets("([{}])"))
}
