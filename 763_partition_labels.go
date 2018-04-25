package main

import (
	"fmt"
	"sort"
)

type Range struct {
	c byte // character
	l int  // left range
	r int  // right range
}

type RangeList []*Range
type ByLeft RangeList
type ByRight RangeList

func (l ByLeft) Len() int {
	return len(l)
}

func (l ByLeft) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByLeft) Less(i int, j int) bool {
	return l[i].l < l[j].l
}

func (l ByRight) Len() int {
	return len(l)
}

func (l ByRight) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByRight) Less(i int, j int) bool {
	return l[i].r < l[j].r
}

func merge(r1 *Range, r2 *Range) bool {
	if r1.r < r2.l || r1.l > r2.r {
		return false
	}
	if r2.l < r1.l {
		r1.l = r2.l
	}
	if r2.r > r1.r {
		r1.r = r2.r
	}
	return true
}

func partitionLabels(S string) []int {
	dist := make(map[byte]*Range)

	for i, s := range S {
		c := byte(s)
		if r, has := dist[c]; has {
			r.r = i
		} else {
			r := &Range{
				c: c,
				l: i,
				r: i,
			}
			dist[c] = r
		}
	}

	ll := RangeList{}
	lr := RangeList{}

	for _, r := range dist {
		ll = append(ll, r)
		lr = append(lr, r)
	}

	sort.Sort(ByLeft(ll))
	sort.Sort(ByRight(lr))

	parts := []int{}
	for _, r1 := range lr {
		if _, has := dist[r1.c]; !has {
			continue
		}
		delete(dist, r1.c)
		for _, r2 := range ll {
			if _, has := dist[r2.c]; !has {
				continue
			}
			if merge(r1, r2) {
				delete(dist, r2.c)
			} else {
				break
			}
		}
		parts = append(parts, r1.r-r1.l+1)
	}
	return parts
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("vhaagbqkaq"))
}
