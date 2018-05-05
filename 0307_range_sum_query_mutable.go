// TODO, another optimal solution?
package main

import "fmt"

type NumArray struct {
	num []int
	sum []int // sum of range(i, end)
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		num: nums,
		sum: make([]int, len(nums)),
	}

	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		na.sum[i] = s
	}

	return na
}

func (this *NumArray) Update(i int, val int) {
	diff := val - this.num[i]
	for j := i; j < len(this.sum); j++ {
		this.sum[j] += diff
	}
	this.num[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.sum[j] - this.sum[i-1]
	} else {
		return this.sum[j]
	}
}

func main() {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))
}
