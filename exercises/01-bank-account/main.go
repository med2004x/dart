package main

// fmt is from Go's standard library.
// You use it to print text and values to the terminal.
// Docs: https://pkg.go.dev/fmt
import "fmt"

/*
Topic: structs and methods

Learn first:
- Structs: https://go.dev/tour/moretypes/2
- Methods: https://go.dev/tour/methods/1
- Pointer receivers: https://go.dev/tour/methods/4

Your task:
Build a small bank account program from scratch.

Requirements:
1. Create an Account struct.
   It should store:
   - owner name
   - balance

2. Write a Deposit method.
   It should:
   - receive an amount
   - reject zero or negative amounts
   - add valid amounts to the balance

3. Write a Withdraw method.
   It should:
   - reject zero or negative amounts
   - reject withdrawals bigger than the balance
   - subtract valid amounts from the balance

4. Write a PrintSummary method.
   It should print:
   Account owner: Sara
   Balance: 70

5. In main:
   - create one account
   - deposit 100
   - withdraw 30
   - print the summary

Strict rule:
Do not add another package yet. Only use fmt.
*/
type Account struct{
Owner string
Balance int
}

func (account Account) addToBalance(amount int)int {
 if amount <=0 {
   return account.Balance
 } else {
   return account.Balance + amount
 }
}
func (account Account) withdrawFromBalance(withdrawal int)int {
   if withdrawal <=0{
      return account.Balance
   } else if withdrawal>account.Balance {
     return account.Balance
   } else {
      return account.Balance-withdrawal
   }
}
func (account Account) printSummary() {
	fmt.Println("Account owner:", account.Owner)
	fmt.Println("Balance:", account.Balance)
}



func main() {
	fmt.Println("Exercise 01: write the bank account program here")
   account:= Account{
      Owner: "Adam",
      Balance: 80,
   }
   account.Balance = account.addToBalance(40)
   account.Balance =account.withdrawFromBalance(20)
   account.printSummary()
}

