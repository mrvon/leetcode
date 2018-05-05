package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	// create a dummy node to mark the head of this list
	dummy := &ListNode{0, head}
	prev := dummy

	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}

	start := prev.Next
	then := start.Next

	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
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
	print_list(reverseBetween(build_list([]int{1, 2, 3, 4, 5}), 2, 4))
	print_list(reverseBetween(build_list([]int{1, 2, 3, 4, 5}), 1, 1))
	print_list(reverseBetween(build_list([]int{1, 2, 3, 4, 5}), 4, 4))
	print_list(reverseBetween(build_list([]int{1, 2, 3, 4, 5}), 1, 5))
}
