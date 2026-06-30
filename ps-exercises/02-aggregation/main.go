package main

import "fmt"

type Summary struct {
	Minimum int
	Maximum int
	Total   int
	Average float64
}

func summarize(values []int) (Summary, bool) {
	// TODO: implement one-pass aggregation.
	return Summary{}, false
}

func main() {
	summary, found := summarize([]int{30, 10, 50})
	fmt.Println(summary, found)
}
