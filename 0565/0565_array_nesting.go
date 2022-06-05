package main

import "fmt"

func arrayNesting(nums []int) int {
	unique := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		unique[nums[i]] = true
	}

	max := 0

	for len(unique) > 0 {
		index := 0
		// Get One
		for index = range unique {
			break
		}
		count := 0
		for {
			if !unique[index] {
				break
			}
			delete(unique, index)
			count++
			index = nums[index]
		}
		if count > max {
			max = count
		}
	}

	return max
}

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}
