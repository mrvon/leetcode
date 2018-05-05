package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traversal(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left, list)
	(*list) = append((*list), root)
	traversal(root.Right, list)
}

func convertBST(root *TreeNode) *TreeNode {
	list := []*TreeNode{}

	traversal(root, &list)

	greater := 0
	for i := len(list) - 1; i >= 0; i-- {
		node := list[i]
		val := node.Val
		node.Val += greater
		greater += val
	}

	return root
}
