package main

import "fmt"

// kmp matcher
func kmp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	p := prefix(pattern)
	q := 0                   // number of chanracters matched
	for i := 0; i < n; i++ { // scan the text from left to right
		for q > 0 && pattern[q] != text[i] {
			q = p[q-1] // next character does not match
		}
		if pattern[q] == text[i] {
			q = q + 1 // next character matches
		}
		if q == m { // is all of P matched?
			return i - m + 1
		}
	}
	return -1
}

// compute prefix function
func prefix(pattern string) []int {
	m := len(pattern)
	p := make([]int, m)
	p[0] = 0
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = p[k-1]
		}
		if pattern[k] == pattern[q] {
			k = k + 1
		}
		p[q] = k
	}
	return p
}

func main() {
	fmt.Println(kmp("", ""))
	fmt.Println(kmp("hello", "hel"))
	fmt.Println(kmp("hello", "el"))
	fmt.Println(kmp("hello", "ll"))
	fmt.Println(kmp("hello", "k"))
	fmt.Println(kmp("mississippi", "issip"))
}
