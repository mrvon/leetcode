/*
flood-fill solution
*/
package main

import "fmt"

var direction = []int{
	-1, 0,
	1, 0,
	0, -1,
	0, 1,
}

func flood_fill(grid [][]byte, i int, j int) {
	m := len(grid)
	n := len(grid[0])

	// marked for visited
	grid[i][j] = '0'

	for k := 0; k < len(direction); k += 2 {
		x := i + direction[k]
		y := j + direction[k+1]

		if x < 0 || x >= m || y < 0 || y >= n {
			continue
		}

		if grid[x][y] == '0' {
			continue
		}

		flood_fill(grid, x, y)
	}
}

func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				flood_fill(grid, i, j)
				count++
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	fmt.Println(numIslands(grid))

	grid = [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(grid))
}
