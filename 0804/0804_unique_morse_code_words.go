package main

import "fmt"

var MorseCodes = []string{
	".-",
	"-...",
	"-.-.",
	"-..",
	".",
	"..-.",
	"--.",
	"....",
	"..",
	".---",
	"-.-",
	".-..",
	"--",
	"-.",
	"---",
	".--.",
	"--.-",
	".-.",
	"...",
	"-",
	"..-",
	"...-",
	".--",
	"-..-",
	"-.--",
	"--..",
}

func uniqueMorseRepresentations(words []string) int {
	unique := make(map[string]bool)
	count := 0
	for _, w := range words {
		s := ""
		for _, r := range w {
			s += MorseCodes[r-rune('a')]
		}
		if !unique[s] {
			unique[s] = true
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
