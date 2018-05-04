package main

import "fmt"

func set(s int, nums []int, m map[int]int) {
	if len(nums) == 0 {
		m[s]++
		return
	}

	set(s+nums[0], nums[1:len(nums)], m)
	set(s-nums[0], nums[1:len(nums)], m)
}

func get(s int, nums []int, m map[int]int) int {
	if len(nums) == 0 {
		return m[-s]
	}

	return get(s+nums[0], nums[1:len(nums)], m) +
		get(s-nums[0], nums[1:len(nums)], m)
}

func findTargetSumWays(nums []int, S int) int {
	m := make(map[int]int)

	// use mid point the split nums into two part
	mid := len(nums) / 2

	// left
	set(0, nums[0:mid], m)

	// right
	return get(S, nums[mid:len(nums)], m)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
