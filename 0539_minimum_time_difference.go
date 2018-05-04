package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	timeNums := []int{}

	for i := 0; i < len(timePoints); i++ {
		t := strings.Split(timePoints[i], ":")
		h, _ := strconv.Atoi(t[0])
		m, _ := strconv.Atoi(t[1])
		timeNums = append(timeNums, h*60+m)
	}

	sort.Ints(timeNums)

	max := 24 * 60
	min := math.MaxInt32

	for i := 1; i < len(timeNums); i++ {
		d := timeNums[i] - timeNums[i-1]
		if d < min {
			min = d
		}
	}

	d := max - timeNums[len(timeNums)-1] + timeNums[0]
	if d < min {
		min = d
	}

	return min
}

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(findMinDifference([]string{"23:59", "00:01"}))
	fmt.Println(findMinDifference([]string{"00:05", "00:01"}))
}
