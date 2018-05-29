package main

import "fmt"

func findWords(words []string) []string {
	mapping := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

	index := func(b byte) int {
		if b >= 'A' && b <= 'Z' { // to lowercase
			b += 'a' - 'A'
		}
		for i := 0; i < len(mapping); i++ {
			m := mapping[i]
			for j := 0; j < len(m); j++ {
				if m[j] == b {
					return i
				}
			}
		}
		// never be here
		return 0
	}

	filter := func(word string) bool {
		if len(word) == 0 {
			return true
		}
		first := index(word[0])

		for i := 1; i < len(word); i++ {
			if index(word[i]) != first {
				return false
			}
		}
		return true
	}

	var list []string

	for i := 0; i < len(words); i++ {
		if filter(words[i]) {
			list = append(list, words[i])
		}
	}

	return list
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}
