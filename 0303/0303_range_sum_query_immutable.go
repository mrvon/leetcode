package main

import "fmt"

type NumArray struct {
	sum []int // sum of range(i, end)
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		sum: make([]int, len(nums)),
	}

	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		na.sum[i] = s
	}

	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.sum[j] - this.sum[i-1]
	} else {
		return this.sum[j]
	}
}

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})

	fmt.Println(na.SumRange(0, 0))
	fmt.Println(na.SumRange(1, 1))
	fmt.Println(na.SumRange(2, 2))
	fmt.Println(na.SumRange(3, 3))
	fmt.Println(na.SumRange(4, 4))
	fmt.Println(na.SumRange(5, 5))
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
