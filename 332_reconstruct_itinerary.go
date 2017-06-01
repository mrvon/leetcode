// *Graph* Depth-first search solution
package main

import (
	"fmt"
	"sort"
)

func find(
	adjacency map[string][]string,
	unique map[string]int,
	departure string,
	path *[]string,
	length int) bool {

	if len(*path) >= length {
		return true
	}

	for i := 0; i < len(adjacency[departure]); i++ {
		arrival := adjacency[departure][i]
		if unique[departure+arrival] <= 0 {
			continue
		}
		(*path) = append(*path, arrival)
		unique[departure+arrival]--
		is_ok := find(adjacency, unique, arrival, path, length)
		if is_ok {
			return true
		}
		unique[departure+arrival]++
		(*path) = (*path)[:len(*path)-1]
	}

	return false
}

func findItinerary(tickets [][]string) []string {
	adjacency := make(map[string][]string)
	unique := make(map[string]int)

	for i := 0; i < len(tickets); i++ {
		departure := tickets[i][0]
		arrival := tickets[i][1]
		adjacency[departure] = append(adjacency[departure], arrival)
		unique[departure+arrival]++
	}

	for _, adj := range adjacency {
		sort.Strings(adj)
	}

	path := []string{"JFK"}
	find(adjacency, unique, "JFK", &path, len(tickets)+1)

	return path
}

func main() {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	fmt.Println(findItinerary(tickets))
	// Return ["JFK", "MUC", "LHR", "SFO", "SJC"].

	tickets = [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	fmt.Println(findItinerary(tickets))
	// Return ["JFK","ATL","JFK","SFO","ATL","SFO"].
	// Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"]. But it is larger in lexical order.

	tickets = [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	fmt.Println(findItinerary(tickets))
	// ["JFK","NRT","JFK","KUL"]

	tickets = [][]string{
		{"EZE", "AXA"},
		{"TIA", "ANU"},
		{"ANU", "JFK"},
		{"JFK", "ANU"},
		{"ANU", "EZE"},
		{"TIA", "ANU"},
		{"AXA", "TIA"},
		{"TIA", "JFK"},
		{"ANU", "TIA"},
		{"JFK", "TIA"},
	}
	fmt.Println(findItinerary(tickets))
}
