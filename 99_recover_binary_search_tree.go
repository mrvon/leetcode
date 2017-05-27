package main

import "fmt"

/*
O(n) Space, straightforward solution

Two elements of a binary search tree (BST) are swapped by mistake, there are two
cases actually:

Original:
	1, 2, 3, 5, 6, 9

First, We swapped the two elements are continuous in the traversal order:

	1, 3, 2, 5, 6, 9

We can see only pair(3, 2) is out of order, reverse this pair is all done.

Second, We swapped the two elements aren't continuous in the traversal order:

	1, 6, 3, 5, 2, 9

We can see pair(6, 3), pair(5, 2) is out of order, swapped first element of
first pair and second element of second pair is all done.
*/
type Pair struct {
	first  *TreeNode
	second *TreeNode
}

func traverse(root *TreeNode, array *[]*TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left, array)
	(*array) = append(*array, root)
	traverse(root.Right, array)
}

func recoverTreeN(root *TreeNode) {
	array := []*TreeNode{}
	traverse(root, &array)

	var first *Pair
	var second *Pair

	for i := 0; i < len(array)-1; i++ {
		if array[i].Val > array[i+1].Val {
			if first == nil {
				first = &Pair{array[i], array[i+1]}
			} else {
				second = &Pair{array[i], array[i+1]}
			}
		}
	}

	if second == nil {
		first.first.Val, first.second.Val = first.second.Val, first.first.Val
	} else {
		first.first.Val, second.second.Val = second.second.Val, first.first.Val
	}
}

/*
O(1) space solution

Morris Traversal
*/

type State struct {
	fnode  *TreeNode
	snode  *TreeNode
	first  *Pair
	second *Pair
}

// O(1) space inorder tree traversal
func morrisTraversal(root *TreeNode, s *State, f func(s *State, node *TreeNode)) {
	for root != nil { // While current is not nil
		if root.Left == nil { // If current does not have left child
			f(s, root)        // print the current's data
			root = root.Right // go to the right child
		} else {
			// find the right child of the rightmost node in current's left
			// subtree, also known as "predecessor" of current node
			var predecessor *TreeNode = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// if predecessor does not have right child, set current node as
				// right child of predecessor, and go to the left child
				predecessor.Right = root
				root = root.Left
			} else {
				// if predecessor does have right child, it must be current node
				// here. print the current's data, set predecessor's right child
				// as nil (recover the tree structure), and go to the right
				// child.
				f(s, root)
				predecessor.Right = nil
				root = root.Right
			}
		}
	}
}

func recoverTree(root *TreeNode) {
	s := &State{}
	morrisTraversal(root, s, func(s *State, node *TreeNode) {
		if s.fnode == nil {
			s.fnode = node
		} else if s.snode == nil {
			s.snode = node
		} else {
			s.fnode, s.snode = s.snode, node
		}
		if s.fnode != nil && s.snode != nil && s.fnode.Val > s.snode.Val {
			if s.first == nil {
				s.first = &Pair{s.fnode, s.snode}
			} else {
				s.second = &Pair{s.fnode, s.snode}
			}
		}
	})
	if s.second == nil {
		s.first.first.Val, s.first.second.Val = s.first.second.Val, s.first.first.Val
	} else {
		s.first.first.Val, s.second.second.Val = s.second.second.Val, s.first.first.Val
	}
}
