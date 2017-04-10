/*
Sort the list by interval.Start, we take the first interval, and check whether
it overlap with next interval. If so, we merge them together, otherwise push
interval into result and take next interval.
*/
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type IntervalList []Interval

func (l IntervalList) Len() int {
	return len(l)
}

func (l IntervalList) Less(i, j int) bool {
	return l[i].Start <= l[j].Start
}

func (l IntervalList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func merge(intervals []Interval) []Interval {
	overlap := func(i1, i2 Interval) bool {
		if i1.Start > i2.End || i2.Start > i1.End {
			return false
		} else {
			return true
		}
	}

	merge := func(i1, i2 Interval) Interval {
		min := func(x, y int) int {
			if x <= y {
				return x
			} else {
				return y
			}
		}

		max := func(x, y int) int {
			if x >= y {
				return x
			} else {
				return y
			}
		}
		return Interval{min(i1.Start, i2.Start), max(i1.End, i2.End)}
	}

	var result []Interval

	// sort by Interval.Start
	sort.Sort(IntervalList(intervals))

	var s *Interval = nil

	for i := 0; i < len(intervals); i++ {
		if s == nil {
			s = &intervals[i]
			continue
		} else {
			if overlap(*s, intervals[i]) {
				*s = merge(*s, intervals[i])
			} else {
				result = append(result, *s)
				s = &intervals[i]
			}
		}
	}

	if s != nil {
		result = append(result, *s)
	}

	return result
}

func main() {
	intervals := []Interval{
		Interval{8, 10},
		Interval{10, 11},
		Interval{1, 3},
		Interval{2, 6},
		Interval{15, 18},
	}
	fmt.Println(merge(intervals))
}
