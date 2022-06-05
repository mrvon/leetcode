package main

import "fmt"

func isSelfDividing(n int) bool {
	m := n

	if m <= 0 {
		return false
	}

	digits := []int{}
	for m > 0 {
		r := m % 10
		if r == 0 {
			return false
		}
		m = m / 10
		digits = append(digits, r)
	}
	for i := 0; i < len(digits); i++ {
		if n%digits[i] != 0 {
			return false
		}
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	numbers := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func main() {
	// left = 1, right = 22
	// Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
	fmt.Println(selfDividingNumbers(1, 22))
}
