package main

// fmt is from Go's standard library.
// You use it to print text and values to the terminal.
// Docs: https://pkg.go.dev/fmt
import "fmt"

/*
Topic: maps and grouping data

Learn first:
- Maps: https://go.dev/tour/moretypes/19
- Mutating maps: https://go.dev/tour/moretypes/22

Your task:
Build an expense tracker.

Requirements:
 1. Create an Expense struct.
    It should store:
    - title
    - category
    - amount

 2. Create a slice with at least 5 expenses.
    Example categories:
    - food
    - transport
    - study

 3. Write a function named Total.
    It should:
    - receive []Expense
    - return the total amount

 4. Write a function named TotalByCategory.
    It should:
    - receive []Expense
    - return map[string]int
    - group amounts by category

5. In main:
  - print the total
  - print every category total

Example output:
Total: 4500
food: 1200
transport: 800
study: 2500

Strict rule:
Amounts are integers. Do not use float64 for money in this exercise.
*/
func main() {
	fmt.Println("Exercise 04: write the expense tracker here")
}
