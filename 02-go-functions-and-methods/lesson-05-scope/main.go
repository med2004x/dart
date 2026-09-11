package main

func userName(firstName, lastName string) string {
	return ("Your username is : "+firstName+lastName+".")
}
func calculateTotal(amount, tax float64)float64 {
	return amount*(1+tax/100)
}
func receipt(customerName, total string) string {
	return ("customer name: "+ customerName +"\n Amount to pay: "+total+"\n")
}


func main() {
	firstName := "Maad"
	lastName := "Miid"
	amount:= 99.99
	tax:= 19
	customerName := "fill"
	total := "59.99"
	print(userName(firstName, lastName) ,"\n")
	print(calculateTotal(amount, float64(tax)),"\n")
	print(receipt(customerName, total),"\n")
}
