package main

import "fmt"

func buildPrimeTable(n int) []int {
	// 0 (is a prime number)
	// -1 (not a prime number)
	prime := make([]int, n+1)
	prime[0] = -1
	prime[1] = -1
	for i := 2; i <= n; i++ {
		if prime[i] == -1 {
			continue
		}

		for j := 2; i*j <= n; j++ {
			// mark it
			prime[i*j] = -1
		}
	}
	return prime
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

func main() {
	t := buildPrimeTable(20)
	for i := 0; i < len(t); i++ {
		if t[i] == 0 {
			fmt.Printf("Number %d is Prime Number\n", i)
		} else {
			fmt.Printf("Number %d is not Prime Number\n", i)
		}
	}
}
