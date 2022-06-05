package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
	sum   float64
	count int
}

func average(result *[]Pair, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level >= len(*result) {
		(*result) = append(*result, Pair{})
	}

	(*result)[level].sum += float64(root.Val)
	(*result)[level].count += 1

	average(result, root.Left, level+1)
	average(result, root.Right, level+1)
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []Pair
	average(&result, root, 0)

	var ave []float64
	for i := 0; i < len(result); i++ {
		ave = append(ave, result[i].sum/float64(result[i].count))
	}
	return ave
}
