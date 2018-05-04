// O(n) runtime complexity, O(1) space complexity
// Idea: place element n in nums[n-1]
package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		n := nums[i]
		if nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], nums[i]
		} else {
			i++
		}
	}

	var disappeared []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			disappeared = append(disappeared, i+1)
		}
	}
	return disappeared
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
