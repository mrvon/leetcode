/*
Put each number in its right place.
For example, When we find 5, then swap it with A[4].
*/
package main

import "fmt"

func firstMissingPositive(nums []int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= l && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return l + 1
}

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(1, firstMissingPositive([]int{}))
	assert(3, firstMissingPositive([]int{1, 2, 0}))
	assert(3, firstMissingPositive([]int{1, 2, 4}))
	assert(3, firstMissingPositive([]int{1, 2}))
	assert(2, firstMissingPositive([]int{3, 4, -1, 1}))
}
