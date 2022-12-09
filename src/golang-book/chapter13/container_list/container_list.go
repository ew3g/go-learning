package main

import (
	"container/list"
	"fmt"
)

func main() {
	// List is a LinkedList
	// Each node has a pointer to the next and previous element
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

}