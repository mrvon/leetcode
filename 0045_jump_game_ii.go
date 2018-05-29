/*
I try to change this problem to a BFS problem, where nodes in level i are all
the nodes that can be reached in i-1th jump.

for example. {2 3 1 1 4}, is

2
3 1
1 4

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

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	bound := nums[0]  // current level bound
	farest := nums[0] // farest bound for next level
	level := 1        // current level

	for i := 1; i < len(nums); i++ {
		if i > bound {
			bound = farest
			farest = max(farest, i+nums[i])
			level++
		} else {
			farest = max(farest, i+nums[i])
		}
	}

	return level
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 1, 1}))
	fmt.Println(jump([]int{4, 3, 3, 1, 1, 2, 3, 4, 3}))
}
