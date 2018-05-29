package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	count := 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			diff := nums[j] - nums[i]
			if diff == k {
				count++
			} else if diff > k {
				break
			}
		}
	}

	return count
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}
