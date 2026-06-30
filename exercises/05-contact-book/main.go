package main

// fmt prints values to the terminal.
// Docs: https://pkg.go.dev/fmt

	import "fmt"



/*
Build the contact book described in README.md.

Strict rules:
- Use a slice, not a map.
- Return bool when find/update/delete succeeds or fails.
- Do not use HTTP yet.
*/

type contact struct {
	ID int
	name string
	number int
	adress string
	mail string
}
func  addContact(contacts []contact, newContacts []contact) []contact {
	contacts = append(contacts, newContacts...)
	return contacts
}
func main (){
	contacts:= []contact {
		
	}
	for _,cont := range contacts {
		fmt.Printf(cont )
	}
}
