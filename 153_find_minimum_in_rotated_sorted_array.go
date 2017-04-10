package main

import (
	"fmt"
	"math"
)

func findMin(nums []int) int {
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
	fmt.Println(findMin([]int{1, 2, 4, 5, 6, 7, 0}))
	fmt.Println(findMin([]int{2, 4, 5, 6, 7, 0, 1}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{2, 4, 5, 6, 7, 0, 1}))
	fmt.Println(findMin([]int{1, 2, 4, 5, 6, 7, 0}))
	fmt.Println(findMin([]int{0, 1, 2, 4, 5, 6, 7}))
}
