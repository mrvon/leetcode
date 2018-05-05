package main

import (
	"fmt"
	"sort"
)

type SortedBytes []byte

func (s SortedBytes) Len() int {
	return len(s)
}

func (s SortedBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortedBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func generate_index() func(word string) int {
	m := make(map[string]int)
	g := 0
	return func(word string) int {
		sb := SortedBytes(word)
		sort.Sort(sb)
		s := string(sb)
		_, b := m[s]
		if !b {
			m[s] = g
			g++
		}
		return m[s]
	}
}

func groupAnagrams(strs []string) [][]string {
	var r [][]string
	index := generate_index()
	for j := 0; j < len(strs); j++ {
		str := strs[j]
		i := index(str)
		if i >= len(r) {
			r = append(r, []string{})
		}
		r[i] = append(r[i], str)
	}
	return r
}

func main() {
	t := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(t))
}
