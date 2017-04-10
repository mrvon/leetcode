package main

import "fmt"

func reverse(s string) []byte {
	b := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}
	return b
}

func reverseStr(s string, k int) string {
	n := []byte{}
	i := 0
	r := true
	for {
		if i+k >= len(s) {
			// not enough
			if r {
				n = append(n, reverse(s[i:len(s)])...)
			} else {
				n = append(n, s[i:len(s)]...)
			}
			break
		} else {
			// enough
			if r {
				n = append(n, reverse(s[i:i+k])...)
			} else {
				n = append(n, s[i:i+k]...)
			}
			i += k
		}
		r = !r
	}
	return string(n)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	// Input: s = "abcdefg", k = 2
	// Output: "bacdfeg"
}
