package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	return buildBST(list, 0, len(list)-1)
}

func buildBST(list []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	return &TreeNode{
		Val:   list[mid],
		Left:  buildBST(list, left, mid-1),
		Right: buildBST(list, mid+1, right),
	}
}
