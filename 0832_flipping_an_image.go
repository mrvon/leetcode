package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		l := 0
		r := len(A[i]) - 1
		for l <= r {
			if l == r {
				// Invert
				if A[i][l] == 0 {
					A[i][l] = 1
				} else {
					A[i][l] = 0
				}
			} else {
				// Swap
				t := A[i][l]
				A[i][l] = A[i][r]
				A[i][r] = t
				// Invert
				if A[i][l] == 0 {
					A[i][l] = 1
				} else {
					A[i][l] = 0
				}
				if A[i][r] == 0 {
					A[i][r] = 1
				} else {
					A[i][r] = 0
				}
			}
			l++
			r--
		}
	}
	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 1}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
