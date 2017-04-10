package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// Naive solution, O(n^2)
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	max_s := math.MinInt32
	for start := 0; start < n; start++ {
		min_h := math.MaxInt32
		for end := start; end < n; end++ {
			min_h = min(min_h, heights[end])
			max_s = max(max_s, min_h*(end-start+1))
		}
	}
	return max_s
}

// Divide and conquer, O(NlogN)
func largestRectangleArea2(heights []int) int {
	var lra func(heights []int) int

	lra = func(heights []int) int {
		n := len(heights)
		if n == 0 {
			return 0
		} else if n == 1 {
			return heights[0] * 1
		}

		// find minimum height
		// TODO, should be use segment tree

		index := 0
		minimum := heights[index]

		for i := 1; i < n; i++ {
			if heights[i] < minimum {
				minimum = heights[i]
				index = i
			}
		}

		return max(
			n*heights[index],                                  // all
			max(lra(heights[:index]), lra(heights[index+1:])), // left or right
		)
	}

	return lra(heights)
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

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
