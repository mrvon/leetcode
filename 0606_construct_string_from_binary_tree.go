package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	r := strconv.Itoa(t.Val)

	if t.Left == nil && t.Right == nil {
		return r
	} else if t.Left == nil {
		return r + "()" + "(" + tree2str(t.Right) + ")"
	} else if t.Right == nil {
		return r + "(" + tree2str(t.Left) + ")"
	} else {
		return r + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
	}
}
