package main

import "fmt"

func subarraySumN(nums []int, k int) int {
	// TODO
	return 0
}

// O(n^2)
func subarraySum(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	optimal := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			if nums[i] == k {
				optimal[i] = 1
			} else {
				optimal[i] = 0
			}
		} else {
			count := 0
			sum := 0
			for j := i; j >= 0; j-- {
				sum += nums[j]
				if sum == k {
					count++
				}
			}
			optimal[i] = optimal[i-1] + count
		}
	}
	return optimal[n-1]
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 5, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 5, 1, 1}, 6))
}
