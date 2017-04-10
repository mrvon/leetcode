package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}

	prev := dummy
	end := dummy.Next
	for {
		start := end
		count := 0
		for end != nil && count < k {
			end = end.Next
			count++
		}

		if count >= k {
			// enough, just reverse it.
			then := start.Next
			for i := 0; i < k-1; i++ {
				start.Next = then.Next
				then.Next = prev.Next
				prev.Next = then
				then = start.Next
			}
			for i := 0; i < k; i++ {
				prev = prev.Next
			}
		} else {
			// not enough
			break
		}
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
	print_list(reverseKGroup(build_list([]int{1, 2}), 1))
	print_list(reverseKGroup(build_list([]int{1, 2}), 2))
	print_list(reverseKGroup(build_list([]int{1, 2, 3}), 1))
	print_list(reverseKGroup(build_list([]int{1, 2, 3}), 2))
	print_list(reverseKGroup(build_list([]int{1, 2, 3}), 3))
	print_list(reverseKGroup(build_list([]int{1, 2, 3, 4}), 2))
	print_list(reverseKGroup(build_list([]int{1, 2, 3, 4}), 3))
	print_list(reverseKGroup(build_list([]int{1, 2, 3, 4, 5}), 2))
}
