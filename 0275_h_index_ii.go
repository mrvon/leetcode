package main

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	h := 0
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2

		if citations[mid] >= n-mid {
			h = n - mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return h
}

func main() {
	fmt.Println(hIndex([]int{1}))
	fmt.Println(hIndex([]int{0, 1}))
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
}
