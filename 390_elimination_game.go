/*
After first elimination, all the remaining numbers are even numbers.  We can
divide them by 2, to get a continuous new sequence from 1 to n/2.  This is a sub
problem, we can recursively solve it, and the second elimination should be
reverse direction.

For example,

lastRemaining(8)
	1,2,3,4,5,6,7,8 // original sequence
	_,2,_,4,_,6,_,8 // first elimination
	_,1,_,2,_,3,_,4 // divide by 2, sub problem(from right to left)
	// should be (5 - lastRemaining(4))

*/
package main

import "fmt"

func lastRemaining(n int) int {
	if n == 1 {
		return n
	} else {
		return 2 * (n/2 + 1 - lastRemaining(n/2))
	}
}

func main() {
	fmt.Println(lastRemaining(1))
	fmt.Println(lastRemaining(2))
	fmt.Println(lastRemaining(3))
	fmt.Println(lastRemaining(4))
	fmt.Println(lastRemaining(8))
}
