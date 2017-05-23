package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	new := make([]int, len(nums))

	copy(new, nums)
	sort.Ints(new)

	i := 0
	j := len(nums) - 1

	for i <= j {
		if nums[i] == new[i] {
			i++
		} else {
			break
		}
	}

	for i <= j {
		if nums[j] == new[j] {
			j--
		} else {
			break
		}
	}

	if j >= i {
		return j - i + 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 8, 9, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{2}))
	fmt.Println(findUnsortedSubarray([]int{}))
}
