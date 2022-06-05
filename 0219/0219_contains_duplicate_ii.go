package main

import (
	"fmt"
	"sort"
)

type Item struct {
	index int // original index
	value int
}

type List []Item

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return l[i].index < l[j].index
	} else {
		return l[i].value < l[j].value
	}
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func containsNearbyDuplicate(nums []int, k int) bool {
	l := List{}

	for i, n := range nums {
		l = append(l, Item{
			index: i,
			value: n,
		})
	}

	sort.Sort(l)

	for i := 1; i < len(l); i++ {
		if l[i].value == l[i-1].value && l[i].index-l[i-1].index <= k {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}
