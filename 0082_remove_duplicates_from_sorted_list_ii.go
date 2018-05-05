package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	d := dummy

	counter := make(map[int]int)
	for p := head; p != nil; {
		counter[p.Val]++
		p = p.Next
	}

	for p := head; p != nil; {
		if counter[p.Val] == 1 {
			d.Next = p
			d = p
		}
		p = p.Next
	}
	d.Next = nil

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
	print_list(deleteDuplicates(build_list([]int{1, 2, 3, 3, 4, 4, 5})))
	print_list(deleteDuplicates(build_list([]int{1, 1, 1, 2, 3})))
	print_list(deleteDuplicates(build_list([]int{1, 2, 2})))
}
