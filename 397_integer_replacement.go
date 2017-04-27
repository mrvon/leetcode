package main

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func ir(n int, m map[int]int) int {
	if n <= 1 {
		return 0
	}

	if c, e := m[n]; e {
		return c
	}

	c := 0
	if n%2 == 0 {
		c = 1 + ir(n/2, m)
	} else {
		c = 1 + min(ir(n-1, m), ir(n+1, m))
	}
	m[n] = c
	return c
}

func integerReplacement(n int) int {
	return ir(n, make(map[int]int))
}
