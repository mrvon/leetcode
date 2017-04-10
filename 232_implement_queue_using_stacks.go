package main

import "fmt"

// ----------------------------------------------------------------------------

type Stack struct {
	S []int
}

func (S *Stack) len() int {
	return len(S.S)
}

func (S *Stack) push(t int) {
	S.S = append(S.S, t)
}

func (S *Stack) pop() (t int) {
	t = S.S[len(S.S)-1]
	S.S = S.S[:len(S.S)-1]
	return
}

func (S *Stack) peek() (t int) {
	t = S.S[len(S.S)-1]
	return
}

// ----------------------------------------------------------------------------

type MyQueue struct {
	s1 Stack
	s2 Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s2.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s1.len() == 0 {
		for this.s2.len() > 0 {
			this.s1.push(this.s2.pop())
		}
	}
	return this.s1.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.s1.len() == 0 {
		for this.s2.len() > 0 {
			this.s1.push(this.s2.pop())
		}
	}
	return this.s1.peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.len() == 0 && this.s2.len() == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Empty())
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
