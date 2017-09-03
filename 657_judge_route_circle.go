package main

import "fmt"

func judgeCircle(moves string) bool {
	x := 0
	y := 0

	for i := 0; i < len(moves); i++ {
		m := moves[i]
		switch m {
		case 'U':
			x++
		case 'D':
			x--
		case 'L':
			y--
		case 'R':
			y++
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(judgeCircle(""))
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
	fmt.Println(judgeCircle("DURDLDRRLL"))
}
