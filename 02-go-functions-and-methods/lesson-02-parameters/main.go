package main

func greeting(name string) {
	print("Hello " + name + ".\n")
}
func userBadge(name, role string) {
	print("Hello :"+ name + "\nYour role is: "+ role + "\n")
}
func productLine(productName, quantity string){
	print("You ordered "+ quantity + " pieces of "+productName)
}

func main() {
	// Build the parameter exercises described in README.md.
	greeting("med")
	userBadge("Med", "admin")
	productLine("pantie", "4")
}
