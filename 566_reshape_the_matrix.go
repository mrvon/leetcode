package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m := len(nums)
	if m == 0 {
		return nums
	}

	n := len(nums[0])

	if m*n != r*c {
		return nums
	}

	p := 0
	q := 0
	matrix := [][]int{}
	for i := 0; i < r; i++ {
		matrix = append(matrix, make([]int, c))
		for j := 0; j < c; j++ {
			matrix[i][j] = nums[p][q]
			if q < n-1 {
				q++
			} else {
				p++
				q = 0
			}
		}
	}
	return matrix
}

func main() {
	nums := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	fmt.Println(matrixReshape(nums, 1, 4))

	nums = [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	fmt.Println(matrixReshape(nums, 1, 5))
}
