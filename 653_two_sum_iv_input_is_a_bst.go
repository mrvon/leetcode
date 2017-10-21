package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func find(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}

	m[root.Val]++

	other := k - root.Val

	if root.Val == other {
		if m[other] >= 2 {
			return true
		}
	} else {
		if m[other] >= 1 {
			return true
		}
	}

	return find(root.Left, k, m) || find(root.Right, k, m)
}

func findTarget(root *TreeNode, k int) bool {
	return find(root, k, make(map[int]int))
}
