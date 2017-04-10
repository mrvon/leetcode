// Dynamic programming solution
package main

import (
	"fmt"
	"math"
)

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// DP
func numSquares(n int) int {
	optimal := make([]int, n+1)
	optimal[0] = 0

	for i := 1; i <= n; i++ {
		min_num := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			// try to split it, find minimum.
			// min_num = min(min_num, optimal[j*j]+optimal[i-j*j])
			// optimal[j*j] must be 1
			min_num = min(min_num, 1+optimal[i-j*j])
		}
		optimal[i] = min_num
	}

	return optimal[n]
}

func bruteforce_numSquares(n int) int {
	s := int(math.Sqrt(float64(n)))
	if s*s == n {
		// need'n split.
		return 1
	}

	min_num := math.MaxInt32
	for i := 1; i <= n/2; i++ {
		// try to split it, find minimum.
		min_num = min(min_num, numSquares(i)+numSquares(n-i))
	}
	return min_num
}

func main() {
	fmt.Println(numSquares(1))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(5))
	fmt.Println(numSquares(6))
	fmt.Println(numSquares(7))
	fmt.Println(numSquares(8))
	fmt.Println(numSquares(16))
	fmt.Println(numSquares(1535))
}
