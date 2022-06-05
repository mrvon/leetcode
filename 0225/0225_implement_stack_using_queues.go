package main

import "fmt"

// ----------------------------------------------------------------------------

type Queue struct {
	Q []int
}

func (Q *Queue) len() int {
	return len(Q.Q)
}

func (Q *Queue) push(t int) {
	Q.Q = append(Q.Q, t)
}

func (Q *Queue) pop() (t int) {
	t = Q.Q[0]
	Q.Q = Q.Q[1:len(Q.Q)]
	return
}

func (Q *Queue) peek() (t int) {
	t = Q.Q[len(Q.Q)-1]
	return
}

// ----------------------------------------------------------------------------

type MyStack struct {
	q1 *Queue
	q2 *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: &Queue{},
		q2: &Queue{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1.push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	top := 0
	for this.q1.len() > 0 {
		top = this.q1.pop()
		if this.q1.len() > 0 {
			this.q2.push(top)
		}
	}
	this.q1, this.q2 = this.q2, this.q1
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
	top := 0
	for this.q1.len() > 0 {
		top = this.q1.pop()
		this.q2.push(top)
	}
	this.q1, this.q2 = this.q2, this.q1
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.len() == 0
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Empty())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
