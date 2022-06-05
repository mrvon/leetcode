package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	S []int
}

func (S *Stack) len() int {
	return len(S.S)
}

func (S *Stack) push(t int) {
	S.S = append(S.S, t)
}

func (S *Stack) pop() (t int) {
	t = S.S[len(S.S)-1]
	S.S = S.S[:len(S.S)-1]
	return
}

func (S *Stack) peek() (t int) {
	t = S.S[len(S.S)-1]
	return
}

func evalRPN(tokens []string) int {
	var s Stack

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == "+" {
			r := s.pop()
			l := s.pop()
			s.push(l + r)
		} else if t == "-" {
			r := s.pop()
			l := s.pop()
			s.push(l - r)
		} else if t == "*" {
			r := s.pop()
			l := s.pop()
			s.push(l * r)
		} else if t == "/" {
			r := s.pop()
			l := s.pop()
			s.push(l / r)
		} else {
			// number
			i, err := strconv.Atoi(t)
			if err != nil {
				panic("invalid number.")
			}
			s.push(i)
		}
	}

	return s.pop()
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
