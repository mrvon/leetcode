package main

import (
	"fmt"
	"sort"
)

func build(s string) map[byte]int {
	dict := make(map[byte]int)
	for _, r := range s {
		b := byte(r)
		if b >= 'A' && b <= 'Z' {
			dict[b+'a'-'A']++
		} else if b >= 'a' && b <= 'z' {
			dict[b]++
		}
	}
	return dict
}

func shortestCompletingWord(licensePlate string, words []string) string {
	sort.SliceStable(words, func(i int, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dict := build(licensePlate)

	for _, w := range words {
		ok := true
		d := build(w)
		for r, c := range dict {
			if d[r] < c {
				ok = false
				break
			}
		}
		if ok {
			return w
		}
	}

	// no answer
	return ""
}

func main() {
	// steps
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	// pest
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	// husband
	fmt.Println(shortestCompletingWord("Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}))
}
