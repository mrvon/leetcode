package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSide(root *TreeNode, depth int, view []int) []int {
	if root == nil {
		return view
	}

	view = rightSide(root.Left, depth+1, view)

	for i := len(view); i <= depth; i++ {
		view = append(view, 0)
	}
	view[depth] = root.Val

	view = rightSide(root.Right, depth+1, view)

	return view
}

func rightSideView(root *TreeNode) []int {
	return rightSide(root, 0, []int{})
}
