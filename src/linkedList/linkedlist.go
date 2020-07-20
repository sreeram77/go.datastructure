package main

import "fmt"

type linkedList struct {
	head *node
}

func (l *linkedList) Insert(d int) {

	n := node{
		data: d,
	}
	if l.head == nil {
		l.head = &n
	} else {
		n.next = l.head
		l.head = &n
	}
}

func (l *linkedList) Print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.data)
	}
}

func (l *linkedList) Delete() int {

	delNode := l.head
	l.head = l.head.next
	return delNode.data
}

func (l *linkedList) Find(f int) bool {
	for n := l.head; n != nil; n = n.next {
		if n.data == f {
			return true
		}
	}

	return false
}
