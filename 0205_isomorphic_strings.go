package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]byte)
	used := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if mapping[s[i]] == 0 {
			if used[t[i]] {
				return false
			}
			used[t[i]] = true
			mapping[s[i]] = t[i]
		} else {
			if mapping[s[i]] != t[i] {
				return false
			}
		}
	}

	return true
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(true, isIsomorphic("egg", "add"))
	assert(false, isIsomorphic("foo", "bar"))
	assert(true, isIsomorphic("paper", "title"))
	assert(true, isIsomorphic("ken", "ken"))
	assert(false, isIsomorphic("kkn", "ten"))
	assert(false, isIsomorphic("aba", "baa"))
	assert(false, isIsomorphic("ab", "aa"))
}
