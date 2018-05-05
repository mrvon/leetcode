package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var buffer []int

	// sorting
	sort.Ints(candidates)

	return sum2(candidates, target, result, buffer)
}

func sum2(candidates []int, target int, result [][]int, buffer []int) [][]int {
	if target < 0 {
		return result
	} else if target == 0 {
		// find it
		result = append(result, make([]int, len(buffer)))
		copy(result[len(result)-1], buffer)
		return result
	}
	// target > 0

	if len(candidates) == 0 {
		return result
	}

	// how many same element
	count := 1
	for i := 1; i < len(candidates); i++ {
		if candidates[i] == candidates[0] {
			count++
		}
	}

	for i := 0; i <= count; i++ {
		result = sum2(candidates[count:len(candidates)], target-i*candidates[0], result, buffer)
		buffer = append(buffer, candidates[0])
	}
	buffer = buffer[0 : len(buffer)-count]

	return result
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
