package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	// The skyline viewed from left or right
	lr := make([]int, len(grid))
	// The skyline viewed from top or bottom
	tb := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			lr[i] = max(lr[i], grid[i][j])
			tb[j] = max(tb[j], grid[i][j])
		}
	}

	// Max we can increase
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += (min(lr[i], tb[j]) - grid[i][j])
		}
	}
	return count
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}))
}
