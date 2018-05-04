package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	first := timeSeries[0]
	total := duration

	for i := 1; i < len(timeSeries); i++ {
		t := timeSeries[i]
		if t-first >= duration {
			// distinct
			total += duration
		} else {
			// overlap
			total += (t - first)
		}
		first = t
	}
	return total
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}
