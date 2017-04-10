package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt32, math.MaxInt32)
}

func isValid(root *TreeNode, min_b int, max_b int) bool {
	if root == nil {
		return true
	}

	if root.Val < min_b || root.Val > max_b {
		return false
	}

	return isValid(root.Left, min_b, min(max_b, root.Val-1)) &&
		isValid(root.Right, max(min_b, root.Val+1), max_b)
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// Another solution using in-order traversal

func isValidBST2(root *TreeNode) bool {
	// in-order traversal
	order := traversal(root, []int{})
	// check it
	for i := 0; i < len(order)-1; i++ {
		if order[i] >= order[i+1] {
			return false
		}
	}
	return true
}

func traversal(root *TreeNode, order []int) []int {
	if root == nil {
		return order
	}

	order = traversal(root.Left, order)
	order = append(order, root.Val)
	order = traversal(root.Right, order)

	return order
}
