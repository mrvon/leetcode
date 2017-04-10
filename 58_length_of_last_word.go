package main

import "fmt"

func lengthOfLastWord(s string) int {
	is_enter_word := false
	word_len := 0

	for i := len(s) - 1; i >= 0; i-- {
		if is_enter_word {
			if s[i] == ' ' {
				break
			} else {
				word_len++
			}
		} else {
			if s[i] != ' ' {
				is_enter_word = true
				word_len++
			}
		}
	}

	return word_len
}

func main() {
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("hello"))
	fmt.Println(lengthOfLastWord("hello world"))
}
