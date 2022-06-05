package main

import "fmt"

/*
For every subarray-sum sum(nums[j:k]) = sum(nums[0:k]) - sum(nums[0:j]).
So we can keep prefix sum's count in `prefixSum`, and note that empty prefix
subarray's sum is 0.

In the loop, we consider every prefix nums[0:i], The problem is that find how
many subarray nums[s:i] end at i whose sum is k, We can solve by find how many
prefix subarray whose sum is `sum-k`.
*/
func subarraySum(nums []int, k int) int {
	// prefix sum -> count
	prefixSum := make(map[int]int)
	// empty prefix's sum is 0
	prefixSum[0]++
	// current prefix sum
	sum := 0
	// subarray sum is k
	k_count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// how many subarray end at i whose sum is k
		k_count += prefixSum[sum-k]
		// one more prefix subarray whose sum is `sum`
		prefixSum[sum]++
	}
	return k_count
}

/*
Another O(n^2) DP solution, Accepted
*/
func subarraySumN2(nums []int, k int) int {
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
