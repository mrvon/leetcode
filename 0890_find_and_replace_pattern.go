// Translate to numeric representation
// abc -> 123
// abb -> 122
// ccc -> 111
package main

import "fmt"

func translate(cache map[string]string, word string) string {
	if nw, ok := cache[word]; ok {
		return nw
	}
	r := []rune{}
	m := make(map[rune]int)
	c := 0
	for _, w := range word {
		if m[w] == 0 {
			c++
			m[w] = c
		}
		r = append(r, rune('0'+m[w]))
	}
	nw := string(r)
	cache[word] = nw
	return nw
}

func findAndReplacePattern(words []string, pattern string) []string {
	cache := make(map[string]string)
	matchs := []string{}
	np := translate(cache, pattern)
	for _, word := range words {
		nw := translate(cache, word)
		if nw == np {
			matchs = append(matchs, word)
		}
	}
	return matchs
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
