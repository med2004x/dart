package main

	import "fmt"


type cartItem struct {
	name     string
	quantity int
}

func addItem(cart []cartItem, newItem []cartItem) []cartItem {
	return append(cart, newItem...)
}
func increaseQ(cart []cartItem, item string) []cartItem {
	for i := range cart {
		if cart[i].name == item {
			cart[i].quantity++
		}
	}
	return cart
}
func remove(cart []cartItem, item string) []cartItem {
    for i := range cart {
        if cart[i].name == item {
            return append(cart[:i], cart[i+1:]...)
        }
    }
    return cart
} 
func totalQ(cart []cartItem) int {
	tot := 0
	for _, item := range cart {
		tot += item.quantity
	}
	return tot
}

func main() {
	carts := []cartItem {}
carts = addItem(carts,[]cartItem{{name: "vegies", quantity: 4},
{name : "fruits", quantity: 99},
{name: "orange", quantity: 78},
})
fmt.Print(carts, "\n")
	carts = increaseQ(carts, "vegies")
fmt.Print(carts, "\n")
    carts = remove(carts, "fruits")
fmt.Print(carts, "\n")
fmt.Print(totalQ(carts), "\n")
}
