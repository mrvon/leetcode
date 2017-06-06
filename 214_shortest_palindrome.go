/*
Given a string S, you are allowed to convert it to a palindrome by adding
characters in front of it. Find and return the shortest palindrome you can find
by performing this transformation.

For example:

Given "aacecaaa", return "aaacecaaa".
Given "abcd", return "dcbabcd".

--------------------------------------------------------------------------------
Key Point:
	Find the longest prefix of given string which is palindrome

*/
package main

import "fmt"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		n = append(n, s[i])
	}
	return string(n)
}

func shortestPalindrome(s string) string {
	for i := len(s); i > 0; i-- {
		ss := s[:i]
		if isPalindrome(ss) {
			return reverse(s[i:]) + s
		}
	}
	// Never be here
	return ""
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("aacecakk"))
}
