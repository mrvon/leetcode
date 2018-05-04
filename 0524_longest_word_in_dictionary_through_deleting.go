package main

import (
	"fmt"
	"sort"
)

func isValid(s string, item string) bool {
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == item[i] {
			i++
			if i >= len(item) {
				return true
			}
		}
	}
	return false
}

func findLongestWord(s string, d []string) string {
	s_map := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		s_map[s[i]]++
	}

	longest_length := 0
	longest_list := []string{}

	for i := 0; i < len(d); i++ {
		item := d[i]

		if !isValid(s, item) {
			continue
		}

		if len(item) > longest_length {
			longest_length = len(item)
			longest_list = []string{}
			longest_list = append(longest_list, item)
		} else if len(item) == longest_length {
			longest_list = append(longest_list, item)
		}
	}

	if len(longest_list) > 0 {
		sort.Strings(longest_list)
		return longest_list[0]
	} else {
		return ""
	}
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}
