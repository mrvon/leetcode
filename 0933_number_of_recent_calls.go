package main

import "fmt"

const Range = 3000

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	limit := t - Range
	i := 0
	for i < len(this.queue) && this.queue[i] < limit {
		i++
	}
	this.queue = this.queue[i:]
	return len(this.queue)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}
