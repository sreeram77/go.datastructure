package main

import "fmt"

func main() {

	list := new(linkedList)

	list.Insert(5)
	list.Insert(4)
	list.Insert(3)
	list.Insert(3)

	fmt.Println(list.Find(5))

	list.Print()
}
