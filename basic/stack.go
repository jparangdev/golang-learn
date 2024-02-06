package basic

import (
	"container/list"
	"fmt"
)

// Stack is a type that represents a stack data structure.
// It is implemented using a linked list from the container/list package.
type Stack struct {
	v *list.List
}

// push pushes a value onto the stack.
func (q *Stack) push(val interface{}) {
	q.v.PushBack(val)
}

// pop removes and returns the element at the top of the stack.
// If the stack is empty, it returns nil.
func (q *Stack) pop() interface{} {
	back := q.v.Back()
	if back != nil {
		return q.v.Remove(back)
	}
	return nil
}

// NewStack returns a new instance of Stack with an empty list as its underlying data structure.
func NewStack() *Stack {
	return &Stack{list.New()}
}

func StackExample() {
	s := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop()) // Output: 3
	fmt.Println(s.pop()) // Output: 2
	fmt.Println(s.pop()) // Output: 1
}
