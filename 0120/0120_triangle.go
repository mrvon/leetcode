package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func wrapper(triangle [][]int, i int, j int) int {
	if i < 0 || i >= len(triangle) {
		return math.MaxInt32
	}

	if j < 0 || j >= len(triangle[i]) {
		return math.MaxInt32
	}

	return triangle[i][j]
}

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			p1 := wrapper(triangle, i-1, j-1)
			p2 := wrapper(triangle, i-1, j)

			triangle[i][j] = triangle[i][j] + min(p1, p2)
		}
	}

	m := math.MaxInt32
	i := len(triangle) - 1
	for j := 0; j <= i; j++ {
		if triangle[i][j] < m {
			m = triangle[i][j]
		}
	}
	return m
}

func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))

	triangle = [][]int{
		[]int{-1},
		[]int{3, 2},
		[]int{-3, 1, -1},
	}
	fmt.Println(minimumTotal(triangle))
}
