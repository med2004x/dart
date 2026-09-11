package main

import "fmt"

func frequencies(values []string) map[string]int {
	// TODO: normalize and count.
	return map[string]int{}
}

func mostCommon(counts map[string]int) (string, int) {
	// TODO: use a deterministic tie-break.
	return "", 0
}

func main() {
	counts := frequencies([]string{"open", "done", "OPEN"})
	status, count := mostCommon(counts)
	fmt.Println(counts, status, count)
}
