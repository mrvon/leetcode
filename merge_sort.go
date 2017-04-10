// merge sort
// stable
package main

import "fmt"

func merge(a []string, left int, mid int, right int) {
	var t []string
	for i := left; i <= right; i++ {
		t = append(t, a[i])
	}
	i := left
	j := mid + 1
	k := left
	for ; i <= mid || j <= right; k++ {
		if i > mid {
			a[k] = t[j-left]
			j++
		} else if j > right {
			a[k] = t[i-left]
			i++
		} else {
			if t[i-left] < t[j-left] {
				a[k] = t[i-left]
				i++
			} else {
				a[k] = t[j-left]
				j++
			}
		}
	}
}

func merge_sort(a []string, left int, right int) {
	if right-left <= 0 {
		return
	}

	mid := left + (right-left)/2
	merge_sort(a, left, mid)
	merge_sort(a, mid+1, right)
	merge(a, left, mid, right)
}

func main() {
	list := []string{
		"goa", "go", "goo", "flash", "to", "the", "zoo", "the", "game", "of",
		"world", "or", "and", "not", "bee", "box", "color",
	}
	merge_sort(list, 0, len(list)-1)
	fmt.Println(list)
}
