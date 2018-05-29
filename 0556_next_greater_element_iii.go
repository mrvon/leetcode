package main

import (
	"fmt"
	"math"
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

func toArr(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func toInt(arr []int) int {
	n := 0
	for i := 0; i < len(arr); i++ {
		n *= 10
		n += arr[i]
	}
	return n
}

func isExceed(arr []int) bool {
	max := toArr(math.MaxInt32)

	if len(arr) > len(max) {
		return true
	} else if len(arr) < len(max) {
		return false
	} else {
		for i := 0; i < len(arr); i++ {
			if arr[i] > max[i] {
				return true
			} else if arr[i] < max[i] {
				return false
			}
		}
	}
	return false
}

func nextGreaterElement(n int) int {
	arr := toArr(n)
	if !next_permutation(arr) {
		return -1
	}
	if isExceed(arr) {
		return -1
	}
	return toInt(arr)
}

func main() {
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(math.MaxInt32))
}
