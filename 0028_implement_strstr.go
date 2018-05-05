package main

import "fmt"

// Naive brute force
func strStr(haystack string, needle string) int {
	h := []byte(haystack)
	n := []byte(needle)

	for i := 0; i <= len(h)-len(n); i++ {
		is_same := true
		for j := 0; j < len(n); j++ {
			if h[i+j] != n[j] {
				is_same = false
				break
			}
		}
		if is_same {
			return i
		}
	}
	return -1
}

// kmp matcher
func kmpStrStr(text string, pattern string) int {
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

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(-1, strStr("hell", "hello"))
	assert(0, strStr("hello", "hello"))
	assert(-1, strStr("hello", " hello"))
	assert(1, strStr(" hello", "hello"))
	assert(6, strStr("world hello", "hello"))
	assert(-1, strStr("world ello", "hello"))
}
