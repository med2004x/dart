package main

// fmt is from Go's standard library.
// You use it to print text and values to the terminal.
// Docs: https://pkg.go.dev/fmt
import "fmt"

/*
Topic: slices, loops, and simple login logic

Learn first:
- Slices: https://go.dev/tour/moretypes/7
- Range loops: https://go.dev/tour/moretypes/16
- If statements: https://go.dev/tour/flowcontrol/5

Your task:
Build a small login checker without maps and without errors package.
That is intentional. First learn the logic with loops.

Requirements:
1. Create a User struct.
   It should store:
   - username
   - password
   - role

2. Create a slice with at least 3 users.

3. Write a function named Authenticate.
   It should receive:
   - users []User
   - username string
   - password string

   It should return:
   - User
   - bool

4. Authentication rules:
   - if username and password match one user, return that user and true
   - otherwise return an empty User and false

5. In main:
   - call Authenticate with a correct login
   - call Authenticate with a wrong login
   - print different messages for success and failure

Expected behavior:
Correct login prints something like:
Welcome admin, role: owner

Wrong login prints:
Invalid username or password

Strict rule:
Do not use maps yet. Do not use the errors package yet.
*/
type user struct {
   username string
   password string
   role string
}


func main() {
	fmt.Println("Exercise 02: write the auth program here")
users:= []user {
  {
   username : "Med",
   password : "Abc123",
   role : "Admin",
  }, {
 username : "Mad",
   password : "Abc1234",
   role : "user",
  }, {
    username : "Medi",
   password : "Abcd123",
   role : "user",
  },
}
authenticate(users ,"Med", "Abc123")
authenticate(users ,"Maed", "Abc1523")

}
func authenticate(users []user, username,password string)string {
   for _, oneUser := range users {
    if oneUser.password == password && oneUser.username == username {
        return ("true") 
    }
}
return ("false")
}