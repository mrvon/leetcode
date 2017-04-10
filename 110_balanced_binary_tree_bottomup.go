// Bottom up, O(N^2)
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func height(root *TreeNode) int { // -1 means already not balance.
	if root == nil {
		return 0
	}

	left_height := height(root.Left)
	if left_height == -1 {
		return -1
	}

	right_height := height(root.Right)
	if right_height == -1 {
		return -1
	}

	if abs(left_height-right_height) > 1 {
		return -1
	}

	return max(left_height, right_height) + 1
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}
