package main

import "fmt"

const (
	RED   = 0
	BLACK = 1
)

type interval struct {
	low  int
	high int
}

type i_node struct {
	i     interval
	m     int
	color int

	parent      *i_node
	left_child  *i_node
	right_child *i_node
}

type i_tree struct {
	root_node *i_node
	size      int
}

var nil_node = __nil_node()

func assert(result bool) {
	if !result {
		panic(fmt.Sprintf("Assert failed!"))
	}
}

func overlap(ix interval, iy interval) bool {
	if ix.high <= iy.low || iy.high <= ix.low {
		return false
	} else {
		return true
	}
}

func __nil_node() *i_node {
	node := &i_node{}
	node.color = BLACK
	node.parent = node
	node.left_child = node
	node.right_child = node

	return node
}

func new_node() *i_node {
	node := &i_node{
		color:       RED,
		left_child:  nil_node,
		right_child: nil_node,
		parent:      nil_node,
	}

	return node
}

func (tree *i_tree) left_rotate(x *i_node) {
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

func (tree *i_tree) right_rotate(y *i_node) {
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

func (tree *i_tree) insert_fixup(z *i_node) {
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

func (tree *i_tree) __insert(i interval) {
	x := tree.root_node
	y := nil_node // trailing pointer of x
	z := new_node()

	tree.size++

	z.i = i

	for x != nil_node {
		y = x

		if z.i.low < x.i.low {
			x = x.left_child
		} else {
			x = x.right_child
		}
	}

	z.parent = y

	if y == nil_node {
		tree.root_node = z
	} else if z.i.low < y.i.low {
		y.left_child = z
	} else {
		y.right_child = z
	}

	tree.insert_fixup(z)
}

func (x *i_node) minimum() *i_node {
	for x.left_child != nil_node {
		x = x.left_child
	}
	return x
}

func (x *i_node) maximum() *i_node {
	for x.right_child != nil_node {
		x = x.right_child
	}
	return x
}

func (x *i_node) predecessor() *i_node {
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

func (x *i_node) successor() *i_node {
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

func (tree *i_tree) transplant(u *i_node, v *i_node) {
	if u.parent == nil_node {
		tree.root_node = v
	} else if u == u.parent.left_child {
		u.parent.left_child = v
	} else {
		u.parent.right_child = v
	}
	v.parent = u.parent
}

func (tree *i_tree) delete_fixup(x *i_node) {
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

func (tree *i_tree) __delete(z *i_node) {
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

func new_tree() *i_tree {
	tree := &i_tree{
		root_node: nil_node,
		size:      0,
	}

	return tree
}

func (tree *i_tree) search(i interval) *i_node {
	x := tree.root_node
	for x != nil_node && !overlap(i, x.i) {
		if x.left_child != nil_node && x.left_child.m >= i.low {
			x = x.left_child
		} else {
			x = x.right_child
		}
	}
	return x
}

func (tree *i_tree) insert(i interval) {
	node := tree.search(i)
	if node == nil_node {
		tree.__insert(i)
	}
}

func (tree *i_tree) delete(i interval) {
	node := tree.search(i)
	if node != nil_node {
		tree.__delete(node)
	}
}

type iterator struct {
	node *i_node
}

func (tree *i_tree) iterator() iterator {
	return iterator{
		node: tree.root_node.minimum(),
	}
}

func (i *iterator) has_next() bool {
	return i.node != nil_node
}

func (i *iterator) get_next() *i_node {
	n := i.node
	i.node = i.node.successor()
	return n
}

type reverse_iterator struct {
	node *i_node
}

func (tree *i_tree) reverse_iterator() reverse_iterator {
	return reverse_iterator{
		node: tree.root_node.maximum(),
	}
}

func (i *reverse_iterator) has_next() bool {
	return i.node != nil_node
}

func (i *reverse_iterator) get_next() *i_node {
	n := i.node
	i.node = i.node.predecessor()
	return n
}

func main() {
	tree := new_tree()

	tree.insert(interval{1, 2})
	tree.insert(interval{3, 4})

	for i := tree.iterator(); i.has_next(); {
		fmt.Println(i.get_next().i)
	}
}
