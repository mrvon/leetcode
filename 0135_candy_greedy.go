/*
Time complexity is exactly O(N) not only the average time cost, and O(N) space
needed.

I also go thru the ratings array two times, from left to right and from right to
left.

Here are several steps of my algorithm:

Assume candies[i] means count of candies give to child i.

	1. From left to right, to make all candies satisfy the condition if
	ratings[i] > ratings[i - 1] then candies[i] > candies[i - 1], just ignore
	the right neighbor as this moment. We can assume leftCandies[i] is a
	solution which is satisfied.

	2. From right to left, almost like step 1, get a solution rightCandies[i]
	which just satisfy the condition if ratings[i] > ratings[i + 1] then
	candies[i] > candies[i + 1]

	3. For now, we have leftCandies[i] and rightCandies[i], so how can we
	satisfy the real condition in the problem? Just make candies[i] equals the
	maximum between leftCanides[i] and rightCandies[i]

Here are something to notice:

	+ Set all leftCandies[i] rightCandies[i] to 1 at the beginning of going
	thru, since each child need at least one candy.

	+ When you implement the algorithm, you do not need the real array
	leftCandies rightCandies, it just help to explain my thought. So there are
	only a candies[i] needed. However, it is O(N) space complexity.

Here is a sample:

	+ ratings: [1,5,3,1]
	+ in step 1, go from left to right, you can get leftCandies [1,2,1,1]
	+ in step 2, go from right to left, you can get rightCandies [1,3,2,1]
	+ in step 3, you can get the real candies [1,3,2,1], just sum it then get
	the answer.
*/
package main

import "fmt"

func candy(ratings []int) int {
	candy := []int{}
	for i := 0; i < len(ratings); i++ {
		candy = append(candy, 1)
	}
	for i := 1; i < len(candy); i++ {
		if ratings[i] > ratings[i-1] && candy[i-1]+1 > candy[i] {
			candy[i] = candy[i-1] + 1
		}
	}
	for i := len(candy) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candy[i+1]+1 > candy[i] {
			candy[i] = candy[i+1] + 1
		}
	}
	total := 0
	for i := 0; i < len(candy); i++ {
		total += candy[i]
	}
	return total
}

func main() {
	fmt.Println(candy([]int{0}))
	fmt.Println(candy([]int{12, 14, 12}))
	fmt.Println(candy([]int{12, 14, 12, 13}))
	fmt.Println(candy([]int{12, 12, 12, 13}))
	fmt.Println(candy([]int{12, 12, 12, 12}))
}
