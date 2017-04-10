/* Dynamic Programming

Bottom up

optimal[i] = max(
	optimal[i-2]+nums[i-1],
	optimal[i-1],
)

*/
package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	optimal := make([]int, n+1)

	optimal[0] = 0
	optimal[1] = nums[0]

	for i := 2; i <= n; i++ {
		optimal[i] = max(
			optimal[i-2]+nums[i-1],
			optimal[i-1],
		)
	}

	return optimal[n]
}

func main() {
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{7}))
	fmt.Println(rob([]int{7, 2, 1}))
	fmt.Println(rob([]int{1, 3, 1}))
}
