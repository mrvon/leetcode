package main

import "fmt"

type Jug struct {
	water    int
	capacity int
}

func canMeasureWater(x int, y int, z int) bool {
	// 1. Fill any of the jugs completely with water.
	// 2. Empty any of the jugs.
	// 3. Pour water from one jug into another till the other jug is completely
	// full or the first jug itself is empty.
	x_jug := Jug{0, x}
	y_jug := Jug{0, y}
	z_jug := Jug{0, z}
}

func main() {
	// Input: x = 3, y = 5, z = 4
	// Output: True
	fmt.Println(canMeasureWater(3, 5, 4))
}
