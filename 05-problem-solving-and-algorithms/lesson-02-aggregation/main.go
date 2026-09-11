package main

import "fmt"

type Summary struct {
	Minimum int
	Maximum int
	Total   int
	Average float64
}

func summarize(values []int) (Summary, bool) {
	if len(values) == 0  {
		return Summary{}, false
	}
	total :=0
	min := values[0]
	max  := values[0]
	average := 0.0
	for _,value := range values {
		
		if value< min {
			min = value
		} else if value >max {
			max = value
		}
		total = total + value
		average = float64(total/(len(values)))
	}

	
	return Summary{
		Minimum: min,
		Maximum: max,
		Total: total,
		Average: average,
	}, true
}

func main() {
	summary, found := summarize([]int{})
	fmt.Println(summary, found)
}
