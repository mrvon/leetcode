package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	divisors := []int{1}

	i := 2

	for ; i*i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i, num/i)
		}
	}

	if i*i == num {
		divisors = append(divisors, i)
	}

	for i := 0; i < len(divisors); i++ {
		num -= divisors[i]
	}

	return num == 0
}

func main() {
	fmt.Println(checkPerfectNumber(4))
	fmt.Println(checkPerfectNumber(28))
}
