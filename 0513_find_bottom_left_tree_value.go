package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	var max_depth int = -1
	var value int
	find(root, 0, &max_depth, &value)
	return value
}

func find(root *TreeNode, depth int, max_depth *int, value *int) {
	if root == nil {
		return
	}

	find(root.Left, depth+1, max_depth, value)
	if depth > *max_depth {
		*max_depth = depth
		*value = root.Val
	}
	find(root.Right, depth+1, max_depth, value)
}
