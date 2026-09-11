package main

import "fmt"

type contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func addContact(contacts []contact, newContact contact) []contact {
	return append(contacts, newContact)
}

func findContactByID(contacts []contact, wantedID int) (contact, bool) {
	for _, c := range contacts {
		if c.ID == wantedID {
			return c, true
		}
	}
	return contact{}, false
}

func updateContactEmail(contacts []contact, wantedID int, email string) ([]contact, bool) {
	for i := range contacts {
		if contacts[i].ID == wantedID {
			contacts[i].Email = email
			return contacts, true
		}
	}
	return contacts, false
}

func deleteContact(contacts []contact, wantedID int) ([]contact, bool) {
	result := make([]contact, 0, len(contacts))
	found := false

	for _, c := range contacts {
		if c.ID == wantedID {
			found = true
			continue
		}
		result = append(result, c)
	}

	if !found {
		return contacts, false
	}
	return result, true
}

func printContacts(contacts []contact) {
	for _, c := range contacts {
		fmt.Printf("%d %s %s %s\n", c.ID, c.Name, c.Email, c.Phone)
	}
}

func main() {
	contacts := []contact{}

	contacts = addContact(contacts, contact{ID: 1, Name: "Adam", Email: "adam@example.com", Phone: "111"})
	contacts = addContact(contacts, contact{ID: 2, Name: "Sara", Email: "sara@example.com", Phone: "222"})

	found, ok := findContactByID(contacts, 1)
	fmt.Println("find:", found, ok)

	var updated bool
	contacts, updated = updateContactEmail(contacts, 1, "new-adam@example.com")
	fmt.Println("update:", updated)

	var deleted bool
	contacts, deleted = deleteContact(contacts, 2)
	fmt.Println("delete:", deleted)

	printContacts(contacts)
}
