package main

import (
	"fmt"
	"math"
)

func shortestToChar(S string, C byte) []int {
	c := rune(C)
	position := []int{}

	for i, s := range S {
		if s == c {
			position = append(position, i)
		}
	}

	distance := []int{}
	for i, s := range S {
		if s == c {
			distance = append(distance, 0)
		} else {
			// binary search
			shortest := math.MaxInt32
			left := 0
			right := len(position) - 1
			for left <= right {
				mid := left + (right-left)/2
				// assert(i != position[mid])
				d := 0
				if i < position[mid] {
					right = mid - 1
					d = position[mid] - i
				} else {
					left = mid + 1
					d = i - position[mid]
				}
				if d < shortest {
					shortest = d
				}
			}
			distance = append(distance, shortest)
		}
	}
	return distance
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
