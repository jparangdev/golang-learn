package basic

import (
	"container/list"
	"fmt"
)

// Queue is a data structure that follows the first-in-first-out (FIFO) principle.
type Queue struct {
	v *list.List
}

// push adds a new element to the end of the queue.
// The element is added to the underlying linked list using PushBack method.
func (q *Queue) push(val interface{}) {
	q.v.PushBack(val)
}

// pop removes and returns the element at the front of the queue.
func (q *Queue) pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

// NewQueue creates a new instance of Queue.
// It initializes the underlying list and returns a pointer to the Queue.
func NewQueue() *Queue {
	return &Queue{list.New()}
}

func QueueExample() {
	q := NewQueue()
	q.push(1)
	q.push(2)
	q.push(3)
	fmt.Println(q.pop()) // Output: 1
	fmt.Println(q.pop()) // Output: 2
	fmt.Println(q.pop()) // Output: 3
}
