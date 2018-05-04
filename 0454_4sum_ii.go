// O(N^2) Solution
package main

import (
	"fmt"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	Sum := make(map[int]int)
	count := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			Sum[A[i]+B[j]]++
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			v := -C[i] - D[j]
			count += Sum[v]
		}
	}

	return count
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
