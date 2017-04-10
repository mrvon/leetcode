package main

import "fmt"

func search(nums []int, target int) bool {
	j := len(nums)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[0] {
			j--
		} else {
			break
		}
	}

	news := []int{}

	for i := 0; i < j; i++ {
		if i == 0 {
			news = append(news, nums[i])
		} else {
			if news[len(news)-1] != nums[i] {
				news = append(news, nums[i])
			}
		}
	}

	return searchWithoutDuplicate(news, target) != -1
}

// code from 33_search_in_rotated_sorted_array.go
func searchWithoutDuplicate(nums []int, target int) int {
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
	fmt.Println(search([]int{2, 4, 5, 6, 7, 0, 0, 1, 2, 2}, 7))
	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{1, 1, 3}, 3))
}
