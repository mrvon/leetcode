package main

import "fmt"

func counterSetBits(num int) int {
	c := 0
	for num > 0 {
		num = num & (num - 1)
		c++
	}
	return c
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(L int, R int) int {
	count := 0
	for i := L; i <= R; i++ {
		if isPrime(counterSetBits(i)) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}
