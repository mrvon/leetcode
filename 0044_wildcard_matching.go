package main

import "fmt"

func isMatch(s string, p string) bool {
	si := 0
	pi := 0
	star := -1
	match := 0

	for si < len(s) {
		if pi < len(p) && (p[pi] == '?' || p[pi] == s[si]) {
			// advancing both pointers
			pi++
			si++
		} else if pi < len(p) && p[pi] == '*' {
			// *, only advancing pattern pointer
			star = pi
			match = si
			pi++
		} else if star != -1 {
			// last pattern was *, advancing string pointer
			pi = star + 1
			match++
			si = match
		} else {
			// current pattern pointer is not star, last pattern pointer was not
			// * characters do not match
			return false
		}
	}

	// check for remaining characters in pattern
	for pi < len(p) && p[pi] == '*' {
		pi++
	}

	return pi == len(p)
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(false, isMatch("aa", "a"))
	assert(true, isMatch("aa", "aa"))
	assert(false, isMatch("aaa", "aa"))
	assert(true, isMatch("aa", "*"))
	assert(true, isMatch("aa", "a*"))
	assert(true, isMatch("ab", "?*"))
	assert(false, isMatch("aab", "c*a*b"))
	assert(true, isMatch("aab", "*"))
	assert(true, isMatch("aab", "***a***a***b***"))
}
