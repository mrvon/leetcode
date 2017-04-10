package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func numbers(root *TreeNode, buf *[]byte, sum *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		// Leaf
		(*buf) = append((*buf), byte(root.Val+'0'))
		i, _ := strconv.Atoi(string(*buf))
		(*buf) = (*buf)[:len(*buf)-1]
		(*sum) += i
	} else {
		(*buf) = append((*buf), byte(root.Val+'0'))
		numbers(root.Left, buf, sum)
		numbers(root.Right, buf, sum)
		(*buf) = (*buf)[:len(*buf)-1]
	}
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	buf := []byte{}
	numbers(root, &buf, &sum)
	return sum
}
