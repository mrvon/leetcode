// Memoized + DFS
// 70ms
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

func topdown(root *TreeNode, memoized map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if m, has := memoized[root]; has {
		return m
	} else {
		m := max(
			root.Val+max(
				topdown(root.Left, memoized),
				topdown(root.Right, memoized),
			),
			0,
		)
		memoized[root] = m
		return m
	}
}

func mps(root *TreeNode, memoized map[*TreeNode]int) int {
	if root == nil {
		return math.MinInt32
	}
	left := mps(root.Left, memoized)
	right := mps(root.Right, memoized)
	return max(
		root.Val+topdown(root.Left, memoized)+topdown(root.Right, memoized),
		max(left, right),
	)
}

func maxPathSum(root *TreeNode) int {
	return mps(root, make(map[*TreeNode]int))
}
