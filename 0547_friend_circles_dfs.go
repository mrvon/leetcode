// using DFS to find count of connected components
package main

import "fmt"

func dfs(M [][]int, visited map[int]bool, i int) {
	for j := 0; j < len(M[i]); j++ {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(M, visited, j)
		}
	}
}

func findCircleNum(M [][]int) int {
	// count of connected components
	count := 0
	visited := make(map[int]bool)

	for i := 0; i < len(M); i++ {
		if !visited[i] {
			visited[i] = true
			count++
			dfs(M, visited, i)
		}
	}

	return count
}

func main() {
	matrix := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	fmt.Println(findCircleNum(matrix))

	matrix = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	fmt.Println(findCircleNum(matrix))
}
