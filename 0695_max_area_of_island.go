package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxAreaOfIsland(grid [][]int) int {
	visited := [][]bool{}
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}

	var maxArea func(x int, y int) int
	maxArea = func(x int, y int) int {
		if x < 0 || x >= len(grid) {
			return 0
		}
		if y < 0 || y >= len(grid[x]) {
			return 0
		}

		if grid[x][y] == 0 {
			return 0
		}

		if visited[x][y] {
			return 0
		}

		visited[x][y] = true

		return 1 + maxArea(x-1, y) + maxArea(x+1, y) + maxArea(x, y-1) + maxArea(x, y+1)
	}

	m := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			m = max(m, maxArea(i, j))
		}
	}
	return m
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}}))
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
	fmt.Println(maxAreaOfIsland([][]int{{1, 1}, {1, 0}}))
}
