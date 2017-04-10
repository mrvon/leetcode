package main

import "fmt"

func max(height []int, lo int, hi int) int {
	if hi-lo < 0 {
		return -1
	}
	m_val := height[lo]
	m_idx := lo
	for i := lo + 1; i <= hi; i++ {
		if height[i] > m_val { // must be large than and not equal to
			m_val = height[i]
			m_idx = i
		}
	}
	return m_idx
}

func fill(height []int, lo int, hi int) int {
	if hi-lo < 2 {
		return 0
	}

	min := height[lo]
	if height[hi] < min {
		min = height[hi]
	}

	d := 0
	for i := lo + 1; i <= hi-1; i++ {
		d = d + (min - height[i])
	}
	return d
}

func aux_trap(height []int, lo int, hi int) int {
	c := 0
	m := max(height, lo, hi)
	if m == -1 {
		return 0
	}

	l := max(height, lo, m-1)
	if l != -1 {
		c += fill(height, l, m)
		c += aux_trap(height, lo, l)
	}

	r := max(height, m+1, hi)
	if r != -1 {
		c += fill(height, m, r)
		c += aux_trap(height, r, hi)
	}

	return c
}

func trap(height []int) int {
	return aux_trap(height, 0, len(height)-1)
}

func main() {
	fmt.Println(trap([]int{2}))
	fmt.Println(trap([]int{1, 2}))
	fmt.Println(trap([]int{1, 0, 2}))
	fmt.Println(trap([]int{0, 1, 0, 2}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
