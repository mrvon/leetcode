package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var buffer []int
	return sum(candidates, target, result, buffer)
}

func sum(candidates []int, target int, result [][]int, buffer []int) [][]int {
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

	for i := 0; i <= target/candidates[0]; i++ {
		for j := 0; j < i; j++ {
			buffer = append(buffer, candidates[0])
		}
		result = sum(
			candidates[1:len(candidates)],
			target-i*candidates[0],
			result,
			buffer,
		)
		buffer = buffer[0 : len(buffer)-i]
	}

	return result
}

func main() {
	c := []int{2, 3, 6, 7}
	t := 7
	fmt.Println(combinationSum(c, t))
}
