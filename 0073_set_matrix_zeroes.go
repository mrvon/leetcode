// O(1) Space solution, inspired by Leetcode discuss.
package main

func setZeroes(matrix [][]int) {
	M := len(matrix)
	N := len(matrix[0])

	first_row := false
	first_column := false

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					first_row = true
				}
				if j == 0 {
					first_column = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if first_row {
		for j := 0; j < N; j++ {
			matrix[0][j] = 0
		}
	}

	if first_column {
		for i := 0; i < M; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
	}
	setZeroes(matrix)
}
