package main

import "fmt"

type stack struct {
	items []item
}

func (s *stack) push(i item) {
	s.items = append(s.items, i)
}

func (s *stack) pop() item {
	last := len(s.items) - 1
	delItem := s.items[last]
	s.items = s.items[0:last]

	return delItem
}

func (s *stack) print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i].data)
	}
}
