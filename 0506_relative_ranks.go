package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Item struct {
	score  int
	origin int // original position
}

type Rank []Item

func (r Rank) Len() int {
	return len(r)
}

func (r Rank) Less(i, j int) bool {
	return r[i].score >= r[j].score
}

func (r Rank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func findRelativeRanks(nums []int) []string {
	var rank Rank
	for i := 0; i < len(nums); i++ {
		rank = append(rank, Item{
			nums[i], i,
		})
	}
	sort.Sort(rank)
	result := make([]string, len(nums))
	for r := 0; r < len(rank); r++ {
		item := rank[r]
		s := ""
		if r == 0 {
			s = "Gold Medal"
		} else if r == 1 {
			s = "Silver Medal"
		} else if r == 2 {
			s = "Bronze Medal"
		} else {
			s = strconv.Itoa(r + 1)
		}
		result[item.origin] = s
	}
	return result
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{3, 4, 2, 5, 1}))
}
