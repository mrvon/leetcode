package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)

	for _, n := range nums {
		if set[n] > 0 {
			return true
		} else {
			set[n]++
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 1}))
}
