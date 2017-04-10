// exact like problem 205
package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	word_list := strings.Split(str, " ")
	mapping := make(map[byte]string)
	used := make(map[string]bool)

	if len(pattern) != len(word_list) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if i >= len(word_list) {
			return false
		}

		if mapping[pattern[i]] == "" {
			if used[word_list[i]] {
				return false
			}
			used[word_list[i]] = true
			mapping[pattern[i]] = word_list[i]
		} else {
			if mapping[pattern[i]] != word_list[i] {
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
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("jquery", "jquery"))
	fmt.Println(wordPattern("aaaa", "aa aa aa"))
	fmt.Println(wordPattern("aaa", "aa aa aa aa"))
}
