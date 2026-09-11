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
type expense struct {
   title string
   category string
   amount int
}



func main() {
	fmt.Println("Exercise 04: write the expense tracker here")
  
 expenses := [] expense {
   {title: "grocery" , category: "food" , amount: 50},
   {title: "car repair" , category: "transport" , amount: 70},
   {title: "gaz" , category: "transport" , amount: 90},
   {title: "net" , category: "study" , amount: 99},
   {title: "investment" , category: "study" , amount: 40},
}
	total := Total(expenses)
	fmt.Println("Total:", total) 

	fmt.Println("--- Category Totals ---")
	
	// 4. Call the TotalByCategory function to get category totals
	categoryTotals := TotalByCategory(expenses)
	
	// Print every category total
	for category, amount := range categoryTotals {
		fmt.Printf("%s: %d\n", category, amount)
	}
}
func Total(expenses []expense) int {
   total:=0
   for _,exps := range expenses {
      total+= exps.amount
   }
   return total
}
func TotalByCategory(expenses []expense) map[string]int {
	// Initialize an empty map to store category totals
	categoryTotals := make(map[string]int)

	// Loop through the expense slice
	for _, exps := range expenses {
		// Use the category as the key and add the amount to its running total.
		// If the category doesn't exist yet, Go initializes it to 0 before adding the value.
		categoryTotals[exps.category] += exps.amount
	}

	return categoryTotals
}

