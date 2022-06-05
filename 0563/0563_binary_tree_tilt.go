package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func find(root *TreeNode) (sum int, tilt int) {
	if root == nil {
		sum = 0
		tilt = 0
		return
	}

	left_sum, left_tilt := find(root.Left)
	right_sum, right_tilt := find(root.Right)

	sum = left_sum + right_sum + root.Val
	tilt = left_tilt + right_tilt + abs(left_sum, right_sum)

	return
}

func findTilt(root *TreeNode) int {
	_, tilt := find(root)
	return tilt
}
