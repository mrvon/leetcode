package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(N int) []*TreeNode {
	// assert(N%2 == 1 && N >= 1)
	nodeList := []*TreeNode{}
	if N == 1 {
		nodeList = append(nodeList, &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		})
	} else if N%2 == 1 {
		for leftN := 1; leftN <= N-2; leftN += 2 {
			rightN := N - leftN - 1
			leftL := allPossibleFBT(leftN)
			rightL := allPossibleFBT(rightN)
			for i := 0; i < len(leftL); i++ {
				for j := 0; j < len(rightL); j++ {
					nodeList = append(nodeList, &TreeNode{
						Val:   0,
						Left:  leftL[i],
						Right: rightL[j],
					})
				}
			}
		}
	}
	return nodeList
}
