/*
Memoized solution
*/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type M struct {
	tree       *TreeNode
	can_choose bool
}

// Naive recursive solution
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func __rob(memoized map[M]int, root *TreeNode, can_choose bool) int {
	if root == nil {
		return 0
	}

	if count, has := memoized[M{root, can_choose}]; has {
		return count
	}

	result := 0
	if !can_choose {
		result = __rob(memoized, root.Left, true) + __rob(memoized, root.Right, true)
	} else {
		result = max(
			root.Val+__rob(memoized, root.Left, false)+__rob(memoized, root.Right, false),
			__rob(memoized, root.Left, true)+__rob(memoized, root.Right, true))
	}
	memoized[M{root, can_choose}] = result
	return result
}

func rob(root *TreeNode) int {
	return __rob(make(map[M]int), root, true)
}
