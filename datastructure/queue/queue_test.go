package datastructure

import "testing"

func TestQueue(t *testing.T) {
	t.Run("it should return an empty queue", func(t *testing.T) {
		q := Queue[int]{}

		if size := q.Size(); size != 0 {
			t.Error("Expected size to be 0, got", size)
		}

		if _, ok := q.Dequeue(); ok {
			t.Error("Expected Dequeue on empty queue to return false, and return", ok)
		}

		if peek := q.Peak(); peek != 0 {
			t.Errorf("Expected Peek on empty queue to return zero value, got %d", peek)
		}
	})

	t.Run("it should increase the size of the queue", func(t *testing.T) {
		q := Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)
		if size := q.Size(); size != 2 {
			t.Error("Expected size to be 2, got", size)
		}
	})

	t.Run("it should peak the first element", func(t *testing.T) {
		q := Queue[int]{}
		q.Enqueue(1)
		if peek := q.Peak(); peek != 1 {
			t.Error("Expected Peek to return 1, got", peek)
		}
	})

	t.Run("it should dequeue elements in FIFO order", func(t *testing.T) {
		q := Queue[int]{}
		q.Enqueue(10)
		q.Enqueue(20)

		val, ok := q.Dequeue()
		if !ok || val != 10 {
			t.Errorf("Expected Dequeue to return 10, got %d", val)
		}

		val, ok = q.Dequeue()
		if !ok || val != 20 {
			t.Errorf("Expected Dequeue to return 20, got %d", val)
		}

		if size := q.Size(); size != 0 {
			t.Errorf("Expected size 0, got %d", size)
		}
	})

}
