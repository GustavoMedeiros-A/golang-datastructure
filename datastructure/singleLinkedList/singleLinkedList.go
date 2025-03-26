package singlelinkedlist

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{value: value, next: list.head}
	list.head = newNode
}

func (list *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Search(value int) (*Node, bool) {
	current := list.head
	for current != nil {
		if current.value == value {
			return current, true
		}
		current = current.next
	}

	return nil, false
}

func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		return
	}
	if list.head.value == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}
