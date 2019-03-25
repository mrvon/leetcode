package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(S string) string {
	T := []rune(S)
	low := 0
	high := len(T) - 1
	for {
		for low < high && !unicode.IsLetter(T[low]) {
			low++
		}
		for low < high && !unicode.IsLetter(T[high]) {
			high--
		}
		if low >= high {
			break
		}
		T[low], T[high] = T[high], T[low]
		low++
		high--
	}
	return string(T)
}

func assert(result bool) {
	if !result {
		panic(fmt.Sprintf("Assert failed!"))
	}
}

func main() {
	assert(reverseOnlyLetters("Test1ng-Leet=code-Q!") == "Qedo1ct-eeLg=ntse-T!")
}
