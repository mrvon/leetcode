package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 {
		return l2
	} else if l2.Val == 0 {
		return l1
	}

	n1 := []int{}
	n2 := []int{}
	r := []int{}

	for l1 != nil {
		n1 = append(n1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		n2 = append(n2, l2.Val)
		l2 = l2.Next
	}

	if len(n1) < len(n2) {
		n1, n2 = n2, n1
	}

	for i := 0; i < len(n1); i++ {
		r = append(r, n1[i])
	}
	d := len(n1) - len(n2)
	for i := d; i < len(n1); i++ {
		r[i] += n2[i-d]
	}

	for i := len(r) - 1; i > 0; i-- {
		r[i-1] += (r[i] / 10)
		r[i] %= 10
	}

	if r[0] >= 10 {
		c := r[0] / 10
		r[0] %= 10
		// prepend
		r = append(r, 0)
		copy(r[1:len(r)], r[0:len(r)-1])
		r[0] = c
	}

	return build_list(r)
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
	print_list(addTwoNumbers(
		build_list([]int{7, 2, 4, 3}),
		build_list([]int{5, 6, 4}),
	))

	print_list(addTwoNumbers(
		build_list([]int{5, 6, 4}),
		build_list([]int{7, 2, 4, 3}),
	))
	// Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	// Output: 7 -> 8 -> 0 -> 7

	print_list(addTwoNumbers(
		build_list([]int{9}),
		build_list([]int{9}),
	))
}
