package main

import (
	"fmt"
	"sort"
)

type Item struct {
	char  byte
	count int
}

type List []Item

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return l[i].count > l[j].count
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func frequencySort(s string) string {
	b := []byte(s)
	m := make(map[byte]int)

	// counting
	for i := 0; i < len(b); i++ {
		m[b[i]]++
	}

	var l List
	for n, c := range m {
		l = append(l, Item{n, c})
	}

	sort.Sort(l)

	// build result
	var r []byte
	for i := 0; i < len(l); i++ {
		for j := 0; j < l[i].count; j++ {
			r = append(r, l[i].char)
		}
	}
	return string(r)
}

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
