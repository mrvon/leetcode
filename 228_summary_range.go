package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ranges := []string{}
	index := -1

	for i := 0; i <= len(nums); i++ {
		if index == -1 {
			index = i
		} else {
			if i == len(nums) || nums[i] != nums[i-1]+1 {
				if index == i-1 {
					// only one
					ranges = append(ranges, strconv.Itoa(nums[index]))
				} else {
					ranges = append(ranges,
						strconv.Itoa(nums[index])+"->"+strconv.Itoa(nums[i-1]))
				}
				index = i
			}
		}
	}

	return ranges
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7, 9, 10}))
}
