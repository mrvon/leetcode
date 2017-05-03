package main

import (
	"fmt"
	"sort"
)

type Item struct {
	b byte
	i int
}

type ByItem []Item

func (b ByItem) Len() int {
	return len(b)
}

func (b ByItem) Less(i, j int) bool {
	return b[i].i < b[j].i
}

func (b ByItem) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type Stack []Item

func (S *Stack) len() int {
	return len(*S)
}

func (S *Stack) push(item Item) {
	(*S) = append(*S, item)
}

func (S *Stack) pop() Item {
	item := (*S)[len(*S)-1]
	(*S) = (*S)[:len(*S)-1]
	return item
}

func (S *Stack) peek() Item {
	return (*S)[len(*S)-1]
}

func removeKdigits(num string, k int) string {
	s := Stack{}
	l := []Item{}

	for i := 0; i < len(num); {
		item := Item{num[i], i}

		if s.len() == 0 {
			s.push(item)
			i++
		} else {
			if item.b >= s.peek().b {
				s.push(item)
				i++
			} else {
				l = append(l, s.pop())
			}
		}
	}

	for s.len() > 0 {
		l = append(l, s.pop())
	}

	if k >= len(l) {
		return "0"
	}

	l = l[k:]

	sort.Sort(ByItem(l))

	b := []byte{}
	for i := 0; i < len(l); i++ {
		if len(b) == 0 && l[i].b == '0' {
			continue
		}
		b = append(b, l[i].b)
	}

	if len(b) == 0 {
		return "0"
	}
	return string(b)
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("010", 2))
	fmt.Println(removeKdigits("112", 1))
}
