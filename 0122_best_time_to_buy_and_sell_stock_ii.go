package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(0, maxProfit([]int{0}))
	assert(7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert(6, maxProfit([]int{1, 2, 3, 4, 5, 7}))
	assert(0, maxProfit([]int{7, 6, 5, 4, 3, 2}))
	assert(2, maxProfit([]int{2, 1, 2, 0, 1}))
	assert(3, maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
	assert(11, maxProfit([]int{3, 4, 2, 1, 2, 0, 1, 7, 8, 0, 1}))
}
