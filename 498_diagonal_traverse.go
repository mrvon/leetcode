package main

import "fmt"

func findDiagonalOrder(matrix [][]int) []int {
	M := len(matrix)
	if M == 0 {
		return []int{}
	}

	N := len(matrix[0])

	var list [][]int
	for i := 0; i < M; i++ {
		list = append(list, []int{})
		index := len(list) - 1

		m := i
		n := 0
		list[index] = append(list[index], matrix[m][n])
		for {
			// next one
			m--
			n++
			if m < 0 || n >= N {
				break
			}

			list[index] = append(list[index], matrix[m][n])
		}
	}

	for j := 1; j < N; j++ {
		list = append(list, []int{})
		index := len(list) - 1

		m := M - 1
		n := j
		list[index] = append(list[index], matrix[m][n])
		for {
			// next one
			m--
			n++
			if m < 0 || n >= N {
				break
			}

			list[index] = append(list[index], matrix[m][n])
		}
	}

	direction := true
	var order []int
	for i := 0; i < len(list); i++ {
		if direction {
			for j := 0; j < len(list[i]); j++ {
				order = append(order, list[i][j])
			}
		} else {
			for j := len(list[i]) - 1; j >= 0; j-- {
				order = append(order, list[i][j])
			}
		}
		direction = !direction
	}
	return order
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(matrix))
}
