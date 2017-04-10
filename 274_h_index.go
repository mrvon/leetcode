package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h := 0
	left := 0
	right := len(citations) - 1

	for left <= right {
		mid := left + (right-left)/2

		if citations[mid] >= mid+1 {
			h = mid + 1
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return h
}

func main() {
	fmt.Println(hIndex([]int{1}))
	fmt.Println(hIndex([]int{1, 0}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
