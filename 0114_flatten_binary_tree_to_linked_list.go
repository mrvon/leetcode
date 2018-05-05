package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	f(root)
}

func f(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil {
		root.Right = f(root.Right)
		return root
	} else {
		r := root.Right
		l := f(root.Left)
		// l must be not nil

		root.Left = nil
		root.Right = l

		// keep going until the last node
		for l.Right != nil {
			l = l.Right
		}

		l.Right = f(r)
		return root
	}
}
