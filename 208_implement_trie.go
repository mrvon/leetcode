// Trie, Prefix tree

package main

import "fmt"

type Trie struct {
	children []*Trie
	end      bool
}

/** Initialize your data structure here. */
func Create() *Trie {
	new := &Trie{}
	new.children = make([]*Trie, 26) // only lowercase letter
	return new
}

func Constructor() Trie {
	t := Create()
	return *t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}

	index := int(word[0] - 'a')
	child := this.children[index]
	if child == nil {
		child = Create()
		this.children[index] = child
	}
	child.Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}

	index := int(word[0] - 'a')
	child := this.children[index]
	if child == nil {
		return false
	} else {
		return child.Search(word[1:])
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	index := int(prefix[0] - 'a')
	child := this.children[index]
	if child == nil {
		return false
	} else {
		return child.StartsWith(prefix[1:])
	}
}

func main() {
	trie := Constructor()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println(trie.Search(""), trie.StartsWith(""))
	fmt.Println(trie.Search("hello"), trie.StartsWith("hello"))
	fmt.Println(trie.Search("world"), trie.StartsWith("world"))
	fmt.Println(trie.Search("wor"), trie.StartsWith("wor"))
	fmt.Println(trie.Search("game"), trie.StartsWith("game"))
}
