package main

/*
def find(node, key):
    for char in key:
        if char in node.children:
            node = node.children[char]
        else:
            return None
    return node

def insert(root : node, s : string):
    node = root
    i    = 0
    n    = length(s)

    while i < n:
        if node.child(s[i]) != nil:
            node = node.child(s[i])
            i = i + 1
        else:
            break

    (* append new nodes, if necessary *)
    while i < n:
        node.child(s[i]) = new node
        node = node.child(s[i])
        i = i + 1
*/

type Trie struct {
	children []*Trie
}

func tonumber(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a')
	} else if b >= 'A' && b <= 'Z' {
		return int(b - 'A' + 26)
	} else {
		panic("Invalid character.")
	}
}

func create() *Trie {
	new := &Trie{}
	new.children = make([]*Trie, 26*2)
	return new
}

func (node *Trie) find(s string) *Trie {
	for i := 0; i < len(s); i++ {
		node = node.children[tonumber(s[i])]
		if node == nil {
			return nil
		}
	}
	return node
}

func (node *Trie) findvalue(s string) string {
	var val []byte
	for i := 0; i < len(s); i++ {
		node = node.children[tonumber(s[i])]
		if node == nil {
			return ""
		}
		val = append(val, s[i])
	}
	return string(val)
}

func (node *Trie) insert(s string) {
	i := 0
	n := len(s)

	for i < n {
		next := node.children[tonumber(s[i])]
		if next == nil {
			break
		} else {
			node = next
			i++
		}
	}

	for i < n {
		new := create()
		node.children[tonumber(s[i])] = new
		node = node.children[tonumber(s[i])]
		i++
	}
}

func assert(result bool) {
	if !result {
		panic("Assert failed.")
	}
}

func main() {
	trie := create()

	trie.insert("hello")
	trie.insert("world")
	trie.insert("quick")
	trie.insert("sort")
	trie.insert("Definitive")

	assert(trie.findvalue("k") == "")
	assert(trie.findvalue("quicksort") == "")
	assert(trie.findvalue("hel") == "hel")
	assert(trie.findvalue("hello") == "hello")
	assert(trie.findvalue("helloworld") == "")
	assert(trie.findvalue("Definition") == "")
	assert(trie.findvalue("Definitive") == "Definitive")
	assert(trie.findvalue("definitive") == "")

	assert(trie.find("Definitive") != nil)
	assert(trie.find("definitive") == nil)
}
