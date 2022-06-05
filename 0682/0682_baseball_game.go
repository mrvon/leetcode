package main

import (
	"fmt"
	"strconv"
)

func last(scores []int, index int) int {
	if len(scores) >= index {
		return scores[len(scores)-index]
	} else {
		return 0
	}
}

func calPoints(ops []string) int {
	scores := []int{}
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op {
		case "+":
			scores = append(scores, last(scores, 1)+last(scores, 2))
		case "D":
			scores = append(scores, last(scores, 1)*2)
		case "C":
			scores = scores[:len(scores)-1]
		default:
			n, _ := strconv.Atoi(op)
			scores = append(scores, n)
		}
	}
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5"}))
	fmt.Println(calPoints([]string{"5", "2"}))
	fmt.Println(calPoints([]string{"5", "2", "C"}))
	fmt.Println(calPoints([]string{"5", "2", "C", "D"}))
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
