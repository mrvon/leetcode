package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	peakIndex := 1
	curr := A[peakIndex-1]
	for peakIndex < len(A) {
		if curr > A[peakIndex] {
			return peakIndex - 1
		}
		curr = A[peakIndex]
		peakIndex++
	}
	return peakIndex
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
}
