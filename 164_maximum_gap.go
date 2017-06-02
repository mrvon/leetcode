package main

import (
	"fmt"
	"sort"
)

// O(N*LogN)
func maximumGap(nums []int) int {
	sort.Ints(nums)

	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > max {
			max = nums[i] - nums[i-1]
		}
	}
	return max
}

// O(N)
// TODO

func main() {
	fmt.Println(maximumGap([]int{}))
	fmt.Println(maximumGap([]int{1, 200, 2, 100, 3}))
}
