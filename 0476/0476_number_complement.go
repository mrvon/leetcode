package main

import "fmt"

func findComplement(num int) int {
	n := int32(num)

	if n == 0 {
		return 1
	}

	for i := 31; i >= 0; i-- {
		if n&(1<<uint(i)) == 0 {
			n |= (1 << uint(i))
		} else {
			break
		}
	}

	return int(^n)
}

func main() {
	fmt.Println(findComplement(0))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(5))
}
