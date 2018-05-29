package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	max := []int{
		math.MinInt32 - 1, // tricky
		math.MinInt32 - 1,
		math.MinInt32 - 1,
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(max); j++ {
			if nums[i] == max[j] {
				break
			} else if nums[i] < max[j] {
				continue
			} else {
				for k := len(max) - 1; k > j; k-- {
					max[k] = max[k-1]
				}
				max[j] = nums[i]
				break
			}
		}
	}

	if max[len(max)-1] == math.MinInt32-1 {
		return max[0]
	} else {
		return max[len(max)-1]
	}
}

func main() {
	fmt.Println(thirdMax([]int{1}))
	fmt.Println(thirdMax([]int{1, 2, 3}))
	fmt.Println(thirdMax([]int{1, 2, 3, 1}))
	fmt.Println(thirdMax([]int{1, 2, 3, 4}))
	fmt.Println(thirdMax([]int{1, 2, -2147483648}))
}
