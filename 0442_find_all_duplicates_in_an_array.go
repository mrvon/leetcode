// O(n) runtime complexity, O(1) space complexity
// Idea: place element n in nums[n-1]
package main

import "fmt"

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); {
		n := nums[i]
		if nums[n-1] != n {
			nums[n-1], nums[i] = nums[i], nums[n-1]
		} else {
			i++
		}
	}

	var duplicates []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			duplicates = append(duplicates, nums[i])
		}
	}
	return duplicates
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// Output:
	// [2,3]
}
