// TODO
package main

import (
	"fmt"
	"math"
)

func removeDuplicate(nums []int) []int {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[0] != nums[i] {
			break
		}
	}
	nums = nums[:i+1]

	new := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] != new[len(new)-1] {
			new = append(new, nums[i])
		}
	}

	return new
}

// code copy from problem 153.
func findMin(nums []int) int {
	nums = removeDuplicate(nums)

	left := 0
	right := len(nums) - 1
	min := math.MaxInt32

	for left <= right {
		mid := left + (right-left)/2

		if nums[left] <= nums[mid] {
			// nums[left...mid] is ascending array
			if nums[left] < min {
				min = nums[left]
			}
			left = mid + 1
		} else {
			// nums[mid...right] is ascending array
			if nums[mid] < min {
				min = nums[mid]
			}
			right = mid - 1
		}
	}

	return min
}

func main() {
	fmt.Println(findMin([]int{1, 10, 10, 10, 10, 10, 12}))
	fmt.Println(findMin([]int{10, 12, 1, 10, 10, 10, 10}))
}
