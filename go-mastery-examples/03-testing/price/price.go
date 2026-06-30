package price

func AddTax(subtotalCents int, taxRate float64) int {
    return subtotalCents + int(float64(subtotalCents)*taxRate)
}
