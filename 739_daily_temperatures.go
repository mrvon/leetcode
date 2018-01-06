// Stack solution
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	warmers := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := -1
	for i := 0; i < len(temperatures); i++ {
		for top >= 0 && temperatures[i] > temperatures[stack[top]] {
			warmers[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}

	return warmers
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{73, 73, 73, 73}))
}
