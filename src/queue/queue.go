package main

import "fmt"

type queue struct {
	front *item
	rear  *item
}

func (q *queue) enqueue(i *item) {
	// Empty Queue
	if q.front == nil {
		q.front = i
		q.rear = i
		return
	}

	q.rear.next = i
	q.rear = i
}

func (q *queue) dequeue() *item {

	// Queue is empty
	if q.front == nil {
		return nil
	}

	del := q.front
	q.front = q.front.next
	return del
}

func (q *queue) print() {
	for tmp := q.front; tmp != nil; tmp = tmp.next {
		fmt.Println(*tmp)
	}
}
