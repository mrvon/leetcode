package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func canJump(nums []int) bool {
	can_reach := func(sp, tp Interval) bool {
		mid := tp.Start + (tp.End-tp.Start)/2
		if mid >= sp.Start && mid <= sp.End {
			return true
		} else {
			return false
		}
	}

	merge := func(sp, tp Interval) Interval {
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

		return Interval{min(sp.Start, tp.Start), max(sp.End, tp.End)}
	}

	var intervals []Interval

	for i := 0; i < len(nums); i++ {
		intervals = append(intervals, Interval{i - nums[i], i + nums[i]})
	}

	if len(intervals) == 0 {
		return true
	} else if len(intervals) == 1 {
		return true
	} else {
		start := intervals[0]
		for i := 1; i < len(intervals); i++ {
			if can_reach(start, intervals[i]) {
				start = merge(start, intervals[i])
			} else {
				return false
			}
		}
	}

	return true
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(false, canJump([]int{0, 2}))
	assert(true, canJump([]int{2, 3, 1, 1, 4}))
	assert(false, canJump([]int{3, 2, 1, 0, 4}))
}
