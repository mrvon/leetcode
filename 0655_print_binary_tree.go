package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func printTree(root *TreeNode) [][]string {
	m := height(root)
	n := int(math.Pow(2, float64(m))) - 1

	matrix := [][]string{}
	for i := 0; i < m; i++ {
		matrix = append(matrix, make([]string, n))
	}

	row := []*TreeNode{}
	row = append(row, root)
	for i := 0; i < m; i++ {
		indent := int(math.Pow(2, float64(m-i-1))) - 1
		step := int(math.Pow(2, float64(m-i)))
		for l := len(row); l > 0; l-- {
			node := row[0]
			row = row[1:]
			if node == nil {
				row = append(row, nil)
				row = append(row, nil)
			} else {
				matrix[i][indent] = strconv.Itoa(node.Val)
				row = append(row, node.Left)
				row = append(row, node.Right)
			}
			indent += step
		}
	}

	return matrix
}
