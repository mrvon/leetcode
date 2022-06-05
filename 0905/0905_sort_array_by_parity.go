// Swap approach, inplace algorithm.
package main

import "fmt"

func sortArrayByParity(A []int) []int {
	low := 0
	high := len(A) - 1
	for {
		for low < high && A[low]%2 == 0 {
			low++
		}
		for low < high && A[high]%2 == 1 {
			high--
		}
		if low >= high {
			break
		}
		A[low], A[high] = A[high], A[low]
		low++
		high--
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{1, 3, 2, 4}))
	fmt.Println(sortArrayByParity([]int{2, 4, 1, 3}))
}
