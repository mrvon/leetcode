package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	cur := len(postorder) - 1
	next_root := func() int {
		r := postorder[cur]
		cur--
		return r
	}

	return build(inorder, next_root)
}

func build(inorder []int, next_root func() int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := next_root()
	mid := 0
	for {
		if mid >= len(inorder) {
			panic("error")
		}
		if inorder[mid] == root {
			break
		}
		mid++
	}

	right := build(inorder[mid+1:len(inorder)], next_root)
	left := build(inorder[0:mid], next_root)

	return &TreeNode{
		Val:   root,
		Left:  left,
		Right: right,
	}
}
