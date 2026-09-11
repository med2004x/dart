package main

import "fmt"

func classify(values []int) (below, zero, above int) {
	for _,value:= range values {
		if value <0 {
			below += 1
		} else if value == 0 {
			zero += 1
		} else {
			above += 1
		}
	}
	return below, zero, above
}

func main() {
	below, zero, above := classify([]int{-2, 0, 5, -1, 3})
	fmt.Println(below, zero, above)
}
