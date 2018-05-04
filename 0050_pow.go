package main

import "fmt"

func pow(x float64, n int, memoized map[int]float64) float64 {
	if n == 0 {
		return 1.0
	} else if n == 1 {
		return x
	}

	var t int = n / 2
	var r int = n % 2

	p, b := memoized[t]
	if b {
		p = memoized[t]
	} else {
		p = pow(x, t, memoized)
	}

	if r > 0 {
		return p * p * x
	} else {
		return p * p
	}
}

func myPow(x float64, n int) float64 {
	memoized := make(map[int]float64)
	if n >= 0 {
		return pow(x, n, memoized)
	} else {
		return 1.0 / pow(x, -n, memoized)
	}
}

func main() {
	fmt.Println(myPow(1, 10))
	fmt.Println(myPow(2, 2))
	fmt.Println(myPow(3, 0))
	fmt.Println(myPow(4, 4))
	fmt.Println(myPow(2, -1))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(2, -3))
	fmt.Println(myPow(2, -4))
	fmt.Println(myPow(2, -5))
	fmt.Println(myPow(34.00515, -3))
	fmt.Println(myPow(0.44, 0))
	fmt.Println(myPow(0.44, 1))
}
