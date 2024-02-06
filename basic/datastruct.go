package basic

import (
	"container/list"
	"fmt"
)

func DataStructExample() {
	ListExample()
}

// ListExample is a function that demonstrates the usage of the list package in Go.
//
// It creates a new doubly-linked list,
// inserts elements into it, and then prints the values in both forward and backward order.
func ListExample() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushBack(1)
	v.InsertAfter(3, e4)
	v.InsertBefore(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
