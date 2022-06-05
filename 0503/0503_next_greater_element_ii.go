// TODO - too slow but accept
package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		g := nextGreater(nums, i)
		if g != -1 {
			result = append(result, g)
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func nextGreater(nums []int, start int) int {
	for i := start + 1; i < len(nums); i++ {
		if nums[i] > nums[start] {
			return nums[i]
		}
	}
	for i := 0; i < start; i++ {
		if nums[i] > nums[start] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
