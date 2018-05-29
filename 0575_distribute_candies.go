package main

func distributeCandies(candies []int) int {
	kind := make(map[int]bool)
	n := 0

	for i := 0; i < len(candies); i++ {
		if !kind[candies[i]] {
			kind[candies[i]] = true
			n++
		}
	}

	if n > len(candies)/2 {
		n = len(candies) / 2
	}

	return n
}
