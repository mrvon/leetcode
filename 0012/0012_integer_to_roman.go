package main

import "fmt"

type Item struct {
	I int
	R string
}

var Roman []Item = []Item{
	Item{1000, "M"},
	Item{900, "CM"},
	Item{500, "D"},
	Item{400, "CD"},
	Item{100, "C"},
	Item{90, "XC"},
	Item{50, "L"},
	Item{40, "XL"},
	Item{10, "X"},
	Item{9, "IX"},
	Item{5, "V"},
	Item{4, "IV"},
	Item{1, "I"},
}

func MaxRoman(num int) Item {
	for i := 0; i < len(Roman); i++ {
		if num >= Roman[i].I {
			return Roman[i]
		}
	}
	panic("error")
}

func intToRoman(num int) string {
	item := MaxRoman(num)
	if item.I == num {
		return item.R
	} else {
		return item.R + intToRoman(num-item.I)
	}
}

func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(intToRoman(i))
	}
}
