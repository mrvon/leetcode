package main

import "fmt"

func transpose(A [][]int) [][]int {
	M := len(A)
	N := len(A[0])
	T := [][]int{}
	for i := 0; i < N; i++ {
		T = append(T, make([]int, M))
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			T[j][i] = A[i][j]
		}
	}
	return T
}

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
