package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	cur := 0
	next_root := func() int {
		r := preorder[cur]
		cur++
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

	left := build(inorder[0:mid], next_root)
	right := build(inorder[mid+1:len(inorder)], next_root)

	return &TreeNode{
		Val:   root,
		Left:  left,
		Right: right,
	}
}
