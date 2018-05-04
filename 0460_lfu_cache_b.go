package main

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	key      int // The key of the item.
	value    int // The value of the item.
	priority int // The priority of the item in the queue.
	time     int // The timestamp of the item to be used.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority < pq[j].priority {
		return true
	} else if pq[i].priority == pq[j].priority {
		return pq[i].time <= pq[j].time
	} else {
		return false
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type LFUCache struct {
	hash      map[int]*Item
	heap      PriorityQueue
	capacity  int
	current   int
	timestamp int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		hash:      make(map[int]*Item),
		heap:      PriorityQueue{},
		capacity:  capacity,
		current:   0,
		timestamp: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	item := this.hash[key]
	if item == nil {
		return -1
	}

	this.timestamp++

	item.time = this.timestamp
	(&this.heap).update(item, item.priority+1)

	return item.value
}

func (this *LFUCache) Set(key int, value int) {
	if this.capacity == 0 {
		return
	}

	// try set
	item := this.hash[key]
	if item != nil {
		this.timestamp++
		item.time = this.timestamp
		item.value = value
		(&this.heap).update(item, item.priority+1)
		return
	}

	// just insert
	if this.current >= this.capacity {
		i := heap.Pop(&this.heap).(*Item)
		delete(this.hash, i.key)
	} else {
		this.current++
	}

	this.timestamp++

	item = &Item{
		key:      key,
		value:    value,
		priority: 0,
		time:     this.timestamp,
	}

	this.hash[key] = item
	heap.Push(&this.heap, item)
}

func assert(b bool) {
	if !b {
		panic("Assert failed!")
	}
}

func test_1() {
	capacity := 2
	obj := Constructor(capacity)
	obj.Set(1, 1)
	obj.Set(2, 2)
	assert(obj.Get(1) == 1)
	obj.Set(3, 3)
	assert(obj.Get(2) == -1)
	assert(obj.Get(3) == 3)
	obj.Set(4, 4)
	assert(obj.Get(1) == -1)
	assert(obj.Get(3) == 3)
	assert(obj.Get(4) == 4)
	obj.Set(3, 1)
	assert(obj.Get(3) == 1)
}

func test_2() {
	capacity := 2
	obj := Constructor(capacity)
	obj.Set(2, 1)
	obj.Set(1, 1)
	obj.Set(2, 3)
	obj.Set(4, 1)
	assert(obj.Get(1) == -1)
	assert(obj.Get(2) == 3)
}

func test_3() {
	capacity := 0
	obj := Constructor(capacity)
	obj.Set(0, 0)
	assert(obj.Get(0) == -1)
}

func main() {
	test_1()
	test_2()
	test_3()
}
