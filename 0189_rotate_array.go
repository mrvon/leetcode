package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)

	out := make([]int, k)

	j := 0
	for i := len(nums) - k; i < len(nums); i++ {
		out[j] = nums[i]
		j++
	}

	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	// copy back
	for i := 0; i < len(out); i++ {
		nums[i] = out[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 5
	rotate(nums, k)
	fmt.Println(nums)
}
