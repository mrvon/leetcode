package main

import "fmt"

// Double linked-list with sentinel & hash implementaion
type Node struct {
	prev  *Node
	next  *Node
	key   string
	value int
}

type AllOne struct {
	nil_node *Node
	nodemap  map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	nil_node := &Node{nil, nil, "", 0}
	nil_node.next = nil_node
	nil_node.prev = nil_node
	return AllOne{
		nil_node: nil_node,
		nodemap:  make(map[string]*Node),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	n := this.nodemap[key]
	if n != nil {
		n.value++
		// adjust node position
		next := n.next
		if next == this.nil_node {
			return
		}
		if n.value > next.value {
			// exchange n and next
			this.nodemap[n.key], this.nodemap[next.key] =
				this.nodemap[next.key], this.nodemap[n.key]
			n.key, next.key = next.key, n.key
			n.value, next.value = next.value, n.value
		}
	} else {
		n := &Node{
			prev:  nil,
			next:  nil,
			key:   key,
			value: 1,
		}
		this.nodemap[n.key] = n
		n.next = this.nil_node.next
		this.nil_node.next.prev = n
		this.nil_node.next = n
		n.prev = this.nil_node
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
		} else {
			// adjust node position
			prev := n.prev
			if prev == this.nil_node {
				return
			}
			if n.value < prev.value {
				// exchange n and prev
				this.nodemap[n.key], this.nodemap[prev.key] =
					this.nodemap[prev.key], this.nodemap[n.key]
				n.key, prev.key = prev.key, n.key
				n.value, prev.value = prev.value, n.value
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.nil_node.prev.key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.nil_node.next.key
}

func (this *AllOne) Inspect() {
	fmt.Println("----------------------------------------")
	head := this.nil_node.next
	for head != this.nil_node {
		fmt.Printf("%s(%d) ", head.key, head.value)
		head = head.next
	}
	fmt.Println()
}

func main() {
	obj := Constructor()
	obj.Inc("a")
	obj.Inc("b")
	obj.Inc("b")
	obj.Inc("b")
	obj.Inc("b")
	obj.Inspect()
	obj.Dec("b")
	obj.Dec("b")
	obj.Inspect()
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
}
