/*
Inspired by

https://leetcode.com/problems/implement-magic-dictionary/discuss/107454/Python-without-*26-factor-in-complexity

A word 'apple' has neighbors '*pple', 'a*ple', 'ap*le', 'app*e', 'appl*'. When
searching for a target word like 'apply', we know that a necessary condition is
a neighbor of 'apply' is a neighbor of some source word in our magic dictionary.
If there is more than one source word that does this, then at least one of those
source words will be different from the target word. Otherwise, we need to check
that the source doesn't equal the target.
*/
package main

import "fmt"

type MagicDictionary struct {
	m map[string][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		m: make(map[string][]string),
	}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		for i := 0; i < len(word); i++ {
			neighbor := word[:i] + "*" + word[i+1:]
			this.m[neighbor] = append(this.m[neighbor], word)
		}
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		neighbor := word[:i] + "*" + word[i+1:]
		count := len(this.m[neighbor])
		if count > 1 {
			return true
		} else if count == 1 {
			if this.m[neighbor][0] != word {
				return true
			}
		}
	}
	return false
}

func main() {
	dict := Constructor()
	dict.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(dict.Search("hello"))
	fmt.Println(dict.Search("hallo"))
	fmt.Println(dict.Search("hellk"))
	fmt.Println(dict.Search("hell"))
	fmt.Println(dict.Search("hellok"))
	fmt.Println(dict.Search("leetcode"))
	fmt.Println(dict.Search("leetcoded"))
}
