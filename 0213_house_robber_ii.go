/* Dynamic Programming

Bottom up

optimal[i] = max(
	exclude[i-2]+nums[i-1],
	include[i-1],
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

	include := make([]int, n+1) // include first element
	exclude := make([]int, n+1) // exclude first element
	optimal := make([]int, n+1)

	include[0] = 0
	include[1] = nums[0]
	for i := 2; i <= n; i++ {
		include[i] = max(
			include[i-2]+nums[i-1],
			include[i-1],
		)
	}

	exclude[0] = 0
	exclude[1] = 0
	for i := 2; i <= n; i++ {
		exclude[i] = max(
			exclude[i-2]+nums[i-1],
			exclude[i-1],
		)
	}

	optimal[0] = 0
	optimal[1] = nums[0]
	for i := 2; i <= n; i++ {
		optimal[i] = max(
			exclude[i-2]+nums[i-1], // select last element, so we exclude first element
			include[i-1],           // discard last element, so we include first element
		)
	}

	return optimal[n]
}

func main() {
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{7}))
	fmt.Println(rob([]int{7, 2, 1}))
	fmt.Println(rob([]int{1, 3, 1}))
	fmt.Println(rob([]int{2, 1, 1, 1}))
}
