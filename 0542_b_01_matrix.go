/*
In-place, BFS, start point is 0 node
*/
package main

import (
	"fmt"
	"math"
)

type Node struct {
	x int
	y int
}

type Queue []Node

func (Q *Queue) push(p Node) {
	(*Q) = append(*Q, p)
}

func (Q *Queue) pop() Node {
	p := (*Q)[0]
	(*Q) = (*Q)[1:]
	return p
}

func (Q *Queue) isempty() bool {
	if len(*Q) == 0 {
		return true
	} else {
		return false
	}
}

var direction = []int{
	-1, 0,
	1, 0,
	0, -1,
	0, 1,
}

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	q := Queue{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				// pretty good!
				q.push(Node{i, j})
			} else {
				matrix[i][j] = math.MaxInt32
			}
		}
	}

	for !q.isempty() {
		p := q.pop()

		for k := 0; k < len(direction); k += 2 {
			x := p.x + direction[k]
			y := p.y + direction[k+1]

			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			if matrix[x][y] < matrix[p.x][p.y]+1 {
				// key point
				continue
			}

			matrix[x][y] = matrix[p.x][p.y] + 1
			q.push(Node{x, y})
		}
	}

	return matrix
}

func print_matrix(distance [][]int) {
	for i := 0; i < len(distance); i++ {
		fmt.Println(distance[i])
	}
}

func main() {
	matrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	matrix = updateMatrix(matrix)
	print_matrix(matrix)

	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	matrix = updateMatrix(matrix)
	print_matrix(matrix)
}
