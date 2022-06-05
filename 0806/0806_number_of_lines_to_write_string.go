package main

import "fmt"

const (
	lineWidth = 100
)

func numberOfLines(widths []int, S string) []int {
	count := 1
	curr := 0
	for _, s := range S {
		w := widths[s-'a']
		if curr+w > lineWidth {
			count++
			curr = 0
		}
		curr += w
	}
	return []int{count, curr}
}

func main() {
	widths := []int{
		4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		10, 10, 10, 10, 10, 10, 10, 10,
	}
	S := "bbbcccdddaaa"

	fmt.Println(numberOfLines(widths, S))
}
