// Naive solution, use min-heap will be better.
package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	s   []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if x < this.min {
		this.min = x
	}

}

func (this *MinStack) Pop() {
	this.s = this.s[0 : len(this.s)-1]
	this.min = math.MaxInt32
	for i := 0; i < len(this.s); i++ {
		if this.s[i] < this.min {
			this.min = this.s[i]
		}
	}
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(0)
	fmt.Println("MIN", s.GetMin())
	fmt.Println("TOP", s.Top())
	s.Pop()
	fmt.Println("MIN", s.GetMin())
	fmt.Println("TOP", s.Top())
	s.Pop()
	fmt.Println("TOP", s.Top())
	s.Pop()
	fmt.Println("TOP", s.Top())
	s.Pop()
}
