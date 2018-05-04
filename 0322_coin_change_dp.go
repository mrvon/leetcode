// bottom-up dynamic programming
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	change := make([]int, amount+1)

	change[0] = 0

	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] { // can change?
				s := change[i-coins[j]]
				if s != -1 { // have valid solution?
					if s+1 < min {
						min = s + 1
					}
				}
			}
		}
		if min == math.MaxInt32 { // can't found solution
			min = -1
		}
		change[i] = min
	}

	return change[amount]
}

func main() {
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}
