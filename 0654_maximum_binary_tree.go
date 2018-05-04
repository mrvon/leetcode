/*
O(n)

Stack
*/
package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// O(n^2)
func constructMaximumBinaryTreeNaive(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := nums[0]
	max_index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			max_index = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:max_index]),
		Right: constructMaximumBinaryTree(nums[max_index+1:]),
	}
}

// O(n)
type Stack struct {
	S []*TreeNode
}

func (S *Stack) push(n *TreeNode) {
	S.S = append(S.S, n)
}

func (S *Stack) pop() *TreeNode {
	n := S.S[len(S.S)-1]
	S.S = S.S[:len(S.S)-1]
	return n
}

func (S *Stack) peek() *TreeNode {
	return S.S[len(S.S)-1]
}

func (S *Stack) bottom() *TreeNode {
	return S.S[0]
}

func (S *Stack) isempty() bool {
	return len(S.S) == 0
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var s Stack

	for i := 0; i < len(nums); i++ {
		n := &TreeNode{
			Val: nums[i],
		}

		if s.isempty() {
			s.push(n)
		} else {
			if n.Val < s.peek().Val {
				s.peek().Right = n
				s.push(n)
			} else {
				p := s.pop()
				for !s.isempty() && n.Val >= s.peek().Val {
					p = s.pop()
				}
				n.Left = p
				if !s.isempty() {
					s.peek().Right = n
				}
				s.push(n)
			}
		}
	}

	if s.isempty() {
		return nil
	} else {
		return s.bottom()
	}
}
