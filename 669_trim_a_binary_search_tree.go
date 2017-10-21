package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	} else {
		root.Left = trimBST(root.Left, L, root.Val)
		root.Right = trimBST(root.Right, root.Val, R)
		return root
	}
}
