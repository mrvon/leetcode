package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func isSubseq(a string, b string) bool {
	// len(a) >= len(b)
	j := 0
	for i := 0; i < len(a); i++ {
		if j < len(b) && b[j] == a[i] {
			j++
		}
	}
	return j == len(b)
}

func findLUSlength(strs []string) int {
	if len(strs) == 0 {
		return 0
	} else if len(strs) == 1 {
		return len(strs[0])
	}

	sort.Sort(ByLength(strs))

	for i := 0; i < len(strs); i++ {
		is_en := false
		is_ok := true
		for j := i + 1; j <= len(strs)-1; j++ {
			is_en = true
			if strs[i] == strs[j] {
				is_ok = false
				break
			}
		}
		if is_en && !is_ok {
			continue
		}

		is_en = false
		is_ok = true
		for j := i - 1; j >= 0; j-- {
			is_en = true
			if isSubseq(strs[j], strs[i]) {
				is_ok = false
				break
			}
		}
		if is_en && !is_ok {
			continue
		}

		return len(strs[i])
	}

	return -1
}

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(3, findLUSlength([]string{"aaa"}))
	assert(3, findLUSlength([]string{"aaa", "aa"}))
	assert(-1, findLUSlength([]string{"aaa", "aaa", "aa"}))
	assert(2, findLUSlength([]string{"aaa", "aaa", "kk"}))
	assert(-1, findLUSlength([]string{"aaa", "aaa", "kk", "kk"}))
	assert(6, findLUSlength([]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"}))
	assert(-1, findLUSlength([]string{"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"}))
}
