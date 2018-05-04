package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	p := head
	length := 1
	for ; p.Next != nil; p = p.Next {
		length++
	}
	p.Next = head // make the list as circle
	tail := p

	k = k % length

	// need rotate right k, we rotate left length-k instead.
	for i := 0; i < length-k; i++ {
		tail = head
		head = head.Next
	}
	tail.Next = nil // break circle

	return head
}

func print(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println("")
}

func build() *ListNode {
	// build singular linked list
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	return node1
}

func main() {
	print(rotateRight(build(), 0))
	print(rotateRight(build(), 1))
	print(rotateRight(build(), 2))
	print(rotateRight(build(), 3))
	print(rotateRight(build(), 4))
	print(rotateRight(build(), 5))
	print(rotateRight(build(), 6))
}
