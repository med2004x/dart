package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) ([]Interval, error) {
	// TODO: validate, copy, sort, and merge.
	return nil, nil
}

func main() {
	fmt.Println(mergeIntervals([]Interval{
		{Start: 3, End: 6},
		{Start: 1, End: 4},
		{Start: 8, End: 10},
	}))
}
