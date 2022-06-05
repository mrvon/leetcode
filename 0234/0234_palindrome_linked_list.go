package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-1-i] {
			return false
		}
	}
	return true
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
	fmt.Println(isPalindrome(build_list([]int{})))
	fmt.Println(isPalindrome(build_list([]int{3, 2, 3})))
	fmt.Println(isPalindrome(build_list([]int{1, 2, 3})))
	fmt.Println(isPalindrome(build_list([]int{1, 2, 3, 4})))
}
