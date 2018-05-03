package main

import (
	"fmt"
)

// Naive, O(n^2)
func twoSumNaive(nums []int, target int) []int {
	r := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				r = append(r, i)
				r = append(r, j)
			}
		}
	}
	return r
}

// Hash table
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		if j, has := m[target-n]; has {
			return []int{i, j}
		}
		m[n] = i
	}

	// error occurs
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
