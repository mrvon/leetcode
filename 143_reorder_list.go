package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	l1 := []*ListNode{}

	for head != nil {
		l1 = append(l1, head)
		head = head.Next
	}

	l2 := make([]*ListNode, len(l1))

	j := 0
	for i := 0; i < len(l2); i += 2 {
		l2[i] = l1[j]
		j++
	}

	j = len(l1) - 1
	for i := 1; i < len(l2); i += 2 {
		l2[i] = l1[j]
		j--
	}

	for i := 0; i < len(l2); i++ {
		if i == len(l2)-1 {
			l2[i].Next = nil
		} else {
			l2[i].Next = l2[i+1]
		}
	}
}

func build_list(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var l *ListNode = &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	p := l
	for i := 1; i < len(nums); i++ {
		n := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		p.Next = n
		p = n
	}

	return l
}

func print_list(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Print("\n")
}

func main() {
	l := build_list([]int{1, 2, 3, 4})
	reorderList(l)
	print_list(l)
}
