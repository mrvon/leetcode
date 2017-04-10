package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	var sum []int
	find(root, &sum)

	max := 0
	counter := make(map[int]int)
	for i := 0; i < len(sum); i++ {
		counter[sum[i]]++
		if counter[sum[i]] > max {
			max = counter[sum[i]]
		}
	}
	var result []int
	for s, c := range counter {
		if c == max {
			result = append(result, s)
		}
	}
	return result
}

func find(root *TreeNode, sum *[]int) int {
	if root == nil {
		return 0
	}

	left := find(root.Left, sum)
	right := find(root.Right, sum)
	s := left + right + root.Val
	*sum = append(*sum, s)

	return s
}
