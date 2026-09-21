package main

// fmt is from Go's standard library.
// You use it to print text and values to the terminal.
// Docs: https://pkg.go.dev/fmt
import "fmt"

/*
Topic: slices, loops, and calculations

Learn first:
- Arrays and slices: https://go.dev/tour/moretypes/6
- Slices: https://go.dev/tour/moretypes/7
- Range loops: https://go.dev/tour/moretypes/16

Your task:
Build a student scoreboard.

Requirements:
 1. Create a Student struct.
    It should store:
    - name
    - scores []int

 2. Write a method named Average.
    It should:
    - loop through the scores
    - calculate the average
    - return a float64

 3. Write a method named Passed.
    It should:
    - return true when average is 50 or higher
    - return false otherwise

4. In main:
  - create at least 3 students
  - print each student's name
  - print their average
  - print whether they passed
*/
func main() {
	fmt.Println("Exercise 03: write the scoreboard program here")
}
