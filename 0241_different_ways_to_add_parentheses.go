/* Divide and conquer

For example

2*3-4*5

Divide into following sub-problem

"2" * "3-4*5"
"2*3" - "4*5"
"2*3-4" * "5"

When the sub-problem is a single number

It's value is the number itself.
*/

package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	var result []int
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '+' || c == '-' || c == '*' {
			for _, m := range diffWaysToCompute(input[0:i]) {
				for _, n := range diffWaysToCompute(input[i+1 : len(input)]) {
					if c == '+' {
						result = append(result, m+n)
					} else if c == '-' {
						result = append(result, m-n)
					} else {
						result = append(result, m*n)
					}
				}
			}
		}
	}
	if len(result) == 0 { // single number
		n, _ := strconv.Atoi(input)
		return []int{
			n,
		}
	} else { // expression with operation
		return result
	}
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
