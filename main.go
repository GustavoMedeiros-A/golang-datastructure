package main

import (
	"fmt"

	"github.com/GustavoMedeiros-A/golang-datastructure/datastructure"
)

func main() {
	q := &datastructure.Queue[int]{}

	// Enqueue some values
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println("Print:", q.Print())
	var removedValue, isRemoved = q.Dequeue()
	fmt.Println("Print:", q.Print())
	println(removedValue, isRemoved)
	var currentHead = q.Peak()
	fmt.Println("Head:", currentHead)
	// Print size
	fmt.Println("Size:", q.Size())
}
