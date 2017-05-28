package main

import "fmt"

type Object struct {
	object_id int
	x         int
	y         int
}

type ListNode struct {
	object_id int
	x         int
	y         int
	next      *ListNode
}

type Scene struct {
	x_list  *ListNode
	y_list  *ListNode
	obj_id  int
	objects map[int]*Object
}

const (
	VisualRange = 5
)

func create_scene() *Scene {
	return &Scene{
		x_list:  &ListNode{},
		y_list:  &ListNode{},
		obj_id:  0,
		objects: make(map[int]*Object),
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func (this *Scene) create_object(x int, y int) int {
	this.obj_id++
	object := &Object{
		object_id: this.obj_id,
		x:         x,
		y:         y,
	}
	this.objects[this.obj_id] = object
	return this.obj_id
}

func (this *Scene) __near_set(object *Object) map[int]bool {
	// X list
	x_set := make(map[int]bool)
	node := this.x_list
	for node.next != nil {
		next := node.next
		if object.x <= next.x {
			if next.x-object.x <= VisualRange {
				x_set[next.object_id] = true
			}
		} else {
			if object.x-next.x <= VisualRange {
				x_set[next.object_id] = true
			} else {
				break
			}
		}
		node = next
	}
	// Y list
	y_set := make(map[int]bool)
	node = this.y_list
	for node.next != nil {
		next := node.next
		if object.y <= next.y {
			if next.y-object.y <= VisualRange {
				y_set[next.object_id] = true
			}
		} else {
			if object.y-next.y <= VisualRange {
				y_set[next.object_id] = true
			} else {
				break
			}
		}
		node = next
	}
	n_set := make(map[int]bool)
	for id := range x_set {
		if id != object.object_id && y_set[id] {
			n_set[id] = true
		}
	}
	return n_set
}

func (this *Scene) send_enter_message(watcher *Object, object *Object) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Enter\n",
		watcher.object_id, watcher.x, watcher.y, object.object_id, object.x, object.y,
	)
}

func (this *Scene) send_leave_message(watcher *Object, object *Object) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Leave\n",
		watcher.object_id, watcher.x, watcher.y, object.object_id, object.x, object.y,
	)
}

func (this *Scene) send_move_message(watcher *Object, object *Object, old_x int, old_y int) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Move to (%d,%d) \n",
		watcher.object_id, watcher.x, watcher.y, object.object_id, old_x, old_y, object.x, object.y,
	)
}

func (this *Scene) enter(object_id int) {
	object := this.objects[object_id]
	if object == nil {
		return
	}

	fmt.Printf("Object[%d](%d,%d) Enter\n", object.object_id, object.x, object.y)

	this.__enter(object)
	near_set := this.__near_set(object)
	for id := range near_set {
		this.send_enter_message(this.objects[id], object)
	}
}

func (this *Scene) __enter(object *Object) {
	// X list
	node := this.x_list
	for node.next != nil {
		next := node.next
		if object.x <= next.x {
			break
		} else {
			node = next
		}
	}
	new_node := &ListNode{
		object_id: object.object_id,
		x:         object.x,
		y:         object.y,
		next:      node.next,
	}
	node.next = new_node
	// Y list
	node = this.y_list
	for node.next != nil {
		next := node.next
		if object.y <= next.y {
			break
		} else {
			node = next
		}
	}
	new_node = &ListNode{
		object_id: object.object_id,
		x:         object.x,
		y:         object.y,
		next:      node.next,
	}
	node.next = new_node
}

func (this *Scene) leave(object_id int) {
	object := this.objects[object_id]
	if object == nil {
		return
	}

	fmt.Printf("Object[%d](%d,%d) Leave\n", object.object_id, object.x, object.y)

	near_set := this.__near_set(object)
	for id := range near_set {
		this.send_leave_message(this.objects[id], object)
	}
	this.__leave(object)
}

func (this *Scene) __leave(object *Object) {
	// X list
	node := this.x_list
	for node.next != nil {
		next := node.next
		if object.object_id == next.object_id {
			node.next = next.next
			break
		} else {
			node = next
		}
	}
	// Y list
	node = this.y_list
	for node.next != nil {
		next := node.next
		if object.object_id == next.object_id {
			node.next = next.next
			break
		} else {
			node = next
		}
	}
}

func (this *Scene) move(object_id int, new_x int, new_y int) {
	object := this.objects[object_id]
	if object == nil {
		return
	}

	old_x := object.x
	old_y := object.y

	fmt.Printf(
		"Object[%d](%d,%d) Move to (%d,%d)\n",
		object.object_id, old_x, old_y, new_x, new_y,
	)

	near_set_before := this.__near_set(object)
	this.__leave(object)
	object.x = new_x
	object.y = new_y
	this.__enter(object)
	near_set_after := this.__near_set(object)

	for id := range near_set_before {
		if near_set_after[id] {
			this.send_move_message(this.objects[id], object, old_x, old_y)
		} else {
			this.send_leave_message(this.objects[id], object)
		}
	}
	for id := range near_set_after {
		if !near_set_before[id] {
			this.send_enter_message(this.objects[id], object)
		}
	}
}

func (this *Scene) print() {
	fmt.Println("-------- X --------")
	node := this.x_list
	for node.next != nil {
		next := node.next
		fmt.Printf("[%d](%d,%d)\n", next.object_id, next.x, next.y)
		node = next
	}
	fmt.Println("-------- Y --------")
	node = this.y_list
	for node.next != nil {
		next := node.next
		fmt.Printf("[%d](%d,%d)\n", next.object_id, next.x, next.y)
		node = next
	}
	fmt.Print("\n")
}

func main() {
	s := create_scene()
	obj_1 := s.create_object(1, 2)
	obj_2 := s.create_object(2, 1)
	obj_3 := s.create_object(3, 3)
	s.enter(obj_1)
	s.enter(obj_2)
	s.enter(obj_3)
	s.leave(obj_1)
	s.move(obj_3, 0, 0)
	// s.print()
}
