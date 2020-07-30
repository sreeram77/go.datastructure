package main

import "fmt"

func main() {
	s := new(stack)
	i := item{
		5,
	}
	s.push(i)
	s.push(item{6})
	s.push(item{8})
	s.print()
	fmt.Println(s.pop())
	s.print()
}
