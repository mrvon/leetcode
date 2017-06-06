/*
Given a string which consists of lowercase or uppercase letters, find the length
of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note: Assume the length of given string will not exceed 1,010.


The answer is very straightforword.
*/
package main

import "fmt"

func longestPalindrome(s string) int {
	dict := make(map[byte]int)

	// accumulate byte
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	// find max odd count character
	var max_byte byte
	var max_size int = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		n := dict[c]
		if n%2 != 0 {
			if n > max_size {
				max_size = n
				max_byte = c
			}
		}
	}

	// delete max odd count character
	delete(dict, max_byte)

	len := max_size
	for _, n := range dict {
		if n%2 == 0 {
			len += n
		} else {
			len += n - 1
		}
	}

	return len
}

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ab"))
	fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("acccb"))
	fmt.Println(longestPalindrome("accca"))
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("aaacccdddd"))
}
