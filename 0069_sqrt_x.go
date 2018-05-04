/*
Newton's method

Formula is that:

Sqrt(x) ->
	X[k+1] = 1.0/2.0*(X[k] + N / X[k])

X[0] = 1
*/
package main

import (
	"fmt"
	"math"
)

const (
	PRECISION = 0.1
)

func newtonSqrt(N float64) float64 {
	X := float64(1.0)

	for {
		O := X
		X = 1.0 / 2.0 * (X + N/X)
		if math.Abs(X-O) < PRECISION {
			break
		}
	}

	return X
}

func mySqrt(x int) int {
	return int(newtonSqrt(float64(x)))
}

func main() {
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(5))
	fmt.Println(newtonSqrt(2.0))
	fmt.Println(newtonSqrt(3.0))
}
