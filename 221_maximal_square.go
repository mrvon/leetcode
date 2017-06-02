/*
Though we can slightly modified 85_maximal_rectangle.go can pass this problem.
But it has another simply approach.

Dynamic programming
*/
package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	// TODO
}
func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
