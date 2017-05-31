package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	m1 := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		m1[list1[i]] = i
	}
	min_sum := math.MaxInt32
	min_res := []string{}
	for i := 0; i < len(list2); i++ {
		if index, has := m1[list2[i]]; has {
			if index+i <= min_sum {
				if index+i < min_sum {
					min_res = []string{list2[i]}
				} else {
					min_res = append(min_res, list2[i])
				}
				min_sum = index + i
			}
		}
	}
	return min_res
}

func main() {
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	))

	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
