package main

import "fmt"

func smallestRangeI(A []int, K int) int {
	min := A[0]
	max := A[0]
	for _, a := range A {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	if max-min <= 2*K {
		return 0
	}
	return max - min - 2*K
}

func main() {
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
	fmt.Println(smallestRangeI([]int{1, 3, 7}, 3))
	fmt.Println(smallestRangeI([]int{1, 3, 8}, 3))
}
