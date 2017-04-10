package main

import "fmt"

func findMaxLength(nums []int) int {
	// max length
	max := 0
	// sum -> index mapping
	sum_list := make(map[int]int)
	sum := 0
	sum_list[sum] = -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 { // we treats 0 as -1 for convenient.
			sum += -1
		} else {
			sum += 1
		}
		if j, b := sum_list[sum]; b {
			// we have got this sum before, it's means that from j to i in the
			// array, the sum is not change, and have equal number of 0 and 1.
			length := i - j
			if length > max {
				max = length
			}
		} else {
			// first time we get this sum, record its index.
			sum_list[sum] = i
		}
	}

	return max
}

func main() {
	fmt.Println(findMaxLength([]int{1, 1, 0, 0}))
}
