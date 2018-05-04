package main

import (
	"fmt"
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)

	min_count := 0
	min_value := 0
	max_count := 0
	max_value := 0
	result := 0

	for i := 0; i < len(nums); i++ {
		if min_count == 0 {
			min_value = nums[i]
			min_count++
		} else {
			if min_value == nums[i] {
				min_count++
			} else {
				if nums[i]-min_value == 1 {
					max_value = nums[i]
					max_count++
					if min_count+max_count > result {
						result = min_count + max_count
					}
				} else {
					min_value = max_value
					min_count = max_count
					max_count = 0
					i--
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(findLHS([]int{1}))
	fmt.Println(findLHS([]int{2, 1}))
	fmt.Println(findLHS([]int{1, 3, 2}))
	fmt.Println(findLHS([]int{1, 3, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 3, 2, 5, 2, 3, 7, 2, 2, 3}))
}
