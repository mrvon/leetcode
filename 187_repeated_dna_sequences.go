package main

import "fmt"

const (
	win_size = 10
)

func findRepeatedDnaSequences(s string) []string {
	results := []string{}

	if len(s) < win_size {
		return results
	}

	m := make(map[string]int)

	for beg, end := 0, win_size; end <= len(s); beg, end = beg+1, end+1 {
		if m[s[beg:end]] == 1 {
			results = append(results, s[beg:end])
		}
		m[s[beg:end]]++
	}

	return results
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	// ["AAAAACCCCC", "CCCCCAAAAA"].
}
