package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		if isSame(s, t) {
			return true
		}
	}
	if isSubtree(s.Left, t) {
		return true
	}
	if isSubtree(s.Right, t) {
		return true
	}
	return false
}
