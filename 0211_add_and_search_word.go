// Trie solution
package main

import "fmt"

type RegTrie struct {
	children []*RegTrie
	end      bool
}

/* Initialize your data structure here. */
func Create() *RegTrie {
	new := &RegTrie{}
	new.children = make([]*RegTrie, 26) // only lowercase letter
	return new
}

/* Inserts a word into the RegTrie. */
func (this *RegTrie) Insert(word string) {
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

/* Returns if the word is in the trie. */
func (this *RegTrie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}

	if word[0] == '.' {
		for _, child := range this.children {
			if child != nil && child.Search(word[1:]) {
				return true
			}
		}
		return false
	} else {
		index := int(word[0] - 'a')
		child := this.children[index]
		if child == nil {
			return false
		} else {
			return child.Search(word[1:])
		}
	}
}

type WordDictionary struct {
	t *RegTrie
}

/*
 * Initialize your data structure here.
 */
func Constructor() WordDictionary {
	return WordDictionary{
		t: Create(),
	}
}

/*
 * Adds a word into the data structure.
 */
func (this *WordDictionary) AddWord(word string) {
	this.t.Insert(word)
}

/*
 * Returns if the word is in the data structure. A word could contain the dot
 * character '.' to represent any one letter.
 */
func (this *WordDictionary) Search(word string) bool {
	return this.t.Search(word)
}

func assert(expect bool, result bool) {
	if expect != result {
		panic(fmt.Sprintf("Assert failed!"))
	}
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	assert(false, obj.Search("pad"))
	assert(true, obj.Search("bad"))
	assert(true, obj.Search(".ad"))
	assert(true, obj.Search("b.."))
	assert(false, obj.Search("ba"))
	assert(false, obj.Search("b..."))
}
