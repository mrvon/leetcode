// Manhattan distance
package main

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	max := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		distance := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if distance <= max {
			return false
		}
	}
	return true
}
