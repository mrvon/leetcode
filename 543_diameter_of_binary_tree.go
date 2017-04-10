package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	length := 0
	diameter(root, &length)
	return length
}

func diameter(root *TreeNode, length *int) int {
	if root == nil {
		return 0
	}

	l := diameter(root.Left, length)
	r := diameter(root.Right, length)

	*length = max(*length, l+r)

	return max(l, r) + 1
}
