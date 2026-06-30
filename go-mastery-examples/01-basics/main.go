package main

import "fmt"

type ProductStatus string

const (
    StatusDraft     ProductStatus = "draft"
    StatusPublished ProductStatus = "published"
)

func main() {
    name := "T-shirt"
    priceCents := int64(4900)
    status := StatusPublished

    fmt.Printf("%s costs %d cents and is %s\n", name, priceCents, status)
}
