package main

import "fmt"

// Double linked-list & hash implementaion
type Node struct {
	prev  *Node
	next  *Node
	key   string
	value int
}

type AllOne struct {
	head    *Node
	tail    *Node
	nodemap map[string]*Node
}

var nil_node *Node = &Node{nil, nil, "", 0}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		head:    nil_node,
		tail:    nil_node,
		nodemap: make(map[string]*Node),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	n := this.nodemap[key]
	if n != nil {
		n.value++
		// adjust node position
		next := n.next
		if next == nil_node {
			return
		}
		if n.value > next.value {
			n.prev, next.prev = next.prev, n.prev
			n.next, next.next = next.next, n.next
		}
		if n == this.head {
			this.head = next
		}
		if next == this.tail {
			this.tail = n
		}
	} else {
		n := &Node{
			prev:  nil_node,
			next:  nil_node,
			key:   key,
			value: 1,
		}
		this.nodemap[key] = n
		this.head.prev = n
		n.next = this.head
		this.head = n
		if this.tail == nil_node {
			this.tail = this.head
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	n := this.nodemap[key]
	if n != nil {
		n.value--
		if n.value == 0 {
			// delete from list
			delete(this.nodemap, key)
			n.prev.next = n.next
			n.next.prev = n.prev
			if n == this.head {
				this.head = n.next
			}
			if n == this.tail {
				this.tail = n.prev
			}
		} else {
			// adjust node position
			prev := n.prev
			if prev == nil_node {
				return
			}
			if n.value > prev.value {
				n.prev, prev.prev = prev.prev, n.prev
				n.next, prev.next = prev.next, n.next
			}
			if prev == this.head {
				this.head = n
			}
			if n == this.tail {
				this.tail = prev
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.tail.key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.head.key
}

func main() {
	obj := Constructor()
	obj.Inc("fuck")
	obj.Inc("fuck")
	obj.Inc("fuck")
	obj.Inc("fool")
	obj.Dec("fool")
	obj.Inc("fun")
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
}
