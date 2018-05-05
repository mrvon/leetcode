// dynamic programming
package main

import (
	"fmt"
	"math"
)

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func integerBreak(n int) int {
	optimal := make([]int, n+1)

	if n < 2 {
		panic("invalid n")
	}

	optimal[0] = 0
	optimal[1] = 0
	optimal[2] = 1

	for i := 3; i <= n; i++ {
		m := math.MinInt32
		for j := 1; j <= i-1; j++ {
			p1 := j * optimal[i-j] // break
			p2 := j * (i - j)      // no break
			m = max(max(p1, p2), m)
		}
		optimal[i] = m
	}

	return optimal[n]
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(3))
	fmt.Println(integerBreak(4))
	fmt.Println(integerBreak(5))
	fmt.Println(integerBreak(6))
	fmt.Println(integerBreak(7))
	fmt.Println(integerBreak(8))
}
