// TODO
package main

import (
	"fmt"
	"strconv"
)

func test(i int) {
	fmt.Print(strconv.Itoa(i))
}

func main() {
	fmt.Print("0")
	for i := 100; i < 110; i++ {
		test(i)
	}
}
