package main

import "fmt"

func customSortString(S string, T string) string {
	s := make(map[byte]int)

	for _, r := range T {
		c := byte(r)
		s[c]++
	}

	result := []byte{}

	for _, r := range S {
		c := byte(r)
		for i := 0; i < s[c]; i++ {
			result = append(result, c)
		}
		delete(s, c)
	}

	for c, n := range s {
		for i := 0; i < n; i++ {
			result = append(result, c)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(customSortString("cba", "abcd"))
}
