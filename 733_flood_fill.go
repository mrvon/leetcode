package main

import "fmt"

var direction = []int{
	-1, 0,
	1, 0,
	0, -1,
	0, 1,
}

func check(image [][]int, visited [][]bool, tr int, tc int, oldColor int) bool {
	if tr < 0 || tr >= len(image) {
		return false
	}

	if tc < 0 || tc >= len(image[tr]) {
		return false
	}

	if visited[tr][tc] {
		return false
	}

	return oldColor == image[tr][tc]
}

func fill(image [][]int, visited [][]bool, sr int, sc int, oldColor int, newColor int) [][]int {
	visited[sr][sc] = true
	image[sr][sc] = newColor

	for i := 0; i < len(direction); i += 2 {
		x := direction[i]
		y := direction[i+1]
		if check(image, visited, sr+x, sc+y, oldColor) {
			fill(image, visited, sr+x, sc+y, oldColor, newColor)
		}
	}

	return image
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var visited [][]bool
	for i := 0; i < len(image); i++ {
		visited = append(visited, make([]bool, len(image[i])))
	}
	return fill(image, visited, sr, sc, image[sr][sc], newColor)
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
	// image = [[1,1,1],[1,1,0],[1,0,1]]
	// sr = 1, sc = 1, newColor = 2
	// Output: [[2,2,2],[2,2,0],[2,0,1]]
}
