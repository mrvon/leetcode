package main

import "fmt"

type RingBuffer struct {
	buffer []string
	head   int
	tail   int
}

func (rb *RingBuffer) Init() {
	rb.buffer = make([]string, 10)
	rb.head = 0
	rb.tail = 0
}

func (rb *RingBuffer) full() bool {
	bufLen := len(rb.buffer)
	next := (rb.tail + 1) % bufLen
	return rb.head == next
}

func (rb *RingBuffer) expand() {
	if !rb.full() {
		return
	}
	bufLen := len(rb.buffer)
	newLen := bufLen * 2
	newBuf := make([]string, newLen)
	newTail := 0
	for !rb.Empty() {
		newBuf[newTail] = rb.Pop()
		newTail++
	}
	rb.buffer = newBuf
	rb.head = 0
	rb.tail = newTail
}

func (rb *RingBuffer) Push(s string) {
	rb.expand()
	bufLen := len(rb.buffer)
	rb.buffer[rb.tail] = s
	rb.tail = (rb.tail + 1) % bufLen
}

func (rb *RingBuffer) Pop() string {
	if rb.Empty() {
		assert(false, "RingBuffer is empty")
	}
	head := rb.head
	bufLen := len(rb.buffer)
	rb.head = (rb.head + 1) % bufLen
	return rb.buffer[head]
}

func (rb *RingBuffer) Empty() bool {
	return rb.head == rb.tail
}

func assert(result bool, msg string) {
	if !result {
		panic(fmt.Sprintf("Assert failed!: %s\n", msg))
	}
}

func main() {
	rb := &RingBuffer{}
	rb.Init()
	assert(rb.Empty(), "rb is empty")
	rb.Push("hello")
	rb.Push("world")
	assert(!rb.Empty(), "rb is not empty")
	assert(rb.Pop() == "hello", "expect hello")
	assert(rb.Pop() == "world", "expect hello")
	for i := 0; i < 1024; i++ {
		v := fmt.Sprintf("VALUE_%d", i)
		rb.Push(v)
	}
	for i := 0; i < 1000; i++ {
		e := fmt.Sprintf("VALUE_%d", i)
		v := rb.Pop()
		assert(v == e, fmt.Sprintf("should be %s, find %s", e, v))
	}
}
