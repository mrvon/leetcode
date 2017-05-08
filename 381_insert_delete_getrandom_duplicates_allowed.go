package main

import (
	"fmt"
	"math/rand"
)

type RandomizedCollection struct {
	m map[int]map[int]bool
	s []int
}

/*
 * Initialize your data structure here.
 */
func Constructor() RandomizedCollection {
	set := RandomizedCollection{}
	set.m = make(map[int]map[int]bool)
	set.s = []int{}
	return set
}

/*
 * Inserts a value to the collection. Returns true if the collection did not
 * already contain the specified element.
 */
func (this *RandomizedCollection) Insert(val int) bool {
	// Next index for val
	next := len(this.s)
	is_first := false

	if _, has := this.m[val]; !has {
		this.m[val] = make(map[int]bool)
		is_first = true
	}
	this.m[val][next] = true
	this.s = append(this.s, val)

	return is_first
}

/*
 * Removes a value from the collection. Returns true if the collection contained
 * the specified element.
 */
func (this *RandomizedCollection) Remove(val int) bool {
	if indexs, has := this.m[val]; !has || len(indexs) == 0 {
		return false
	} else {
		index := 0
		for index = range this.m[val] {
			break
		}
		delete(this.m[val], index)
		last_index := len(this.s) - 1
		last_val := this.s[last_index]
		this.m[last_val][index] = true
		this.s[index] = last_val
		delete(this.m[last_val], last_index)
		this.s = this.s[:last_index]
		return true
	}
}

/*
 * Get a random element from the collection.
 */
func (this *RandomizedCollection) GetRandom() int {
	size := len(this.s)
	r := rand.Int() % size
	return int(this.s[r])
}

func main() {
	obj := Constructor()
	obj.Insert(0)
	obj.Remove(0)
	obj.Insert(-1)
	obj.Remove(0)
	fmt.Println(obj.GetRandom())
}
