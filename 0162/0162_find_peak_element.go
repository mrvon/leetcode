package main

import (
	"fmt"
	"math"
)

// O(n) solution
func n_findPeakElement(nums []int) int {
	get := func(nums []int, index int) int {
		if index == -1 || index == len(nums) {
			return math.MinInt32
		} else {
			return nums[index]
		}
	}

	for i := 0; i < len(nums); i++ {
		l := get(nums, i-1)
		r := get(nums, i+1)
		m := get(nums, i)
		if m >= l && m >= r {
			return i
		}
	}

	return -1
}

// O(logN)
func findPeakElement(nums []int) int {
	get := func(nums []int, index int) int {
		if index == -1 || index == len(nums) {
			return math.MinInt32
		} else {
			return nums[index]
		}
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		l := get(nums, mid-1)
		r := get(nums, mid+1)
		m := get(nums, mid)

		if m >= l && m >= r {
			return mid
		}

		if l <= m {
			left = mid + 1
		} else if r <= m {
			right = mid - 1
		} else { // m <= l && m <= r
			left = mid + 1
			// or right = mid - 1, is ok.
		}
	}

	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{math.MinInt32}))
	fmt.Println(findPeakElement([]int{1}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}
