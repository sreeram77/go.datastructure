package main

import "fmt"

type queue struct {
	items []item
}

func (q *queue) push(i item) {
	q.items = append(q.items, i)
}

func (q *queue) pop() item {
	last := len(q.items) - 1
	delItem := q.items[last]
	q.items = q.items[0:last]

	return delItem
}

func (q *queue) print() {
	for i := len(q.items) - 1; i >= 0; i-- {
		fmt.Println(q.items[i].data)
	}
}
