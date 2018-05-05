/*
First let's review some statement for tree in graph theory:

(1) A tree is an undirected graph in which any two vertices are connected by
exactly one path.

(2) Any connected graph who has n nodes with n-1 edges is a tree.

(3) The degree of a vertex of a graph is the number of edges incident to the
vertex.

(4) A leaf is a vertex of degree 1. An internal vertex is a vertex of degree at
least 2.

(5) A path graph is a tree with two or more vertices that is not branched at
all.

(6) A tree is called a rooted tree if one vertex has been designated the root.

(7) The height of a rooted tree is the number of edges on the longest downward
path between root and a leaf.  OK. Let's stop here and look at our problem.

Our problem want us to find the minimum height trees and return their root
labels. First we can think about a simple case -- a path graph.

For a path graph of n nodes, find the minimum height trees is trivial. Just
designate the middle point(s) as roots.

Despite its triviality, let design a algorithm to find them.

Suppose we don't know n, nor do we have random access of the nodes. We have to
traversal. It is very easy to get the idea of two pointers. One from each end
and move at the same speed. When they meet or they are one step away, (depends
on the parity of n), we have the roots we want.

This gives us a lot of useful ideas to crack our real problem.

For a tree we can do some thing similar. We start from every end, by end we mean
vertex of degree 1 (aka leaves). We let the pointers move the same speed. When
two pointers meet, we keep only one of them, until the last two pointers meet or
one step away we then find the roots.

It is easy to see that the last two pointers are from the two ends of the
longest path in the graph.

The actual implementation is similar to the BFS topological sort. Remove the
leaves, update the degrees of inner vertexes. Then remove the new leaves. Doing
so level by level until there are 2 or 1 nodes left. What's left is our answer!

The time complexity and space complexity are both O(n).

Note that for a tree we always have V = n, E = n-1.
*/
package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjacency := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		adjacency[edges[i][0]] = append(adjacency[edges[i][0]], edges[i][1])
		adjacency[edges[i][1]] = append(adjacency[edges[i][1]], edges[i][0])
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(adjacency[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		n -= len(leaves)
		new_leaves := []int{}
		for i := 0; i < len(leaves); i++ {
			l := leaves[i]
			r := adjacency[l][0]
			for j := 0; j < len(adjacency[r]); j++ {
				if adjacency[r][j] == l {
					copy(adjacency[r][j:], adjacency[r][j+1:])
					adjacency[r] = adjacency[r][:len(adjacency[r])-1]
					break
				}
			}
			adjacency[l] = adjacency[l][1:]
			if len(adjacency[r]) == 1 {
				new_leaves = append(new_leaves, r)
			}
		}
		leaves = new_leaves
	}

	return leaves
}

func main() {
	fmt.Println(findMinHeightTrees(
		4,
		[][]int{{1, 0}, {1, 2}, {1, 3}},
	))

	fmt.Println(findMinHeightTrees(
		6,
		[][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}},
	))
}
