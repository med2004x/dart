package main

import "fmt"

func classify(values []int) (below, zero, above int) {
	// TODO: classify every value exactly once.
	return 0, 0, 0
}

func main() {
	below, zero, above := classify([]int{-2, 0, 5, -1, 3})
	fmt.Println(below, zero, above)
}
