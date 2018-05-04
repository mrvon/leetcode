package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	m := 0 // max
	c := 0 // counter
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// reset
			if c > m {
				m = c
			}
			c = 0
		} else {
			// count
			c++
		}
	}
	if c > m {
		return c
	} else {
		return m
	}
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
