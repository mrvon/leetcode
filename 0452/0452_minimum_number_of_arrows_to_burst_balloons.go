package main

import (
	"fmt"
	"sort"
)

type Balloon struct {
	s int
	e int
}

func findMinArrowShots(points [][]int) int {
	bs := []Balloon{}
	be := []Balloon{}
	bm := make(map[Balloon]bool)

	for i := 0; i < len(points); i++ {
		b := Balloon{points[i][0], points[i][1]}
		bs = append(bs, b)
		be = append(be, b)
		bm[b] = true
	}

	// Sorted by start
	sort.Slice(bs, func(i int, j int) bool {
		return bs[i].s < bs[j].s
	})

	// Sorted by end
	sort.Slice(be, func(i int, j int) bool {
		return be[i].e < be[j].e
	})

	count := 0

	for len(bs) > 0 {
		end_item := be[0]
		be = be[1:]
		if !bm[end_item] {
			continue
		}
		for len(bs) > 0 {
			start_item := bs[0]
			if start_item.s > end_item.e {
				break
			}
			delete(bm, start_item)
			bs = bs[1:]
		}
		count++
	}

	return count
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
