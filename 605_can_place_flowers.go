package main

import "fmt"

func canPlace(flowerbed []int, n int, cur int) bool {
	if n <= 0 {
		return true
	}
	for i := cur; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			return canPlace(flowerbed, n-1, i+1)
		}
	}
	return false
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	return canPlace(flowerbed, n, 0)
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
}
