package main

import "fmt"

/*
Use a sliding window s[i:j], always add the new character, and remove the first
window character if the extension isn't ok. So in each step, either extend the
window by one or move it by one.
*/

func maxElement(counter map[byte]int) int {
	var max_count int

	for _, c := range counter {
		if c > max_count {
			max_count = c
		}
	}

	return max_count
}

func characterReplacement(s string, k int) int {
	left := 0
	right := 0
	counter := make(map[byte]int)
	for right < len(s) {
		counter[s[right]]++
		right++
		if right-left-maxElement(counter) > k {
			counter[s[left]]--
			left++
		}
	}
	return right - left
}

func main() {
	fmt.Println(characterReplacement("A", 2))
	fmt.Println(characterReplacement("ABAB", 0))
	fmt.Println(characterReplacement("ABAB", 1))
	fmt.Println(characterReplacement("ABAB", 2))
}
