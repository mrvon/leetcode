/*
Dynamic programming
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

type Cpoint struct {
	X int
	Y int
	C int
}
type Slope struct {
	s float64
	c int
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func sameSlope(slope_list []Slope) int {
	set := make(map[float64]int)
	m := 0

	for i := 0; i < len(slope_list); i++ {
		slope := slope_list[i]
		set[slope.s] += slope.c
		if set[slope.s] > m {
			m = set[slope.s]
		}
	}

	return m
}

func countPoint(points []Point) []Cpoint {
	set := make(map[Point]int)
	for i := 0; i < len(points); i++ {
		set[points[i]]++
	}

	cpoints := []Cpoint{}
	for p, c := range set {
		cpoints = append(cpoints, Cpoint{p.X, p.Y, c})
	}
	return cpoints
}

func maxPoints(points []Point) int {
	cpoints := countPoint(points)

	if len(cpoints) == 0 {
		return 0
	}

	optimal := make([]int, len(cpoints))
	// only one point
	optimal[0] = cpoints[0].C

	for i := 1; i < len(cpoints); i++ {
		p1 := cpoints[i]
		slope_list := []Slope{}
		for j := 0; j < i; j++ {
			p2 := cpoints[j]
			if p1.X == p2.X {
				slope_list = append(slope_list, Slope{math.MaxFloat64, p2.C})
			} else {
				slope_list = append(slope_list, Slope{float64(p1.Y-p2.Y) / float64(p1.X-p2.X), p2.C})
			}
		}
		count := sameSlope(slope_list) + p1.C
		optimal[i] = max(optimal[i-1], count)
	}

	return optimal[len(cpoints)-1]
}

func main() {
	fmt.Println(maxPoints(
		[]Point{
			Point{1, 1},
			Point{0, 0},
			Point{1, 1},

			// Point{1, 1},
			// Point{2, 2},
			// Point{3, 3},
			// Point{4, 4},
			// Point{5, 5},
			// Point{5, 5},
			// Point{5, 5},
			// Point{10, 8},
			// Point{0, 5},
			// Point{5, 0},
		},
	))
}
