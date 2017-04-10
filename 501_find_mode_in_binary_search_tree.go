package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goNext(list []int, cur int) int {
	cur++
	for cur < len(list)-1 && list[cur] == list[cur+1] {
		cur++
	}
	return cur
}

func findMode(root *TreeNode) []int {
	var list []int
	list = find(root, list)

	var result []int
	max := 0
	cur := -1
	for {
		old := cur
		cur = goNext(list, cur)
		if cur >= len(list) {
			break
		}
		length := cur - old
		if length > max {
			max = length
		}
	}

	cur = -1
	for {
		old := cur
		cur = goNext(list, cur)
		if cur >= len(list) {
			break
		}
		length := cur - old
		if length == max {
			result = append(result, list[cur])
		}
	}

	return result
}

// Inorder traversal
func find(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = find(root.Left, list)
	list = append(list, root.Val)
	list = find(root.Right, list)

	return list
}
