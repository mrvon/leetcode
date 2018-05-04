// top-down memoization
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	memoized := make([]int, amount+1)
	return change(coins, amount, memoized)
}

func change(coins []int, amount int, memoized []int) int {
	if memoized[amount] != 0 {
		return memoized[amount]
	}

	min := math.MaxInt32

	for i := 0; i < len(coins); i++ {
		if amount < coins[i] {
			continue
		}

		if amount == coins[i] {
			if min > 1 {
				min = 1
			}
		} else { // amount > coins[i]
			s := change(coins, amount-coins[i], memoized)
			if s >= 0 && s+1 < min {
				min = s + 1
			}
		}
	}

	if min == math.MaxInt32 {
		min = -1
	}

	memoized[amount] = min
	return min
}

func main() {
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}
