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

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := List{}

	for i, n := range nums {
		l = append(l, Item{
			index: i,
			value: n,
		})
	}

	sort.Sort(l)

	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[j].value-l[i].value > t {
				break
			}
			if abs(l[i].index-l[j].index) <= k {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{0, 10, 22, 15, 0, 5, 22, 12, 1, 5}, 3, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 3, 6, 2}, 1, 2))
}
