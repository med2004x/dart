package main

import "fmt"

type Graph map[string][]string

func shortestPath(graph Graph, start, target string) ([]string, bool) {
	// TODO: BFS with queue, visited set, and predecessor map.
	return nil, false
}

func main() {
	graph := Graph{
		"A": {"B", "C"},
		"B": {"A", "D"},
		"C": {"A", "D"},
		"D": {"B", "C"},
	}
	fmt.Println(shortestPath(graph, "A", "D"))
}
