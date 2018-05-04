/*
Recursive with memoized, 	6ms
Recursive without memoized, 139ms
*/
package main

import "fmt"

type R struct {
	A int
	B int
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func predict(nums []int, s int, e int, memoized map[R]R) (A int, B int) {
	if r, existed := memoized[R{s, e}]; existed {
		A = r.A
		B = r.B
		return
	}

	l := e - s + 1
	if l <= 0 {
		A = 0
		B = 0
		return
	} else if l == 1 {
		A = nums[s]
		B = 0
		return
	} else if l == 2 {
		A = max(nums[s], nums[e])
		B = min(nums[s], nums[e])
		return
	}

	// l >= 3
	A1, B1 := predict(nums, s+1, e, memoized)
	A2, B2 := predict(nums, s, e-1, memoized)

	A31 := nums[s] + B1
	A32 := nums[e] + B2

	if A31 >= A32 {
		A = A31
		B = A1
	} else {
		A = A32
		B = A2
	}

	memoized[R{s, e}] = R{A, B}
	return
}

func PredictTheWinner(nums []int) bool {
	A, B := predict(nums, 0, len(nums)-1, make(map[R]R))
	return A >= B
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(true, PredictTheWinner([]int{0}))
	assert(true, PredictTheWinner([]int{1}))
	assert(true, PredictTheWinner([]int{1, 2}))
	assert(false, PredictTheWinner([]int{1, 5, 2}))
	assert(true, PredictTheWinner([]int{1, 5, 233, 7}))
	assert(true, PredictTheWinner([]int{1, 5, 2, 4}))
}
