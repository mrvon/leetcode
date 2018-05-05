/*
Though we can slightly modified 85_maximal_rectangle.go can pass this problem.
But it has another simply approach.

Dynamic programming
*/
package main

import (
	"fmt"
)

func min(x int, y int, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= z {
		return y
	} else {
		return z
	}
}

// optimal(i,j) represents the side length of the maximum square whose bottom
// right corner is the cell with index(i,j) in the original matrix.
func maximalSquare(matrix [][]byte) int {
	optimal := [][]int{}
	max := 0
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		optimal = append(optimal, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				optimal[i][j] = 0
			} else {
				if i == 0 || j == 0 {
					optimal[i][j] = 1
				} else {
					optimal[i][j] = min(
						optimal[i-1][j],
						optimal[i][j-1],
						optimal[i-1][j-1],
					) + 1
				}
			}
			if optimal[i][j] > max {
				max = optimal[i][j]
			}
		}
	}
	return max * max
}

func main() {
	matrix := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	fmt.Println(maximalSquare(matrix))

	matrix = [][]byte{
		[]byte(""),
	}
	fmt.Println(maximalSquare(matrix))

	matrix = [][]byte{
		[]byte("1"),
	}
	fmt.Println(maximalSquare(matrix))
}
