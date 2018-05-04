package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return gen(nums, [][]int{}, []int{})
}

func gen(nums []int, sets [][]int, set []int) [][]int {
	if len(nums) == 0 {
		ns := make([]int, len(set))
		copy(ns, set)
		sets = append(sets, ns)
	} else {
		// handle duplicate
		c := 1
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[0] {
				c++
			} else {
				break
			}
		}

		// without nums[0]
		sets = gen(nums[c:], sets, set)

		for i := 0; i < c; i++ {
			set = append(set, nums[0])
			sets = gen(nums[c:], sets, set)
		}
	}
	return sets
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
