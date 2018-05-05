// Brain Fuck, I hate this problem really.
package main

import (
	"fmt"
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	size := 1
	count := 9
	for {
		if n <= count*size {
			start := int(math.Pow10(size - 1))
			number := start + (n-1)/size
			index := (n - 1) % size
			digits := strconv.Itoa(number)
			digit := int(digits[index] - '0')
			return digit
		} else {
			n -= (count * size)
			count *= 10
			size += 1
		}
	}
}

func main() {
	fmt.Println(findNthDigit(1))
	fmt.Println(findNthDigit(9))
	fmt.Println(findNthDigit(10), findNthDigit(11))
	fmt.Println(findNthDigit(12), findNthDigit(13))
	fmt.Println(findNthDigit(14), findNthDigit(15))
	fmt.Println(findNthDigit(190), findNthDigit(191), findNthDigit(192))
	fmt.Println(findNthDigit(193), findNthDigit(194), findNthDigit(195))
}
