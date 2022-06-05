package main

import "fmt"

func anagram(s1 string, s2 string) bool {
	// len(s1) == len(s2)
	counter := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		counter[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		if counter[s2[i]] <= 0 {
			return false
		}
		counter[s2[i]]--
	}
	return true
}

// Top-Down Recursive solution, Accepted
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		// Ok, All done!
		return true
	}
	if len(s1) != len(s2) {
		// Needn't try!
		return false
	}
	if !anagram(s1, s2) {
		// Needn't try!
		return false
	}
	n := len(s1)
	// Split into two parts
	for i := 1; i < n; i++ {
		// Remain not Exchange
		if isScramble(s1[:i], s2[:i]) &&
			isScramble(s1[i:], s2[i:]) {
			return true
		}
		// Exchange two children
		if isScramble(s1[:i], s2[n-i:]) &&
			isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isScramble("t", "t"))
	fmt.Println(isScramble("eat", "tae"))
	fmt.Println(isScramble("great", "rgtae"))
	fmt.Println(isScramble("great", "rtgae"))
	fmt.Println(isScramble("ccabcbabcbabbbbcbb", "bbbbabccccbbbabcba"))
}
