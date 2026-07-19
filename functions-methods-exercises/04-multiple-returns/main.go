package main

import "slices"

func findNum(nums []int, wNum int) bool {
	for _, num :=range nums {
		if num == wNum {
			return true
		}
	}
	return false
}
func findName(names []string, wName string) bool{
	return slices.Contains(names, wName)
}
func safeDiv(num, denom int) bool {
	if num/denom != 0 {
		return false
	}
	return true
}
func main() {
	print(findNum([]int{1, 3, 4, 5, 5, 5, 6, 9, 8, 7, 8}, 4),"\n")
	print(findName([]string{"Adam", "Anwar", "Akram"}, "Akram"),"\n")
	print(safeDiv(44, 8),"\n")
	print(findNum([]int{1, 3, 4, 5, 5, 5, 6, 9, 8, 7, 8}, 99),"\n")
	print(findName([]string{"Adam", "Anwar", "Akram"}, "Abram"),"\n")
	print(safeDiv(44, 78),"\n")
}
