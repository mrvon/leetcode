package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	var result []int
	for i := 0; i < len(findNums); i++ {
		n := findNums[i]
		// find n in the num2
		j := 0
		for ; j < len(nums); j++ {
			if n == nums[j] {
				break
			}
		}
		// find next great number in the num2
		is_find := false
		for j = j + 1; j < len(nums); j++ {
			if nums[j] > n {
				is_find = true
				break
			}
		}
		if is_find {
			result = append(result, nums[j])
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
