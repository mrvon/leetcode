/*
Dynamic programming

Base on the thought in 174_dungeon_game_dfs, easy translate to a dp solution.
*/
package main

import (
	"fmt"
	"math"
)

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

// Calculate cur cell least hp
func min_hp(previous_hp int, cur_cell int) int {
	previous_hp -= cur_cell
	if previous_hp <= 0 {
		previous_hp = 1
	}
	return previous_hp
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	optimal := [][]int{}
	for i := 0; i < m; i++ {
		optimal = append(optimal, make([]int, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			optimal[i][j] = math.MaxInt32
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				// Our knight should have at least 1 hp at last.
				optimal[i][j] = min_hp(1, dungeon[i][j])
			} else {
				if i+1 < m {
					optimal[i][j] = min(optimal[i][j], min_hp(optimal[i+1][j], dungeon[i][j]))
				}
				if j+1 < n {
					optimal[i][j] = min(optimal[i][j], min_hp(optimal[i][j+1], dungeon[i][j]))
				}
			}
		}
	}

	return optimal[0][0]
}

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(dungeon))
}
