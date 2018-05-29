package main

import "fmt"

var candidates []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var buffer []int
	return sum3(0, k, n, result, buffer)
}

func sum3(
	i int, // consider element candidates[i]
	k int, // remain k element to select
	n int, // sum remainder
	result [][]int, buffer []int) [][]int {

	if n < 0 {
		return result
	} else if n == 0 {
		// find a solution
		if k == 0 { // is it exact count of element?
			result = append(result, make([]int, len(buffer)))
			copy(result[len(result)-1], buffer)
		}
		return result
	}
	// n > 0

	if i >= len(candidates) { // haven't candidates to select
		return result
	}

	// select candidates[i]
	buffer = append(buffer, candidates[i])
	result = sum3(i+1, k-1, n-candidates[i], result, buffer)
	buffer = buffer[0 : len(buffer)-1]

	// discard candidates[i]
	result = sum3(i+1, k, n, result, buffer)

	return result
}

func main() {
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(3, 7))
}
