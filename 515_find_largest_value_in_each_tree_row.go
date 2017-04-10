package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var max []int
	max = largest(root, 0, max)
	return max
}

func largest(root *TreeNode, depth int, max []int) []int {
	if root == nil {
		return max
	}

	if len(max) <= depth {
		max = append(max, math.MinInt32)
	}

	if root.Val > max[depth] {
		max[depth] = root.Val
	}

	max = largest(root.Left, depth+1, max)
	max = largest(root.Right, depth+1, max)

	return max
}
