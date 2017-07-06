package main

import (
	"fmt"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums) - 1
	return max(nums[0]*nums[1]*nums[n], nums[n-2]*nums[n-1]*nums[n])
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-2, -1, 3, 4}))
	fmt.Println(maximumProduct([]int{-3, -2, -1, 4}))
	fmt.Println(maximumProduct([]int{-4, -3, -2, -1}))
	fmt.Println(maximumProduct([]int{-3, -2, -1}))
}
