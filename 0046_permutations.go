/*
This algorithm is work for the following two problem

https://leetcode.com/problems/permutations/?tab=Description
https://leetcode.com/problems/permutations-ii/?tab=Description

Discription:

Given a collection of numbers that might contain duplicates, return all possible
unique permutations.
*/
package main

import (
	"fmt"
	"sort"
)

func next_permutation(arr []int) bool {
	n := len(arr) - 1

	k := -1
	for i := n - 1; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			k = i
			break
		}
	}

	// last permutation
	if k == -1 {
		return false
	}

	l := n
	for i := n; i > k; i-- {
		if arr[k] < arr[i] {
			l = i
			break
		}
	}

	arr[k], arr[l] = arr[l], arr[k]

	i := k + 1
	j := n
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return true
}

func permute(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for {
		perm := make([]int, len(nums))
		copy(perm, nums)
		result = append(result, perm)

		if !next_permutation(nums) {
			break
		}
	}

	return result
}

func main() {
	// for problem permutations
	fmt.Println(permute([]int{}))
	fmt.Println(permute([]int{1, 2, 3}))

	// for problem permutations-ii
	fmt.Println(permute([]int{1, 1, 1}))
}
