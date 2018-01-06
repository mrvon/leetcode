package main

import "fmt"

func hasAlternatingBits(n int) bool {
	first_bit := true
	last_bit := 0
	for n > 0 {
		if !first_bit && n&0x1 == last_bit {
			return false
		}
		last_bit = n & 0x1
		first_bit = false
		n = n >> 1
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(11))
}
