/*
Take second example in the description:
liked list: 0->1->2->3->4
I highlighed the subset G in linked list with color red.
The problem is just to count how many red part there are.
One red part is one connected components.
To do this, we just need to count head of red parts.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	if head == nil {
		return 0
	}

	M := make(map[int]bool)
	for _, g := range G {
		M[g] = true
	}

	components := 0

	pre := head.Val
	if M[pre] {
		components++
	}

	for p := head.Next; p != nil; p = p.Next {
		cur := p.Val
		if !M[pre] && M[cur] {
			components++
		}
		pre = cur
	}

	return components
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

func main() {
	fmt.Println(numComponents(build_list([]int{0, 1, 2, 3}), []int{0, 1, 3}))
}
