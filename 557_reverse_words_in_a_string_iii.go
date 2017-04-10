package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	reverse := func(s string) string {
		b := []byte(s)
		l := 0
		r := len(b) - 1
		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
		return string(b)
	}

	words := strings.Split(s, " ")
	rwords := []string{}
	for i := 0; i < len(words); i++ {
		rwords = append(rwords, reverse(words[i]))
	}
	return strings.Join(rwords, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
