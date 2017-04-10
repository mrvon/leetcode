// Slide window
package main

import "fmt"

func findAnagrams(s string, p string) []int {
	anagrams := []int{}

	// need characters
	working := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		working[p[i]]--
	}

	for i := 0; i < len(s); i++ {
		if i >= len(p) {
			b := s[i-len(p)]
			working[b]--
			if working[b] == 0 {
				delete(working, b)
			}
		}

		working[s[i]]++
		if working[s[i]] == 0 {
			delete(working, s[i])
		}

		if len(working) == 0 {
			anagrams = append(anagrams, i-len(p)+1)
		}
	}

	return anagrams
}

func main() {
	fmt.Println(findAnagrams("abcdeabc", "abc"))
}
