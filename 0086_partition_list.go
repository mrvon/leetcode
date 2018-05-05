// Damn smart using dummy node
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy_less := &ListNode{0, nil}
	dummy_great := &ListNode{0, nil}

	less := dummy_less
	great := dummy_great

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			great.Next = head
			great = great.Next
		}
		head = head.Next
	}

	less.Next = dummy_great.Next
	great.Next = nil

	return dummy_less.Next
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
	print_list(partition(build_list([]int{1, 4, 3, 2, 5, 2}), 3))
	print_list(partition(build_list([]int{1, 4, 3, 2, 5, 2}), 1))
	print_list(partition(build_list([]int{1, 4, 3, 2, 5, 2}), 0))
	print_list(partition(build_list([]int{1, 4, 3, 2, 5, 2}), 6))

	print_list(partition(build_list([]int{1, 1}), 2))
	print_list(partition(build_list([]int{1, 1}), 0))
	print_list(partition(build_list([]int{2, 1}), 2))
}
