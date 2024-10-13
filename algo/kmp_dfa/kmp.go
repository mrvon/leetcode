package main

import "fmt"

const Radix = 256

func dfa(pattern string) [][]int {
	dfa := [][]int{}
	m := len(pattern)
	for i := 0; i < Radix; i++ {
		dfa = append(dfa, make([]int, m))
	}
	dfa[pattern[0]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < Radix; c++ {
			dfa[c][j] = dfa[c][x] // copy mismatch cases
		}
		dfa[pattern[j]][j] = j + 1 // set match case
		x = dfa[pattern[j]][x]
	}
	return dfa
}

func kmp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	dfa := dfa(pattern)
	i := 0
	j := 0
	for ; i < n && j < m; i++ {
		j = dfa[text[i]][j]
	}
	if j == m {
		return i - m // found
	}
	return -1 // not found
}

func main() {
	fmt.Println(kmp("", ""))
	fmt.Println(kmp("hello", "hel"))
	fmt.Println(kmp("hello", "el"))
	fmt.Println(kmp("hello", "ll"))
	fmt.Println(kmp("hello", "k"))
	fmt.Println(kmp("mississippi", "issip"))
}
