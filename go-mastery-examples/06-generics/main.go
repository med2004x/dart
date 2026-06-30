package main

import "fmt"

type Set[T comparable] struct {
    items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) { s.items[value] = struct{}{} }
func (s *Set[T]) Has(value T) bool { _, ok := s.items[value]; return ok }
func (s *Set[T]) Remove(value T) { delete(s.items, value) }

func main() {
    categories := NewSet[string]()
    categories.Add("clothes")
    categories.Add("electronics")
    fmt.Println(categories.Has("clothes"))
}
