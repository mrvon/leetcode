package main

import (
	"fmt"
	"math"
)

/*
l >= w, l * w == area

firstly, I wrote the first version of the program, following

func ConstructRectangle(area int) []int {
	min_l := int(math.Ceil(math.Sqrt(float64(area))))
	for l := min_l; l <= area; l++ {
		if area%l == 0 {
			return []int{
				l, area / l,
			}
		}
	}
	// never be here
	return []int{}
}

It's very slow, because it's will very slow when area is a prime number, In that
case, it need to try (A = area - sqrt(area) + 1) times to find the answer.

the second version in the bottom is better than first version, it's because
it at most try B = sqrt(area) times. At most of time, A is much large than B.
*/

func constructRectangle(area int) []int {
	max_w := int(math.Sqrt(float64(area)))
	for w := max_w; w > 0; w-- {
		l := area / w
		if l*w == area {
			return []int{
				l, w,
			}
		}
	}
	// never be here
	return []int{}
}

func main() {
	fmt.Println(constructRectangle(1))
	fmt.Println(constructRectangle(2))
	fmt.Println(constructRectangle(3))
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(5))
	fmt.Println(constructRectangle(6))
	fmt.Println(constructRectangle(7))
	fmt.Println(constructRectangle(8))
	fmt.Println(constructRectangle(9))
}
