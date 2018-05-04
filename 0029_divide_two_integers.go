/*
In this problem, we are asked to divide two integers. However, we are not
allowed to use division, multiplication and mod operations. So, what else can we
use? Yeah, bit manipulations.

Let's do an example and see how bit manipulations work.

Suppose we want to divide 15 by 3, so 15 is dividend and 3 is divisor. Well,
division simply requires us to find how many times we can subtract the divisor
from the the dividend without making the dividend negative.

Let's get started. We subtract 3 from 15 and we get 12, which is positive. Let's
try to subtract more. Well, we shift 3 to the left by 1 bit and we get 6.
Subtracting 6 from 15 still gives a positive result. Well, we shift again and
get 12. We subtract 12 from 15 and it is still positive. We shift again,
obtaining 24 and we know we can at most subtract 12. Well, since 12 is obtained
by shifting 3 to left twice, we know it is 4 times of 3. How do we obtain this
4? Well, we start from 1 and shift it to left twice at the same time. We add 4
to an answer (initialized to be 0). In fact, the above process is like 15 = 3 *
4 + 3. We now get part of the quotient (4), with a remainder 3.

Then we repeat the above process again. We subtract divisor = 3 from the
remaining dividend = 3 and obtain 0. We know we are done. No shift happens, so
we simply add 1 << 0 to the answer.

Now we have the full algorithm to perform division.

According to the problem statement, we need to handle some exceptions, such as
overflow.

Well, two cases may cause overflow:

1. divisor = 0;
2. dividend = INT_MIN and divisor = -1 (because abs(INT_MIN) = INT_MAX + 1).

Of course, we also need to take the sign into considerations, which is
relatively easy.

Putting all these together, we have the following code.
*/
package main

import (
	"fmt"
	"math"
)

func abs(n int64) int64 {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func divide(dividend int, divisor int) int {
	// check overflow
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}

	dvd := abs(int64(dividend))
	dvs := abs(int64(divisor))
	res := 0

	for dvd >= dvs {
		tmp := dvs
		multiple := int64(1)
		for dvd >= (tmp << 1) {
			tmp <<= 1
			multiple <<= 1
		}
		dvd -= tmp
		res += int(multiple)
	}

	// check sign
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		return res
	} else {
		return -res
	}
}

func main() {
	fmt.Println(divide(15, 3))
	fmt.Println(divide(15, -3))
	fmt.Println(divide(-15, 3))
	fmt.Println(divide(-15, -3))
	fmt.Println(divide(-2147483648, -1))
}
