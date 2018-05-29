package main

import "fmt"

func leastBricks(wall [][]int) int {
	for i := 0; i < len(wall); i++ {
		for j := 0; j < len(wall[i]); j++ {
			if j > 0 {
				wall[i][j] += wall[i][j-1]
			}
		}
	}

	counter := make(map[int]int)
	max := 0 // no cross
	for i := 0; i < len(wall); i++ {
		for j := 0; j < len(wall[i])-1; j++ {
			counter[wall[i][j]]++
			if counter[wall[i][j]] > max {
				max = counter[wall[i][j]]
			}
		}
	}

	// cross
	return len(wall) - max
}

func main() {
	wall := [][]int{
		[]int{1, 2, 2, 1},
		[]int{3, 1, 2},
		[]int{1, 3, 2},
		[]int{2, 4},
		[]int{3, 1, 2},
		[]int{1, 3, 1, 1},
	}
	fmt.Println(leastBricks(wall))
}
