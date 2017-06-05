package main

import (
	"fmt"
	"math"
)

// Time O(n), Space O(1)
// Generalization of Boyer–Moore majority vote algorithm
func majorityElement(nums []int) []int {
	// At most have two numbers appear more than 1/3, If a number's
	// frequent > 1/3, then it must have recorded as x or y, because other
	// number's frequent < 1/3, they can't kick x or y out.
	// notable here, x_major, y_major's initial value is not relevant
	x_major := 0
	x_count := 0
	y_major := 1
	y_count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == x_major {
			x_count++
		} else if nums[i] == y_major {
			y_count++
		} else if x_count == 0 {
			x_major = nums[i]
			x_count++
		} else if y_count == 0 {
			y_major = nums[i]
			y_count++
		} else {
			x_count--
			y_count--
		}
	}
	x_count = 0
	y_count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == x_major {
			x_count++
		} else if nums[i] == y_major {
			y_count++
		}
	}
	bound := int(math.Floor(float64(len(nums)) / 3.0))
	majors := []int{}
	if x_count > bound {
		majors = append(majors, x_major)
	}
	if y_count > bound {
		majors = append(majors, y_major)
	}
	return majors
}

func main() {
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
	fmt.Println(majorityElement([]int{3, 0, 3, 4}))
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 3, 3, 3, 3}))
	fmt.Println(majorityElement([]int{1, 1, 3, 3, 3, 3}))
}
