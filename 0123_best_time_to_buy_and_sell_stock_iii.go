// TODO
package main

import (
	"fmt"
	"math"
)

func mp(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min_price := math.MaxInt32
	max_profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
		} else {
			profix := prices[i] - min_price
			if profix > max_profit {
				max_profit = profix
			}
		}
	}

	return max_profit
}

func maxProfit(prices []int) int {
	max := 0

	for i := 0; i < len(prices); i++ {
		// split into two part, prices[:i], prices[i:]
		p := mp(prices[:i]) + mp(prices[i:])
		if p > max {
			max = p
		}
	}

	return max
}

func main() {
	fmt.Println(maxProfit([]int{}))
	fmt.Println(maxProfit([]int{1, 2}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
