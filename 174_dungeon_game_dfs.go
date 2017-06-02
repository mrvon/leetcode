/*
*Graph* Depth-first search

Time limit exceeded
*/
package main

import "fmt"

type P struct {
	x int
	y int
}

// Calculate cur cell least hp
func min_hp(previous_hp int, cur_cell int) int {
	previous_hp -= cur_cell
	if previous_hp <= 0 {
		previous_hp = 1
	}
	return previous_hp
}

func reverse_dfs(dungeon [][]int, visited map[P]bool, hpmap map[P]int, start P, cur_hp int) {
	if start.x < 0 || start.y < 0 {
		return
	}
	visited[start] = true
	min_hp := min_hp(cur_hp, dungeon[start.x][start.y])
	if value, has := hpmap[start]; !has || value > min_hp {
		hpmap[start] = min_hp
	}
	left := P{start.x - 1, start.y}
	reverse_dfs(dungeon, visited, hpmap, left, min_hp)
	up := P{start.x, start.y - 1}
	reverse_dfs(dungeon, visited, hpmap, up, min_hp)
	visited[start] = false
}

func calculateMinimumHP(dungeon [][]int) int {
	start := P{len(dungeon) - 1, len(dungeon[0]) - 1}
	visited := make(map[P]bool)
	hpmap := make(map[P]int)
	cur_hp := 1
	reverse_dfs(dungeon, visited, hpmap, start, cur_hp)
	return hpmap[P{0, 0}]
}

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(dungeon))
}
