package main

import (
	"fmt"
	"math"
)

// O(N)
func minSubArrayLen(s int, nums []int) int {
	min := func(x, y int) int {
		if x <= y {
			return x
		} else {
			return y
		}
	}

	n := len(nums)

	sum := 0
	start := 0
	min_len := math.MaxInt32

	for end := 0; end < n; end++ {
		sum += nums[end]

		for sum >= s {
			min_len = min(min_len, end-start+1)
			// start position move forward
			sum -= nums[start]
			start++
		}
	}

	if min_len == math.MaxInt32 {
		return 0
	} else {
		return min_len
	}
}

// O(N^2) dynamic programming
func minSubArrayLen2(s int, nums []int) int {
	min := func(x, y int) int {
		if x <= y {
			return x
		} else {
			return y
		}
	}

	n := len(nums)
	if n == 0 {
		return 0
	}

	optimal := make([]int, n)

	for i := 0; i < n; i++ {
		sum := 0
		j := 0

		for j = i; j >= 0; j-- {
			sum += nums[j]
			if sum >= s {
				if i == 0 || optimal[i-1] == 0 {
					optimal[i] = i - j + 1
				} else {
					optimal[i] = min(optimal[i-1], i-j+1)
				}
				break
			}
		}

		if j < 0 {
			// no solution
			optimal[i] = 0
		}
	}

	return optimal[n-1]
}

func main() {
	fmt.Println(minSubArrayLen(100, []int{}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
