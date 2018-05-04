package main

import "fmt"

/*
O(N) Time

We iterate through the input array exactly once, keeping track of the running
sum mod k of the elements in the process. If we find that a running sum value at
index j has been previously seen before in some earlier index i in the array,
then we know that the sub-array (i,j] contains a desired sum.
*/
func checkSubarraySum(nums []int, k int) bool {
	if k == 0 { // special case
		if len(nums) <= 1 {
			return false
		} else {
			if nums[0] == 0 && nums[1] == 0 {
				return true
			} else {
				return false
			}
		}
	}
	// k is not 0

	sum_map := make(map[int]int)
	sum_map[0] = -1 // fake index -1 for convenient

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= k
		if index, is_exist := sum_map[sum]; is_exist {
			if i-index >= 2 { // cause "at least two element"
				return true
			}
		} else {
			sum_map[sum] = i
		}
	}

	return false
}

// O(N^2), straightforward recursive solution, 200ms, too slow but accept.
func __checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	} else if len(nums) == 2 {
		if k == 0 {
			return (nums[0] + nums[1]) == 0
		} else {
			return (nums[0]+nums[1])%k == 0
		}
	} else {
		// without last element in nums
		is_exist := checkSubarraySum(nums[0:len(nums)-1], k)
		if is_exist {
			return true
		}
		// with last element in nums
		sum := nums[len(nums)-1]
		for i := len(nums) - 2; i >= 0; i-- {
			sum += nums[i]
			if k == 0 {
				if sum == 0 {
					return true
				}
			} else {
				if sum%k == 0 {
					return true
				}
			}
		}
		return false
	}
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(true, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	assert(true, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	assert(false, checkSubarraySum([]int{23, 2}, 6))
	assert(false, checkSubarraySum([]int{23, 2, 6, 4, 7}, 0))
	assert(true, checkSubarraySum([]int{0, 0}, 0))
}
