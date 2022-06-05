package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func countSquares(matrix [][]int) int {
	M := len(matrix)
	N := len(matrix[0])
	dp := [][]int{}
	for i := 0; i < M; i++ {
		dp = append(dp, make([]int, N))
	}
	sum := 0
	for i := 0; i < M; i++ {
		if matrix[i][0] == 1 {
			dp[i][0] = 1
			sum += 1
		}
	}
	for j := 1; j < N; j++ {
		if matrix[0][j] == 1 {
			dp[0][j] = 1
			sum += 1
		}
	}
	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				sum += dp[i][j]
			}
		}
	}
	return sum
}

func main() {
	matrix := [][]int{
		// {0, 1, 1, 1},
		// {1, 1, 1, 1},
		// {0, 1, 1, 1},

		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(countSquares(matrix))
}
