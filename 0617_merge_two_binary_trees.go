package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func value(t *TreeNode) int {
	if t == nil {
		return 0
	} else {
		return t.Val
	}
}

func left(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	} else {
		return t.Left
	}
}

func right(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	} else {
		return t.Right
	}
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	return &TreeNode{
		Val:   value(t1) + value(t2),
		Left:  mergeTrees(left(t1), left(t2)),
		Right: mergeTrees(right(t1), right(t2)),
	}
}
