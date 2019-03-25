// Stack approach
package main

import "fmt"

func minAddToMakeValid(S string) int {
	stack := []rune{}
	count := 0
	for _, s := range S {
		if s == '(' {
			stack = append(stack, s)
		} else {
			// must be ')'
			if len(stack) == 0 {
				count++
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	// additive remain '(' in stack
	return count + len(stack)
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("((()))"))
	fmt.Println(minAddToMakeValid("()))(("))
}
