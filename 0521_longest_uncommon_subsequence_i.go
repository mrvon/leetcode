package main

import "fmt"

func findLUSlength(a string, b string) int {
	na := len(a)
	nb := len(b)

	if nb > na {
		return nb
	} else if na > nb {
		return na
	} else {
		if a == b {
			return -1
		} else {
			return na
		}
	}
}

func main() {
	fmt.Println(findLUSlength("aba", "abaa"))
	fmt.Println(findLUSlength("aba", "cdc"))
	fmt.Println(findLUSlength("aaa", "aaa"))
	fmt.Println(findLUSlength("aaa", "aa"))
}
