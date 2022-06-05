package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	list := minimum(root, []int{})
	min := math.MaxInt32
	for i := 1; i < len(list); i++ {
		d := list[i] - list[i-1]
		if d < min {
			min = d
		}
	}
	return min
}

func minimum(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = minimum(root.Left, list)
	list = append(list, root.Val)
	list = minimum(root.Right, list)
	return list
}
