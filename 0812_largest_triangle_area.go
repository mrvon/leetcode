package main

import (
	"fmt"
	"math"
)

func area2(x []int, y []int, z []int) float64 {
	// https: //www.mathopenref.com/coordtrianglearea.html
	return math.Abs(
		float64(x[0])*float64(y[1]-z[1])+
			float64(y[0])*float64(z[1]-x[1])+
			float64(z[0])*float64(x[1]-y[1])) / float64(2.0)
}

func largestTriangleArea(points [][]int) float64 {
	m := float64(0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				m = math.Max(m, area2(points[i], points[j], points[k]))
			}
		}
	}
	return m
}

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}
