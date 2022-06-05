/*
find every house, we find the min_distance away from heaters,
using binary search.

the max(min_distance) in all house is the answer we need.
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func binary_search(heaters []int, house int) int {
	left := 0
	right := len(heaters) - 1
	mid := 0
	for left <= right {
		mid = left + (right-left)/2
		if house == heaters[mid] {
			// find it
			return mid
		} else if house < heaters[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return mid
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)

	max_distance := 0

	for i := 0; i < len(houses); i++ {
		house := houses[i]
		min_distance := math.MaxInt32

		mid := binary_search(heaters, house)

		if house == heaters[mid] {
			min_distance = min(min_distance, 0)
		} else if house > heaters[mid] {
			min_distance = min(min_distance, house-heaters[mid])
			if mid+1 < len(heaters) {
				min_distance = min(min_distance, heaters[mid+1]-house)
			}
		} else {
			min_distance = min(min_distance, heaters[mid]-house)
			if mid-1 >= 0 {
				min_distance = min(min_distance, house-heaters[mid-1])
			}
		}

		max_distance = max(max_distance, min_distance)
	}

	return max_distance
}

func main() {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(findRadius([]int{1, 2, 3}, []int{1, 3}))
	fmt.Println(findRadius([]int{1, 2, 3}, []int{3}))
}
