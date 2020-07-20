package main

import "fmt"

func main() {
	q := new(queue)
	i := item{
		5,
	}
	q.push(i)
	q.push(item{6})
	q.push(item{8})
	q.print()
	fmt.Println(q.pop())
	q.print()
}
