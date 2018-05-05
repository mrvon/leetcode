package main

import "fmt"

func subsets(nums []int) [][]int {
	return gen(nums, [][]int{}, []int{})
}

func gen(nums []int, sets [][]int, set []int) [][]int {
	if len(nums) == 0 {
		ns := make([]int, len(set))
		copy(ns, set)
		sets = append(sets, ns)
	} else {
		sets = gen(nums[1:], sets, set)
		set = append(set, nums[0])
		sets = gen(nums[1:], sets, set)
	}
	return sets
}

func main() {
	fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{1}))
	fmt.Println(subsets([]int{1, 2}))
	fmt.Println(subsets([]int{1, 2, 3}))
}
