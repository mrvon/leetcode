// Max Heap
package main

import "fmt"

func Less(array []string, i int, j int) bool {
	return array[i] < array[j]
}

func Up(array []string, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !Less(array, j, i) {
			break
		}
		// swap
		array[i], array[j] = array[j], array[i]
		j = i
	}
}

func Down(array []string, i int, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		j2 := j1 + 1
		if j2 < n && !Less(array, j1, j2) {
			j = j2 // = 2*i + 2 // right child
		}
		if !Less(array, j, i) {
			break
		}
		array[i], array[j] = array[j], array[i]
		i = j
	}
}

func HeapSort(array []string) {
	// build heap
	for i := 1; i < len(array); i++ {
		Up(array, i)
	}

	n := len(array) - 1
	for n > 0 {
		// array[0] is the root and largest value.
		// The swap moves it in front of the sorted elements.
		array[n], array[0] = array[0], array[n]
		// the heap size is reduced by one
		n--
		// the swap ruined the heap property, so restore it.
		Down(array, 0, n)
	}
}

func main() {
	array := []string{
		"goa", "go", "goo", "flash", "to", "the", "zoo", "the", "game", "of",
		"world", "or", "and", "not", "bee", "box", "color",
	}
	HeapSort(array)
	fmt.Println(array)
}
