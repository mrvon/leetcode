package main

import "fmt"

func component(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var c []int
	for n > 0 {
		r := n % 10
		n = n / 10
		c = append(c, r)
	}
	return c
}

func isHappy(n int) bool {
	loop := make(map[int]int)

	for loop[n] == 0 {
		if n == 1 {
			return true
		}

		loop[n]++
		c := component(n)
		n = 0
		for i := 0; i < len(c); i++ {
			n += c[i] * c[i]
		}
	}

	return false
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(8))
}
