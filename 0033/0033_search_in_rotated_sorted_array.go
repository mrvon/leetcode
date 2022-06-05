// O(logN), Revised binary search
package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if nums[left] <= nums[mid] {
			// nums[left...mid] is ascending array
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// nums[mid...right] is ascending array
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 8, 0, 1, 2}

	fmt.Println(search(arr, 4))
	fmt.Println(search(arr, 5))
	fmt.Println(search(arr, 6))
	fmt.Println(search(arr, 7))
	fmt.Println(search(arr, 8))
	fmt.Println(search(arr, 9))
	fmt.Println(search(arr, 0))
	fmt.Println(search(arr, 1))
	fmt.Println(search(arr, 2))
}
