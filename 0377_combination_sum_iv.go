package main

import "fmt"

/* original recursive solution, Time limit exceeded.
func combinationSum4(nums []int, target int) int {
	if target < 0 {
		return 0
	} else if target == 0 {
		return 1
	}

	// target > 0

	count := 0
	for i := 0; i < len(nums); i++ {
		count += combinationSum4(nums, target-nums[i])
	}
	return count
}
*/

// translate recursive solution to dynamic programming

// dynamic programming solution
func combinationSum4(nums []int, target int) int {
	sum := make([]int, target+1)

	sum[0] = 1

	for i := 1; i <= target; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				count += sum[i-nums[j]]
			}
		}
		sum[i] = count
	}

	return sum[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4

	fmt.Println(combinationSum4(nums, target))
}
