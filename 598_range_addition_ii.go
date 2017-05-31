package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	for i := 0; i < len(ops); i++ {
		a := ops[i][0]
		b := ops[i][1]
		if a < m {
			m = a
		}
		if b < n {
			n = b
		}
	}
	return m * n
}

func main() {
	fmt.Println(maxCount(
		3,
		3,
		[][]int{{2, 2}, {3, 3}},
	))

	fmt.Println(maxCount(
		40000,
		40000,
		[][]int{},
	))

	fmt.Println(maxCount(
		39999,
		39999,
		[][]int{{19999, 19999}},
	))
}
