// insert sort
// stable
// very useful in small array
package main

import "fmt"

func insert_sort(a []string, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; j > left && a[j] < a[j-1]; j-- {
			// exchange
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	list := []string{
		"goa", "go", "goo", "flash", "to", "the", "zoo", "the", "game", "of",
		"world", "or", "and", "not", "bee", "box", "color",
	}
	insert_sort(list, 0, len(list)-1)
	fmt.Println(list)
}
