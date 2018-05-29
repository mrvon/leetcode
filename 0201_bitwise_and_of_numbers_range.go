/*
Consider the from low to high, if n > m, the lowest bit will be 0.  and then we
could translate the problem to sub-problem.
*/
package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	if n <= m {
		return m
	} else {
		return rangeBitwiseAnd(m>>1, n>>1) << 1
	}
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(1, 31))
	fmt.Println(rangeBitwiseAnd(20000, 2147483647))
}
