/*
assume we have n bulb. so we will toogle n round.
bulb 1, it will toogle 1 times(in 1th toogle).
bulb 2, it will toogle 2 times(in 1th, 2th toogle).
bulb 3, it will toogle 2 times(in 1th, 3th toogle).
bulb 4, it will toogle 3 times(in 1th, 2th, 4th toogle).
bulb 5, it will toogle 2 times(in 1th, 5th toogle).
...
bulb n

As you can see, if we toogle odd number times, the bulb will on in the end,
otherwise it'll off. But how do we know how many times it'll be toogle? In blub
n, the magic is how many disivor factor in number n. if the number of factor is
odd, it'll be on in the end. And exactly when n = sqrt(n)*sqrt(n), the number of
factor is odd.
*/
package main

import (
	"fmt"
)

/*
The original method describe above.
But is too slow. Time limit exceeded.

func bulbSwitch(n int) int {
	c := 0
	for i := 1; i <= n; i++ {
		s := int(math.Sqrt(float64(i)))
		if s*s == i {
			i++
			c++
		}
	}
	return c
}
*/

// Better one, accept
func bulbSwitch(n int) int {
	i := 1
	c := 0
	for i*i <= n {
		i++
		c++
	}
	return c
}

func main() {
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(2))
	fmt.Println(bulbSwitch(3))
	fmt.Println(bulbSwitch(4))
	fmt.Println(bulbSwitch(5))
	fmt.Println(bulbSwitch(6))
	fmt.Println(bulbSwitch(16))
	fmt.Println(bulbSwitch(32))
}
