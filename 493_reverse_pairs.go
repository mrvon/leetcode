package main

import (
	"fmt"
)

func mergeSort(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return 0
	}

	mid := n / 2

	total := mergeSort(nums[:mid]) + mergeSort(nums[mid:])

	// nums[:mid] and nums[mid:] is sorted
	j := mid
	for i := 0; i < mid; i++ {
		for j < n {
			if nums[i] > 2*nums[j] {
				j++
			} else {
				break
			}
		}
		count := j - mid
		total += count
	}

	// merge-sort entirely nums[:]
	news := make([]int, n)
	for i := 0; i < n; i++ {
		news[i] = nums[i]
	}
	i := 0
	j = mid
	k := 0
	for i < mid || j < n {
		if i >= mid {
			nums[k] = news[j]
			k++
			j++
		} else if j >= n {
			nums[k] = news[i]
			k++
			i++
		} else {
			if news[i] <= news[j] {
				nums[k] = news[i]
				k++
				i++
			} else {
				nums[k] = news[j]
				k++
				j++
			}
		}
	}

	return total
}

func reversePairs(nums []int) int {
	return mergeSort(nums)
}

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs([]int{5, 4, 3, 2, 1}))
}
