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

Example output:
Sara average: 73.33 passed: true
Youssef average: 44.00 passed: false

Strict rule:
Do not hardcode the average. It must be calculated from the slice.
*/
type student struct {
   name string
   scores []int
}
func (stud student) Average() float64 {
   if len(stud.scores) == 0 { 
      return 0.0
   }
sum:=0
for _,score := range stud.scores {
   sum += score
}
avrg := float64(sum)/float64(len(stud.scores))
return float64(avrg)
}
func (stud student) Passed() bool {
   if stud.Average()>= 50 {
      return true
   } else {
      return false
   }
}
func (stud student) Stat() string {
   if stud.Passed()== true {
      return ("Passed")
   } else {
      return ("didn't pass")
   }
}

func main() {
   students := []student{
		{name: "Sara", scores: []int{70, 80, 75}}, // Average: 75.0
		{name: "Youssef", scores: []int{40, 44, 50}}, // Average: 44.66
		{name: "Ahmed", scores: []int{90, 85, 78, 92}}, // Average: 86.25
	}
   for _,stud := range students {
      avg := stud.Average()
      passedStat := stud.Stat()
   
   
   fmt.Printf("The student %s has an average of %f, and %s.\n", stud.name ,avg,passedStat )
} }