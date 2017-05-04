/*
Double linked-list with sentinel & hash implementation

The principle is maintain a double linked-list, and pushes all element with same
value n into a single hash bucket(n).

	bucket(1) <-> bucket(2) <-> bucket(5) <-> bucket(7)

INC operation, O(1), if key is exists and it's value is n, just move the key to
the bucket(n+1). Otherwise, push it into bucket(1).

DEC operation, O(1), if key is exists and it's value is n, just move the key to
the bucket(n-1). Otherwise, do nothing.

MIN operation, O(1), random get a key from head node.

MAX operation, O(1), random get a key from tail node.
*/
package main

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	keys  map[string]bool
	value int
}

type AllOne struct {
	nil_node *Node
	nodemap  map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	nil_node := &Node{
		prev:  nil,
		next:  nil,
		keys:  make(map[string]bool),
		value: 0,
	}
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
		next := n.next
		delete(n.keys, key)

		if next.value == n.value+1 {
			next.keys[key] = true
			this.nodemap[key] = next
		} else {
			o := &Node{
				prev:  nil,
				next:  nil,
				keys:  make(map[string]bool),
				value: n.value + 1,
			}

			n.next = o
			o.prev = n
			o.next = next
			next.prev = o

			o.keys[key] = true
			this.nodemap[key] = o
		}

		if len(n.keys) == 0 {
			// delete node n
			n.prev.next = n.next
			n.next.prev = n.prev
		}
	} else {
		if head := this.nil_node.next; head.value == 1 {
			head.keys[key] = true
			this.nodemap[key] = head
		} else {
			o := &Node{
				prev:  nil,
				next:  nil,
				keys:  make(map[string]bool),
				value: 1,
			}
			o.next = this.nil_node.next
			this.nil_node.next.prev = o
			this.nil_node.next = o
			o.prev = this.nil_node

			o.keys[key] = true
			this.nodemap[key] = o
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	n := this.nodemap[key]
	if n != nil {
		prev := n.prev
		delete(n.keys, key)

		if n.value == 1 {
			delete(this.nodemap, key)
		} else {
			if prev.value == n.value-1 {
				prev.keys[key] = true
				this.nodemap[key] = prev
			} else {
				o := &Node{
					prev:  nil,
					next:  nil,
					keys:  make(map[string]bool),
					value: n.value - 1,
				}

				prev.next = o
				o.prev = prev
				o.next = n
				n.prev = o

				o.keys[key] = true
				this.nodemap[key] = o
			}
		}

		if len(n.keys) == 0 {
			// delete node n
			n.prev.next = n.next
			n.next.prev = n.prev
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	for key := range this.nil_node.prev.keys {
		return key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	for key := range this.nil_node.next.keys {
		return key
	}
	return ""
}

func (this *AllOne) Inspect() {
	fmt.Println("----------------------------------------")
	head := this.nil_node.next
	for head != this.nil_node {
		fmt.Printf("%v(%d) ", head.keys, head.value)
		head = head.next
	}
	fmt.Println()
}

func test1() {
	obj := Constructor()
	obj.Inc("hello")
	obj.Inspect()
	obj.Inc("world")
	obj.Inspect()
	obj.Inc("leet")
	obj.Inspect()
	obj.Inc("code")
	obj.Inspect()
	obj.Inc("DS")
	obj.Inspect()
	obj.Inc("leet")
	obj.Inspect()
	fmt.Println(obj.GetMaxKey())
}

func test2() {
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

func test3() {
	obj := Constructor()
	obj.Inc("a")
	obj.Inc("b")
	obj.Dec("b")
	obj.Dec("b")
	obj.Inspect()
	obj.Inc("b")
	obj.Inspect()
	obj.Dec("b")
	obj.Dec("b")
	obj.Inspect()
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
}

func main() {
	// test1()
	// test2()
	test3()
}
