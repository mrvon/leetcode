/*
1.Two Queue and add all the Pacific border to one queue; Atlantic border to
another queue.

2.Keep a visited matrix for each queue. In the end, add the cell visited by two
queue to the result.
*/
package main

import "fmt"

type Grid struct {
	x int
	y int
}

type Queue struct {
	Q []Grid
}

func (Q *Queue) len() int {
	return len(Q.Q)
}

func (Q *Queue) push(g Grid) {
	Q.Q = append(Q.Q, g)
}

func (Q *Queue) pop() (g Grid) {
	g = Q.Q[0]
	Q.Q = Q.Q[1:len(Q.Q)]
	return
}

func (Q *Queue) peek() (g Grid) {
	g = Q.Q[len(Q.Q)-1]
	return
}

var direction = []int{
	-1, 0,
	1, 0,
	0, -1,
	0, 1,
}

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return [][]int{}
	}
	n := len(matrix[0])

	p_visited := [][]bool{}
	a_visited := [][]bool{}

	for i := 0; i < m; i++ {
		p_visited = append(p_visited, make([]bool, len(matrix[i])))
		a_visited = append(a_visited, make([]bool, len(matrix[i])))
	}

	p_queue := Queue{}
	a_queue := Queue{}

	for j := 0; j < n; j++ {
		p_queue.push(Grid{0, j})
		p_visited[0][j] = true
		a_queue.push(Grid{m - 1, j})
		a_visited[m-1][j] = true
	}
	for i := 0; i < m; i++ {
		p_queue.push(Grid{i, 0})
		p_visited[i][0] = true
		a_queue.push(Grid{i, n - 1})
		a_visited[i][n-1] = true
	}

	for p_queue.len() > 0 {
		grid := p_queue.pop()

		for i := 0; i < len(direction); i += 2 {
			x := grid.x + direction[i]
			y := grid.y + direction[i+1]

			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			if p_visited[x][y] {
				continue
			}

			if matrix[grid.x][grid.y] > matrix[x][y] {
				continue
			}

			p_visited[x][y] = true
			p_queue.push(Grid{x, y})
		}
	}

	for a_queue.len() > 0 {
		grid := a_queue.pop()

		for i := 0; i < len(direction); i += 2 {
			x := grid.x + direction[i]
			y := grid.y + direction[i+1]

			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			if a_visited[x][y] {
				continue
			}

			if matrix[grid.x][grid.y] > matrix[x][y] {
				continue
			}

			a_visited[x][y] = true
			a_queue.push(Grid{x, y})
		}
	}

	results := [][]int{}

	for i := 0; i < len(p_visited); i++ {
		for j := 0; j < len(p_visited[i]); j++ {
			if p_visited[i][j] && a_visited[i][j] {
				results = append(results, []int{i, j})
			}
		}
	}

	return results
}

func main() {
	// Given the following 5x5 matrix:
	//   Pacific ~   ~   ~   ~   ~
	//        ~  1   2   2   3  (5) *
	//        ~  3   2   3  (4) (4) *
	//        ~  2   4  (5)  3   1  *
	//        ~ (6) (7)  1   4   5  *
	//        ~ (5)  1   1   2   4  *
	//           *   *   *   *   * Atlantic
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlantic(matrix))

	matrix = [][]int{
		{1, 2},
		{4, 3},
	}
	fmt.Println(pacificAtlantic(matrix))
}
