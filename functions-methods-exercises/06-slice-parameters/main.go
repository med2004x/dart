package main

func countItems(items []string) int {
	return len(items)
}

func totalScores(scores []int) int {
	sum :=0
	for _,score := range scores {
		sum = sum+score
	}
	return sum
}
func highestScore(scores []int) int {
	max := 0
	for _, score := range scores {
		if score>max {
			max = score
		}
	}
	return max
}
func containsItem(items []string, wantedItem string) bool {
	for _, item := range items {
		if item == wantedItem {
			return true
		}
	}
	return false
}


func main() {
	print(countItems([]string {"AAd", "ads", "ee"}),"\n")
	print(totalScores([]int{4,5,5,5,5,7,8,9,88}),"\n")
	print(highestScore([]int{77,88,99,10000}),"\n")
	print(containsItem([]string{"a", "ba","dd","c"}, "a"),"\n")
	print(countItems([]string {}),"\n")
	print(totalScores([]int{}),"\n")
	print(highestScore([]int{}),"\n")
	print(containsItem([]string{}, "a"),"\n")

}
