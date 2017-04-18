package main

import (
	"fmt"
	"strconv"
)

func optimalDivision(nums []int) string {
	if len(nums) == 0 {
		return ""
	} else if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	} else if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	} else {
		right := []byte{}
		right = append(right, strconv.Itoa(nums[1])...)
		for i := 2; i < len(nums); i++ {
			right = append(right, '/')
			right = append(right, strconv.Itoa(nums[i])...)
		}
		return fmt.Sprintf("%d/(%s)", nums[0], right)
	}
}

func main() {
	nums := []int{1000, 100, 10, 2}
	fmt.Println(optimalDivision(nums))
}
