package main

import "fmt"

// Straightforward
func longestConsecutive1(nums []int) int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	best := 0
	for len(m) > 0 {
		count := 0
		// Get one
		n := 0
		for n = range m {
			break
		}
		count += 1
		delete(m, n)
		for i := n + 1; m[i] > 0; i++ {
			count += 1
			delete(m, i)
		}
		for i := n - 1; m[i] > 0; i-- {
			count += 1
			delete(m, i)
		}
		if count > best {
			best = count
		}
	}
	return best
}

// More elegant
func longestConsecutive(nums []int) int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	best := 0
	for n := range m {
		if m[n-1] > 0 {
			continue
		}
		count := 0
		for i := n; m[i] > 0; i++ {
			count++
		}
		if count > best {
			best = count
		}
	}
	return best
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 0, 1}))
	fmt.Println(longestConsecutive([]int{0, -1, 0, 1}))
}
