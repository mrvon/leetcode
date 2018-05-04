package main

import "fmt"

// -1 is cann't rotate
// 0 is can rotate and value will not be change
// 1 is can rotate and value will be change
var rotateMap = []int{
	0,
	0,
	1,
	-1,
	-1,
	1,
	1,
	-1,
	0,
	1,
}

func digits(n int) []int {
	l := []int{}
	for n > 0 {
		l = append(l, n%10)
		n = n / 10
	}
	return l
}

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		sum := 0
		for _, d := range digits(i) {
			t := rotateMap[d]
			if t == -1 {
				sum = -1 // set a flag
				break
			}
			sum += t
		}
		if sum > 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(rotatedDigits(2))
	fmt.Println(rotatedDigits(10))
	fmt.Println(rotatedDigits(20))
	fmt.Println(rotatedDigits(30))
	fmt.Println(rotatedDigits(50))
	fmt.Println(rotatedDigits(100))
	fmt.Println(rotatedDigits(857))
}
