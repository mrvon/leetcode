// Math
package main

import "fmt"

func length2(v []int) int {
	return v[0]*v[0] + v[1]*v[1]
}

func dot(v1 []int, v2 []int) int {
	return v1[0]*v1[1] + v2[0]*v2[1]
}

func vector(p1 []int, p2 []int) []int {
	return []int{
		p1[0] - p2[0],
		p1[1] - p2[1],
	}
}

func valid(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	v1 := vector(p1, p2)
	v2 := vector(p2, p3)
	v3 := vector(p3, p4)
	v4 := vector(p4, p1)

	l1 := length2(v1)
	l2 := length2(v2)
	l3 := length2(v3)
	l4 := length2(v4)

	if l1 == 0 || l2 == 0 || l3 == 0 || l4 == 0 {
		return false
	}

	d1 := dot(v1, v2)
	d2 := dot(v2, v3)
	d3 := dot(v3, v4)
	d4 := dot(v4, v1)

	if d1 != 0 {
		return false
	}

	return d1 == d2 && d2 == d3 && d3 == d4 &&
		l1 == l2 && l2 == l3 && l3 == l4
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return valid(p1, p2, p3, p4) ||
		valid(p1, p2, p4, p3) ||
		valid(p1, p3, p2, p4) ||
		valid(p1, p3, p4, p2) ||
		valid(p1, p4, p2, p3) ||
		valid(p1, p4, p3, p2)
}

func main() {
	fmt.Println(validSquare(
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
		[]int{0, 1},
	))

	fmt.Println(validSquare(
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
	))

	fmt.Println(validSquare(
		[]int{1, 1},
		[]int{5, 3},
		[]int{3, 5},
		[]int{7, 7},
	))
}
