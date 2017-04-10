/* merge sort for linked list */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(first, second *ListNode) *ListNode {
	dummy := &ListNode{}
	d := dummy

	for first != nil || second != nil {
		if first == nil {
			d.Next = second
			second = second.Next
		} else if second == nil {
			d.Next = first
			first = first.Next
		} else {
			if first.Val <= second.Val {
				d.Next = first
				first = first.Next
			} else {
				d.Next = second
				second = second.Next
			}
		}
		d = d.Next
	}

	d.Next = nil

	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	if count <= 1 {
		return head
	}

	first := head
	second := head

	// cut into two part
	prev := head

	for i := 0; i < count/2; i++ {
		prev = second
		second = second.Next
	}
	prev.Next = nil

	first = sortList(first)
	second = sortList(second)
	return merge(first, second)
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
	print_list(sortList(build_list([]int{})))
	print_list(sortList(build_list([]int{3, 2, 4, 1})))
	print_list(sortList(build_list([]int{4, 3, 2, 1})))
}
