package main

import "fmt"

func insert(nums []int, n int) []int {
	for i := 0; i < len(nums); i++ {
		if n < nums[i] {
			nums = append(nums, 0)
			copy(nums[i+1:], nums[i:len(nums)-1])
			nums[i] = n
			return nums
		}
	}
	return append(nums, n)
}

func delete(nums []int, n int) []int {
	for i := 0; i < len(nums); i++ {
		if n == nums[i] {
			copy(nums[i:len(nums)-1], nums[i+1:])
			return nums[:len(nums)-1]
		}
	}
	return nums
}

func medianOdd(nums []int) float64 {
	n := len(nums)
	return float64(nums[n/2])
}

func medianEven(nums []int) float64 {
	n := len(nums)
	return float64(nums[n/2-1]+nums[n/2]) / float64(2)
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var mf func(nums []int) float64

	if k%2 == 0 {
		mf = medianEven
	} else {
		mf = medianOdd
	}

	sorted := []int{}
	results := []float64{}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			sorted = insert(sorted, nums[i])
		} else {
			sorted = insert(sorted, nums[i])
			if i > k-1 {
				sorted = delete(sorted, nums[i-k])
			}
			results = append(results, mf(sorted))
		}
	}
	return results
}

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 4))
}
