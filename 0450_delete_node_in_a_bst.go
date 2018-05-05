package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteThis(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root.Left
	} else {
		right := root.Right

		if right.Left == nil {
			right.Left = root.Left
			return right
		} else {
			// Find leftmost node in the right subtree
			for right.Left.Left != nil {
				right = right.Left
			}
			new_root := right.Left
			right.Left = new_root.Right
			new_root.Left = root.Left
			new_root.Right = root.Right

			return new_root
		}
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		root = deleteThis(root)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
