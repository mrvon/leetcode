package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		j := i
		k := 0
		for {
			j++
			k++
			if j >= m || k >= n {
				break
			}
			if matrix[j][k] != matrix[j-1][k-1] {
				return false
			}
		}
	}

	for i := 1; i < n; i++ {
		j := 0
		k := i
		for {
			j++
			k++
			if j >= m || k >= n {
				break
			}
			if matrix[j][k] != matrix[j-1][k-1] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}
