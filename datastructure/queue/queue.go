package datastructure

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	lenght int
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value}

	if q.head == nil {
		q.tail = newNode
		q.head = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.lenght++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var empty T
		return empty, false
	}

	value := q.head.value
	q.head = q.head.next
	q.lenght--

	if q.head == nil {
		q.tail = nil
	}
	return value, true
}

func (q *Queue[T]) Peak() T {
	if q.head == nil {
		var empty T
		return empty
	}

	return q.head.value
}

func (q *Queue[T]) Size() int {
	return q.lenght
}

func (q *Queue[T]) Print() string {
	if q.head == nil {
		return "nil"
	}
	var result string
	current := q.head
	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}
	result += "nil"
	return result
}
