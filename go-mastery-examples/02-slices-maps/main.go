package main

import "fmt"

type Product struct {
    Name     string
    Category string
}

func CountByCategory(products []Product) map[string]int {
    counts := make(map[string]int)
    for _, p := range products {
        counts[p.Category]++
    }
    return counts
}

func main() {
    products := []Product{
        {Name: "T-shirt", Category: "clothes"},
        {Name: "Shoes", Category: "clothes"},
        {Name: "Phone", Category: "electronics"},
    }
    fmt.Println(CountByCategory(products))
}
