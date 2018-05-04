package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
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

	// Push internal before newInterval, their wouldn't overlap with
	// newInterval.
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i].End < newInterval.Start {
			result = append(result, intervals[i])
		} else {
			break
		}
	}

	// Merge all interval with newInterval
	for ; i < len(intervals); i++ {
		if overlap(intervals[i], newInterval) {
			newInterval = merge(intervals[i], newInterval)
		} else {
			break
		}
	}
	result = append(result, newInterval)

	// Push others after the newInterval, their wouldn't overlap with
	// newInterval.
	for ; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}

	return result
}

func main() {
	intervals := []Interval{
		Interval{1, 3},
		Interval{6, 9},
	}
	new_interval := Interval{2, 5}

	fmt.Println(insert(intervals, new_interval))
}

// another implementation using problem 56's merge function
// that's very easy.
/*
func insert(intervals []Interval, newInterval Interval) []Interval {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}
*/
