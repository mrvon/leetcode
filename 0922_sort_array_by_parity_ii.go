package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	evenIndex := 0
	oddIndex := 1
	length := len(A)
	for {
		for evenIndex < length && A[evenIndex]%2 == 0 { // Even
			evenIndex += 2
		}
		if evenIndex >= length {
			break
		}
		for oddIndex < length && A[oddIndex]%2 == 1 { // Odd
			oddIndex += 2
		}
		if oddIndex >= length {
			break
		}
		A[evenIndex], A[oddIndex] = A[oddIndex], A[evenIndex]
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
