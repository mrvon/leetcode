// DFS
// 500ms
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

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func topDown(root *TreeNode, c int, m *int) {
	if root != nil {
		c += root.Val
		(*m) = max(*m, c)
		topDown(root.Left, c, m)
		topDown(root.Right, c, m)
	}
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	ml := 0
	mr := 0
	topDown(root.Left, 0, &ml)
	topDown(root.Right, 0, &mr)
	left := maxPathSum(root.Left)
	right := maxPathSum(root.Right)
	return max(root.Val+ml+mr, max(left, right))
}
