// Base on 84_largest_rectangle_in_histogram, very easy.
package main

import "fmt"

// code copy from 84_largest_rectangle_in_histogram.go
func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

type Stack []int

func (S *Stack) len() int {
	return len(*S)
}

func (S *Stack) push(n int) {
	(*S) = append(*S, n)
}

func (S *Stack) pop() int {
	n := (*S)[len(*S)-1]
	(*S) = (*S)[:len(*S)-1]
	return n
}

func (S *Stack) peek() int {
	return (*S)[len(*S)-1]
}

// http://www.geeksforgeeks.org/largest-rectangle-under-histogram/
// Stack, O(N)
func largestRectangleArea(heights []int) int {
	// Create an empty stack. The stack holds indexes of heights[] array
	// The bars stored in stack are always in increasing order of their heights.
	var s Stack

	// Max area
	max_area := 0

	// Run through all bars of given histogram
	i := 0
	for i < len(heights) {
		if s.len() == 0 || heights[s.peek()] <= heights[i] {
			// If this bar is higher than the bar on top stack, push it to stack
			s.push(i)
			// Next
			i++
		} else {
			// If this bar is lower than top of stack, then calculate area of
			// rectangle with stack top as the smallest (or minimum height) bar.
			// 'i' is 'right index' for the top and element before top in stack
			// is 'left index'
			top := s.pop()

			// Calculate the area with heights[top] stack as smallest bar
			area := 0
			if s.len() == 0 {
				area = heights[top] * i
			} else {
				area = heights[top] * (i - s.peek() - 1)
			}

			max_area = max(max_area, area)
		}
	}

	// Now pop the remaining bars from stack and calculate area with every
	// popped bar as the smallest bar
	for s.len() > 0 {
		top := s.pop()
		area := 0
		if s.len() == 0 {
			area = heights[top] * i
		} else {
			area = heights[top] * (i - s.peek() - 1)
		}
		max_area = max(max_area, area)
	}

	return max_area
}

// -------------------------------------------------------

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	max_area := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}

		max_area = max(max_area, largestRectangleArea(heights))
	}

	return max_area
}

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
