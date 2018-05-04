package main

import "fmt"

type State struct {
	text  string
	max   int
	trace []byte
	perms []string
}

func isletter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func flip(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 'A'
	} else {
		return c + 'a' - 'A'
	}
}

func perm(s *State, index int) {
	if index > s.max {
		s.perms = append(s.perms, string(s.trace))
		return
	}

	s.trace[index] = s.text[index]
	if isletter(s.text[index]) {
		perm(s, index+1)
		s.trace[index] = flip(s.text[index])
		perm(s, index+1)
	} else {
		perm(s, index+1)
	}
}

func letterCasePermutation(S string) []string {
	s := &State{
		text:  S,
		max:   len(S) - 1,
		trace: make([]byte, len(S)),
		perms: []string{},
	}

	perm(s, 0)

	return s.perms
}

func main() {
	fmt.Println(letterCasePermutation("a"))
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation(("3z4")))
	fmt.Println(letterCasePermutation(("12345")))
}
