package main

import (
	"fmt"
	"sort"
	"strconv"
)

type StrList []string

func (s StrList) Len() int {
	return len(s)
}

func (s StrList) Less(i, j int) bool {
	x := s[i]
	y := s[j]

	return x+y >= y+x
}

func (s StrList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	strs := []string{}

	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}

	sort.Sort(StrList(strs))

	result := []byte{}
	for i := 0; i < len(strs); i++ {
		result = append(result, strs[i]...)
	}

	// remove leading zero
	for len(result) > 0 && result[0] == '0' {
		result = result[1:]
	}

	if len(result) == 0 {
		return "0"
	} else {
		return string(result)
	}
}

func main() {
	fmt.Println(largestNumber([]int{3, 301, 30, 303, 34, 5, 9}))
	fmt.Println(largestNumber([]int{30, 303}))
	fmt.Println(largestNumber([]int{30, 301}))
	fmt.Println(largestNumber([]int{34, 341}))
	fmt.Println(largestNumber([]int{0, 0}))
}
