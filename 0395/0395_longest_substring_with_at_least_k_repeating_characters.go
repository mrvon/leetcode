// Divide and conquer
package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for c, n := range m {
		if n < k {
			max := 0
			for _, t := range strings.Split(s, string(c)) {
				l := longestSubstring(t, k)
				if l > max {
					max = l
				}
			}
			return max
		}
	}
	return len(s)
}

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(3, longestSubstring("aaabb", 3))
	assert(5, longestSubstring("ababbc", 2))
	assert(6, longestSubstring("aaabbb", 3))
	assert(0, longestSubstring("ababacb", 3))
}
