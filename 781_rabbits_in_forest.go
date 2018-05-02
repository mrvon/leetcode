// Group same answer's rabbit together
package main

import (
	"fmt"
	"math"
)

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}

	s := make(map[int]int)

	for _, a := range answers {
		s[a]++
	}

	count := 0

	for a, c := range s {
		if a == 0 {
			count += c
		} else {
			count += int(math.Ceil(float64(c)/float64(a+1))) * (a + 1)
		}
	}

	return count
}

func main() {
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{1, 1, 2, 2, 2}))
	fmt.Println(numRabbits([]int{1, 1, 2, 2, 2, 3, 3}))
}
