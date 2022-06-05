package main

import (
	"container/list"
	"fmt"
)

type LRUItem struct {
	Value int
	Node  *list.Element
}

type LRUCache struct {
	M map[int]*LRUItem
	L *list.List
	C int
}

func Constructor(capacity int) LRUCache {
	assert(capacity > 0)
	return LRUCache{
		M: make(map[int]*LRUItem),
		L: list.New(),
		C: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	item, existed := this.M[key]
	if existed {
		this.L.MoveToBack(item.Node)
		return item.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	item, existed := this.M[key]
	if existed {
		item.Value = value
		this.Get(key)
		return
	}
	node := this.L.PushBack(key)
	item = &LRUItem{
		Value: value,
		Node:  node,
	}
	this.M[key] = item
	if this.L.Len() <= this.C {
		return
	}
	front := this.L.Front()
	if front == nil {
		return
	}
	this.L.Remove(front)
	delete(this.M, front.Value.(int))
}

func assert(result bool) {
	if !result {
		panic(fmt.Sprintf("Assert failed!"))
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert(cache.Get(1) == 1)  // returns 1
	cache.Put(3, 3)            // evicts key 2
	assert(cache.Get(2) == -1) // returns -1 (not found)
	cache.Put(4, 4)            // evicts key 1
	assert(cache.Get(1) == -1) // returns -1 (not found)
	assert(cache.Get(3) == 3)  // returns 3
	assert(cache.Get(4) == 4)  // returns 4
}
