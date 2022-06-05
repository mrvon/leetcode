package main

import "fmt"

// O(n^3)
func __countSubstrings(s string) int {
	isPalindromic := func(s string) bool {
		i := 0
		j := len(s) - 1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	count := 0
	for end := 1; end <= len(s); end++ {
		for start := end - 1; start >= 0; start-- {
			sub := s[start:end]
			if isPalindromic(sub) {
				count++
			}
		}
	}
	return count
}

// Idea is start from each index and try to extend palindrome for both odd and
// even length.
func countSubstrings(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}

	count := 0

	extend := func(left int, right int) {
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
			count++
		}
	}

	for mid := 0; mid < l; mid++ {
		extend(mid, mid)
		extend(mid, mid+1)
	}

	return count
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
