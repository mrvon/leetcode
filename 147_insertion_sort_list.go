package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		d := dummy
		for d.Next != nil && d.Next.Val < head.Val {
			d = d.Next
		}
		next := head.Next
		head.Next = d.Next
		d.Next = head
		head = next
	}

	return dummy.Next
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
	print_list(insertionSortList(build_list([]int{})))
	print_list(insertionSortList(build_list([]int{1, 2, 3, 4})))
	print_list(insertionSortList(build_list([]int{4, 3, 2, 1})))
	print_list(insertionSortList(build_list([]int{3, 2, 1, 4})))
}
