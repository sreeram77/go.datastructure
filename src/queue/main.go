package main

import "fmt"

func main() {
	q := new(queue)

	q.enqueue(&item{data: 6})
	q.enqueue(&item{data: 7})
	q.enqueue(&item{data: 8})

	//q.print()

	fmt.Println(*q.dequeue())

	q.print()
}
