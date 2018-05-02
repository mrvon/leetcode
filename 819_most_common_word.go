package main

import (
	"fmt"
	"strings"
)

func trim(word string) string {
	start := 0
	end := len(word)
	for start < end {
		if word[start] >= 'a' && word[start] <= 'z' {
			break
		}
		if word[start] >= 'A' && word[start] <= 'Z' {
			break
		}
		start++
	}
	for end > 0 {
		if word[end-1] >= 'a' && word[end-1] <= 'z' {
			break
		}
		if word[end-1] >= 'A' && word[end-1] <= 'Z' {
			break
		}
		end--
	}

	return word[start:end]
}

func mostCommonWord(paragraph string, banned []string) string {
	bm := make(map[string]bool)
	fq := make(map[string]int)

	for _, w := range banned {
		bm[w] = true
	}

	words := strings.Fields(paragraph)
	for _, w := range words {
		t := strings.ToLower(trim(w))
		if bm[t] {
			continue
		}
		fq[t]++
	}

	max_word := ""
	max_count := 0

	for w, c := range fq {
		if max_count == 0 {
			max_word = w
			max_count = c
		} else {
			if c > max_count {
				max_word = w
				max_count = c
			}
		}
	}

	return max_word
}

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}
