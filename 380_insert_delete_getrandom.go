package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	m map[int]int
	s []int
}

/*
 * Initialize your data structure here.
 */
func Constructor() RandomizedSet {
	set := RandomizedSet{}
	set.m = make(map[int]int)
	set.s = []int{}
	return set
}

/*
 * Inserts a value to the set. Returns true if the set did not already contain
 * the specified element.
 */
func (this *RandomizedSet) Insert(val int) bool {
	// Next index for val
	next := len(this.s)

	if this.m[val] == 0 {
		// If val not in set, store it's index+1
		this.m[val] = next + 1
		this.s = append(this.s, val)
		return true
	} else {
		// If val in set
		return false
	}
}

/*
 * Removes a value from the set. Returns true if the set contained the
 * specified element.
 */

func (this *RandomizedSet) Remove(val int) bool {
	if this.m[val] == 0 {
		// If val not is set
		return false
	} else {
		// If val in set
		last_val := this.s[len(this.s)-1]
		this.m[last_val] = this.m[val]
		this.s[this.m[val]-1] = last_val
		this.m[val] = 0
		// Remove last element
		this.s = this.s[:len(this.s)-1]
		return true
	}
}

/*
 * Get a random element from the set.
 */
func (this *RandomizedSet) GetRandom() int {
	size := len(this.s)
	if size == 0 {
		return 0
	}
	r := rand.Int() % size
	return int(this.s[r])
}

func main() {
	obj := Constructor()

	for i := 0; i < 100; i++ {
		obj.Insert(i)
		obj.Insert(-i)
	}

	for i := 0; i < 100; i++ {
		obj.Remove(i)
		obj.Remove(-i)
	}

	fmt.Println(obj.GetRandom())
}
