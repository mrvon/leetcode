package main

import (
	"fmt"
	"math"
)

func imageSmoother(M [][]int) [][]int {
	h := len(M)
	w := len(M[0])

	grey := func(x int, y int) int {
		count := 0
		sum := 0
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i < 0 || i >= h {
					continue
				}
				if j < 0 || j >= w {
					continue
				}
				count++
				sum += M[i][j]
			}
		}
		return int(math.Floor(float64(sum) / float64(count)))
	}

	N := [][]int{}
	for i := 0; i < h; i++ {
		N = append(N, []int{})
		for j := 0; j < w; j++ {
			N[i] = append(N[i], grey(i, j))
		}
	}
	return N
}

func main() {
	fmt.Println(
		imageSmoother(
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1}}))
}
