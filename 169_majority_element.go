package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Straightforward solution
// time O(n), space O(n)
func majorityElementS(nums []int) int {
	dict := make(map[int]int)
	size := len(nums)

	for _, n := range nums {
		dict[n]++
	}

	bound := int(math.Floor(float64(size) / 2.0))
	for n, c := range dict {
		if c > bound {
			return n
		}
	}

	return 0
}

// Randomized
func majorityElementR(nums []int) int {
	size := len(nums)
	bound := int(math.Floor(float64(size) / 2.0))
	for {
		r := nums[rand.Int()%size]
		count := 0
		for i := 0; i < size; i++ {
			if nums[i] == r {
				count++
			}
			if count > bound {
				return r
			}
		}
	}
}

// Time O(n), Space O(1)
// Boyer–Moore majority vote algorithm
func majorityElementM(nums []int) int {
	major := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count++
		} else if nums[i] == major {
			count++
		} else {
			count--
		}
	}
	return major
}

// Time O(n), Space O(1)
// Bit manipulation
// Tricky here, you must notice different between int and int32
func majorityElementB(nums []int) int {
	size := len(nums)
	bound := int(math.Floor(float64(size) / 2.0))
	major := int32(0)
	mask := int32(1)
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < size; j++ {
			if (int32(nums[j]) & mask) != 0 {
				count++
			}
			if count > bound {
				major |= mask
				break
			}
		}
		mask = mask << 1
	}
	return int(major)
}

// Divide and conquer
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	l := majorityElement(nums[:mid])
	r := majorityElement(nums[mid:])
	if l == r {
		return l
	} else {
		lcount := 0
		rcount := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == l {
				lcount++
			} else if nums[i] == r {
				rcount++
			}
		}
		if lcount > rcount {
			return l
		} else {
			return r
		}
	}
}

func main() {
	fmt.Println(majorityElementS([]int{1, 2, 2, 2, 4}))
	fmt.Println(majorityElementS([]int{1, 2, 4, 4, 4}))
	fmt.Println(majorityElementS([]int{-2147483648}))
	fmt.Println(majorityElementS([]int{1}))

	fmt.Println(majorityElementR([]int{1, 2, 2, 2, 4}))
	fmt.Println(majorityElementR([]int{1, 2, 4, 4, 4}))
	fmt.Println(majorityElementR([]int{-2147483648}))
	fmt.Println(majorityElementR([]int{1}))

	fmt.Println(majorityElementM([]int{1, 2, 2, 2, 4}))
	fmt.Println(majorityElementM([]int{1, 2, 4, 4, 4}))
	fmt.Println(majorityElementM([]int{-2147483648}))
	fmt.Println(majorityElementM([]int{1}))

	fmt.Println(majorityElementB([]int{1, 2, 2, 2, 4}))
	fmt.Println(majorityElementB([]int{1, 2, 4, 4, 4}))
	fmt.Println(majorityElementB([]int{-2147483648}))
	fmt.Println(majorityElementB([]int{1}))

	fmt.Println(majorityElement([]int{1, 2, 2, 2, 4}))
	fmt.Println(majorityElement([]int{1, 2, 4, 4, 4}))
	fmt.Println(majorityElement([]int{-2147483648}))
	fmt.Println(majorityElement([]int{1}))
}
