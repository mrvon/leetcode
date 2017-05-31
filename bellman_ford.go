/*
The Bellman-Ford algorithm solves the single-source shortest-paths problem in
the general case in which edge weights may be negative. Give a weighted,
directed graph G = (V, E) with source s and weight function w: E -> R, the
Bellman-Ford algorithm returns a boolean value indicating whether or not there
is a negative-weight cycle that is reachable from the source. If there is such a
cycle, the algorithm indicates that no solution exists. If there is no such
cycle, the algorithm produces the shortest paths and their weights.

The algorithm relaxes edges, progressively decreasing an estimate v.d on the
weight of a shortest path from the source s to each vertex v in V until it
archieves the actual shortest-path weight. The algorithm return TRUE if and only
if the graph contains no negative-weight cycles that are reachable from the
source.
*/

package main

func main() {
}
