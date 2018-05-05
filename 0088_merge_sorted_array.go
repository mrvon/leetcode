package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)

	j := 0
	k := 0
	for i := 0; i < len(nums3); i++ {
		if j >= m {
			nums3[i] = nums2[k]
			k++
		} else if k >= n {
			nums3[i] = nums1[j]
			j++
		} else {
			if nums1[j] <= nums2[k] {
				nums3[i] = nums1[j]
				j++
			} else {
				nums3[i] = nums2[k]
				k++
			}
		}
	}

	for i := 0; i < len(nums3); i++ {
		nums1[i] = nums3[i]
	}
}

func assert(b bool) {
	if !b {
		panic("assert failed.")
	}
}

func main() {
	nums1 := []int{1, 5, 6, 6, 0, 0, 0, 0}
	nums2 := []int{2, 3, 7, 8}
	merge(nums1, 4, nums2, 4)
	fmt.Println(nums1)

	// nums1 := []int{1}
	// nums2 := []int{}
	// merge(nums1, 1, nums2, 0)
	// fmt.Println(nums1)
}
