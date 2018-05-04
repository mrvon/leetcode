package main

import "fmt"

func countPrimes(n int) int {
	prime := make([]int, n)
	for i := 2; i < n; i++ {
		if prime[i] == -1 {
			// not a prime
			continue
		}

		for j := 2; i*j < n; j++ {
			// mark it
			prime[i*j] = -1
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if prime[i] == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(4))
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(6))
	fmt.Println(countPrimes(100))
}
