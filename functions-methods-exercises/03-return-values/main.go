package main

	func add(int1, int2 int) int {
		return int1+int2
	}
	func prod(num1, num2 int) int {
		return num1 * num2
	}
	func discounted(price, discount float32) float32 {
		return price*(1-discount/100)
	}
	func isPassing(score int) bool {
		if score>=50 { return true }
		return false
	}

func main() {
	sum := add(2, 99)
	product := prod(99, 69)
	discount := discounted(79.99, 10)
	passing := isPassing(49)
	print(sum,"\n", product,"\n", discount,"\n", passing)
}
