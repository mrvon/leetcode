package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l := len(nums)

	if l == 0 {
		// error
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		// error
		return 0
	}

	// otherwise

	m := (l - 1) / 2
	n := 0

	if nums[m-1] == nums[m] {
		n = m
		m = m - 1
	} else if nums[m+1] == nums[m] {
		n = m + 1
	} else {
		return nums[m]
	}

	if len(nums[:m])%2 != 0 {
		return singleNonDuplicate(nums[:m])
	} else if len(nums[n+1:])%2 != 0 {
		return singleNonDuplicate(nums[n+1:])
	} else {
		// error
		return 0
	}
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
