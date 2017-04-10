/*
Dynamic programming solution
*/
package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// O(n), Accepted
// https://discuss.leetcode.com/topic/4417/possibly-simplest-solution-with-o-n-time-complexity
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// store the result that is the max we have found so far
	r := nums[0]

	// imax/imin stores the max/min product of
	// subarray the ends with the current nums A[i]
	imax := r
	imin := r
	for i := 1; i < len(nums); i++ {
		// multiplied by a negative makes big number smaller, small number
		// bigger, so we redefine the extremums by swapping them
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		// max/min product for the current number is either the current number
		// itself or the max/min by the previous number times the current one
		imax = max(nums[i], imax*nums[i])
		imin = min(nums[i], imin*nums[i])

		// the newly computed max value is a candidate for our global result
		r = max(r, imax)
	}

	return r
}

// O(n^2), Accepted
func maxProduct2(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	optimal := make([]int, n)
	optimal[0] = nums[0]

	for i := 1; i < n; i++ {
		product := nums[i]
		max_product := nums[i]
		for j := i - 1; j >= 0; j-- {
			product *= nums[j]
			max_product = max(max_product, product)
		}
		optimal[i] = max(optimal[i-1], max_product)
	}

	return optimal[n-1]
}

func main() {
	fmt.Println(maxProduct([]int{}))
	fmt.Println(maxProduct([]int{2}))
	fmt.Println(maxProduct([]int{2, 3}))
	fmt.Println(maxProduct([]int{2, 3, -2}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{2, 3, -2, -4}))
	fmt.Println(maxProduct([]int{2, 3, -2, 7}))

	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, 0}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, -2, -1}))
}
