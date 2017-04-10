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

func dist(matrix [][]int, i int, j int) int {
	m := len(matrix)
	n := len(matrix[0])

	s := make(map[Node]int)
	q := Queue{}

	p := Node{i, j}
	q.push(p)
	s[p] = 1

	for !q.isempty() {
		p := q.pop()

		for k := 0; k < len(direction); k += 2 {
			x := p.x + direction[k]
			y := p.y + direction[k+1]

			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			if matrix[x][y] == 0 {
				return s[p]
			}

			np := Node{x, y}
			if s[np] > 0 {
				continue
			}

			q.push(np)
			s[np] = s[p] + 1
		}
	}

	return math.MaxInt32
}

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				distance[i][j] = 0
			} else {
				distance[i][j] = dist(matrix, i, j)
			}
		}
	}

	return distance
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
	distance := updateMatrix(matrix)
	print_matrix(distance)

	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	distance = updateMatrix(matrix)
	print_matrix(distance)
}
