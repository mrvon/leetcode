package main

import (
	"fmt"
	"sort"
)

type Balloon struct {
	s int
	e int
}

type ByStart []Balloon

func (b ByStart) Len() int {
	return len(b)
}

func (b ByStart) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByStart) Less(i, j int) bool {
	return b[i].s < b[j].s
}

type ByEnd []Balloon

func (b ByEnd) Len() int {
	return len(b)
}

func (b ByEnd) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByEnd) Less(i, j int) bool {
	return b[i].e < b[j].e
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

	sort.Sort(ByStart(bs))
	sort.Sort(ByEnd(be))

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
