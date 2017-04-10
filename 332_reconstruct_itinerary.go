// Topological sort using Kahn's algorithm
package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
}

func main() {
	// tickets := [][]string{
	//     {"MUC", "LHR"},
	//     {"JFK", "MUC"},
	//     {"SFO", "SJC"},
	//     {"LHR", "SFO"},
	// }
	// fmt.Println(findItinerary(tickets))
	// Return ["JFK", "MUC", "LHR", "SFO", "SJC"].

	// tickets = [][]string{
	//     {"JFK", "SFO"},
	//     {"JFK", "ATL"},
	//     {"SFO", "ATL"},
	//     {"ATL", "JFK"},
	//     {"ATL", "SFO"},
	// }
	// fmt.Println(findItinerary(tickets))
	// Return ["JFK","ATL","JFK","SFO","ATL","SFO"].
	// Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"]. But it is larger in lexical order.

	tickets := [][]string{
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
