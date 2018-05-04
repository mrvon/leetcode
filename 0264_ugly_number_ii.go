package main

import "fmt"

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= z {
		return y
	} else {
		return z
	}
}

func nthUglyNumber(n int) int {
	list := []int{1}

	i2 := 0
	i3 := 0
	i5 := 0

	for i := 1; i < n; i++ {
		x := 2 * list[i2]
		y := 3 * list[i3]
		z := 5 * list[i5]

		m := min(x, y, z)
		list = append(list, m)

		if x == m {
			i2++
		}
		if y == m {
			i3++
		}
		if z == m {
			i5++
		}
	}

	return list[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
