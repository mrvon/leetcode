package main

import "fmt"

const (
	RED   = 0
	BLACK = 1
)

type rb_node struct {
	key   int
	val   int
	color int

	parent      *rb_node
	left_child  *rb_node
	right_child *rb_node
}

type rb_tree struct {
	root_node *rb_node
	size      int
}

var nil_node = __nil_node()

func assert(result bool) {
	if !result {
		panic(fmt.Sprintf("Assert failed!"))
	}
}

func __nil_node() *rb_node {
	node := &rb_node{}
	node.color = BLACK
	node.parent = node
	node.left_child = node
	node.right_child = node

	return node
}

func new_node() *rb_node {
	node := &rb_node{
		color:       RED,
		left_child:  nil_node,
		right_child: nil_node,
		parent:      nil_node,
	}

	return node
}

func (tree *rb_tree) left_rotate(x *rb_node) {
	y := x.right_child           // Set y
	x.right_child = y.left_child // Turn y's left subtree into x's right subtree

	if y.left_child != nil_node {
		y.left_child.parent = x
	}

	y.parent = x.parent // Link x's parent to y

	if x.parent == nil_node {
		tree.root_node = y
	} else if x == x.parent.left_child {
		x.parent.left_child = y
	} else {
		x.parent.right_child = y
	}

	y.left_child = x // Put x on y's left
	x.parent = y
}

func (tree *rb_tree) right_rotate(y *rb_node) {
	x := y.left_child            // Set x
	y.left_child = x.right_child // Turn x's right subtree into y's left subtree

	if x.right_child != nil_node {
		x.right_child.parent = y
	}

	x.parent = y.parent // Link y's parent to x

	if y.parent == nil_node {
		tree.root_node = x
	} else if y == y.parent.left_child {
		y.parent.left_child = x
	} else {
		y.parent.right_child = x
	}

	x.right_child = y // Put y on x's right
	y.parent = x
}

func (tree *rb_tree) insert_fixup(z *rb_node) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left_child {
			y := z.parent.parent.right_child
			if y.color == RED {
				z.parent.color = BLACK      // CASE 1
				y.color = BLACK             // CASE 1
				z.parent.parent.color = RED // CASE 1
				z = z.parent.parent         // CASE 1
			} else {
				if z == z.parent.right_child {
					z = z.parent        // CASE 2
					tree.left_rotate(z) // CASE 2
				}
				z.parent.color = BLACK             // CASE 3
				z.parent.parent.color = RED        // CASE 3
				tree.right_rotate(z.parent.parent) // CASE 3
			}
		} else {
			y := z.parent.parent.left_child

			if y.color == RED {
				z.parent.color = BLACK      // CASE 1
				y.color = BLACK             // CASE 1
				z.parent.parent.color = RED // CASE 1
				z = z.parent.parent         // CASE 1
			} else {
				if z == z.parent.left_child {
					z = z.parent         // CASE 2
					tree.right_rotate(z) // CASE 2
				}
				z.parent.color = BLACK            // CASE 3
				z.parent.parent.color = RED       // CASE 3
				tree.left_rotate(z.parent.parent) // CASE 3
			}
		}
	}

	tree.root_node.color = BLACK
}

func (tree *rb_tree) __insert(insert_key int, insert_val int) {
	x := tree.root_node
	y := nil_node // trailing pointer of x
	z := new_node()

	tree.size++

	z.key = insert_key
	z.val = insert_val

	for x != nil_node {
		y = x

		if z.key < x.key {
			x = x.left_child
		} else {
			x = x.right_child
		}
	}

	z.parent = y

	if y == nil_node {
		tree.root_node = z
	} else if z.key < y.key {
		y.left_child = z
	} else {
		y.right_child = z
	}

	tree.insert_fixup(z)
}

func (x *rb_node) minimum() *rb_node {
	for x.left_child != nil_node {
		x = x.left_child
	}
	return x
}

func (x *rb_node) maximum() *rb_node {
	for x.right_child != nil_node {
		x = x.right_child
	}
	return x
}

func (x *rb_node) predecessor() *rb_node {
	if x.left_child != nil_node {
		return x.left_child.maximum()
	}

	y := x.parent

	for y != nil_node && x == y.left_child {
		x = y
		y = y.parent
	}

	return y
}

func (x *rb_node) successor() *rb_node {
	if x.right_child != nil_node {
		return x.right_child.minimum()
	}

	y := x.parent

	for y != nil_node && x == y.right_child {
		x = y
		y = y.parent
	}
	return y
}

func (tree *rb_tree) transplant(u *rb_node, v *rb_node) {
	if u.parent == nil_node {
		tree.root_node = v
	} else if u == u.parent.left_child {
		u.parent.left_child = v
	} else {
		u.parent.right_child = v
	}
	v.parent = u.parent
}

func (tree *rb_tree) delete_fixup(x *rb_node) {
	for x != tree.root_node && x.color == BLACK {
		if x == x.parent.left_child {
			w := x.parent.right_child

			if w.color == RED {
				w.color = BLACK            // CASE 1
				x.parent.color = RED       // CASE 1
				tree.left_rotate(x.parent) // CASE 1
				w = x.parent.right_child   // CASE 1
			}

			if w.left_child.color == BLACK && w.right_child.color == BLACK {
				w.color = RED // CASE 2
				x = x.parent  // CASE 2
			} else {
				if w.right_child.color == BLACK {
					w.left_child.color = BLACK // CASE 3
					w.color = RED              // CASE 3
					tree.right_rotate(w)       // CASE 3
					w = x.parent.right_child   // CASE 3
				}

				w.color = x.parent.color    // CASE 4
				x.parent.color = BLACK      // CASE 4
				w.right_child.color = BLACK // CASE 4
				tree.left_rotate(x.parent)  // CASE 4
				x = tree.root_node          // CASE 4
			}
		} else {
			w := x.parent.left_child

			if w.color == RED {
				w.color = BLACK             // CASE 1
				x.parent.color = RED        // CASE 1
				tree.right_rotate(x.parent) // CASE 1
				w = x.parent.left_child     // CASE 1
			}

			if w.right_child.color == BLACK && w.left_child.color == BLACK {
				w.color = RED // CASE 2
				x = x.parent  // CASE 2
			} else {
				if w.left_child.color == BLACK {
					w.right_child.color = BLACK // CASE 3
					w.color = RED               // CASE 3
					tree.left_rotate(w)         // CASE 3
					w = x.parent.left_child     // CASE 3
				}

				w.color = x.parent.color    // CASE 4
				x.parent.color = BLACK      // CASE 4
				w.left_child.color = BLACK  // CASE 4
				tree.right_rotate(x.parent) // CASE 4
				x = tree.root_node          // CASE 4
			}
		}
	}
}

func (tree *rb_tree) __delete(z *rb_node) {
	x := nil_node
	y := z
	y_original_color := y.color

	tree.size--

	if z.left_child == nil_node {
		x = z.right_child
		tree.transplant(z, z.right_child)
	} else if z.right_child == nil_node {
		x = z.left_child
		tree.transplant(z, z.left_child)
	} else {
		y = z.right_child.minimum()
		y_original_color = y.color
		x = y.right_child

		if y.parent == z {
			if x == nil_node {
				x.parent = y // This line is very important
			} else {
				assert(x.parent == y)
			}
		} else {
			tree.transplant(y, y.right_child)
			y.right_child = z.right_child
			y.right_child.parent = y
		}

		tree.transplant(z, y)
		y.left_child = z.left_child
		y.left_child.parent = y
		y.color = z.color
	}

	if y_original_color == BLACK {
		tree.delete_fixup(x)
	}
}

// ----------------------- API -----------------------
func new_tree() *rb_tree {
	tree := &rb_tree{
		root_node: nil_node,
		size:      0,
	}

	return tree
}

func (tree *rb_tree) search(search_key int) *rb_node {
	var __search func(tree *rb_tree, node *rb_node, search_key int) *rb_node

	__search = func(tree *rb_tree, node *rb_node, search_key int) *rb_node {
		if node == nil_node || node.key == search_key {
			return node
		}

		if search_key < node.key {
			return __search(tree, node.left_child, search_key)
		} else {
			return __search(tree, node.right_child, search_key)
		}
	}

	return __search(tree, tree.root_node, search_key)
}

func (tree *rb_tree) insert(insert_key int, insert_val int) {
	node := tree.search(insert_key)
	if node == nil_node {
		tree.__insert(insert_key, insert_val)
	}
}

func (tree *rb_tree) delete(delete_key int) {
	node := tree.search(delete_key)
	if node != nil_node {
		tree.__delete(node)
	}
}

type iterator struct {
	node *rb_node
}

func (tree *rb_tree) iterator() iterator {
	return iterator{
		node: tree.root_node.minimum(),
	}
}

func (i *iterator) has_next() bool {
	return i.node != nil_node
}

func (i *iterator) get_next() *rb_node {
	n := i.node
	i.node = i.node.successor()
	return n
}

type reverse_iterator struct {
	node *rb_node
}

func (tree *rb_tree) reverse_iterator() reverse_iterator {
	return reverse_iterator{
		node: tree.root_node.maximum(),
	}
}

func (i *reverse_iterator) has_next() bool {
	return i.node != nil_node
}

func (i *reverse_iterator) get_next() *rb_node {
	n := i.node
	i.node = i.node.predecessor()
	return n
}

func main() {
	tree := new_tree()

	m := 100
	for i := 1; i < m; i++ {
		tree.insert(i, i*3)
	}

	for i := 1; i < m; i++ {
		node := tree.search(i)
		assert(node != nil)
		assert(node.val == (i * 3))
	}

	for iter := tree.iterator(); iter.has_next(); {
		node := iter.get_next()
		fmt.Printf("%v %v\n", node.key, node.val)
	}

	for iter := tree.reverse_iterator(); iter.has_next(); {
		node := iter.get_next()
		fmt.Printf("%v %v\n", node.key, node.val)
	}

	for i := 1; i < m; i++ {
		tree.delete(i)
	}

	for i := 1; i < m; i++ {
		node := tree.search(i)
		assert(node != nil)
		assert(node.val == 0)
	}
}
