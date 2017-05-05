/*
Array[A, B, C, D, E]

F(0) = (0A) + (1B) + (2C) + (3D) + (4E)
F(1) = (4A) + (0B) + (1C) + (2D) + (3E)
F(2) = (3A) + (4B) + (0C) + (1D) + (2E)
F(3) = (2A) + (3B) + (4C) + (0D) + (1E)
F(4) = (1A) + (2B) + (3C) + (4D) + (0E)

C = (1A) + (1B) + (1C) + (1D) + (1E)

We can contruct F(1) from F(0)

F(0) - C = (-1A) + (0B) + (1C) + (2D) + (3E)
F(0) - C + (5A) = (4A) + (0B) + (1C) + (2D) + (3E)
				= F(1)

And contruct F(2) from F(1)
F(1) - C = (3A) + (-1B) + (0C) + (1D) + (2E)
F(1) - C + (5B) = (3A) + (4B) + (0C) + (1D) + (2E)

and so on.
*/
package main

import (
	"fmt"
)

func maxRotateFunction(A []int) int {
	F0 := 0
	C := 0
	N := len(A)

	for i := 0; i < N; i++ {
		F0 += (A[i] * i)
		C += A[i]
	}

	max := F0
	FN := F0
	for i := 1; i < N; i++ {
		// contruct FN
		FN = FN - C + A[i-1]*N
		if FN > max {
			max = FN
		}
	}
	return max
}

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
	fmt.Println(maxRotateFunction([]int{}))
	fmt.Println(maxRotateFunction([]int{-2147483648, -2147483648}))
	fmt.Println(maxRotateFunction([]int{1, 2, 3}))
	fmt.Println(maxRotateFunction([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
