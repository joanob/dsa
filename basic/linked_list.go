package dsa

import "errors"

// Linked lists are a series of nodes connected by pointers
// Each node contains a reference to the next node
// Linked lists are more efficient for insertion and deletion than arrays as nodes are not sequential in memory
// Reading a linked list is less efficient than reading an array because access is sequential

type linkedList[T comparable] struct {
	head *linkedListNode[T]
}

type linkedListNode[T comparable] struct {
	data T
	next *linkedListNode[T]
}

func CreateLinkedList[T comparable]() *linkedList[T] {
	return &linkedList[T]{head: nil}
}

// Insert data in the end of the linked list
func (list *linkedList[T]) InsertLast(data T) {
	newNode := &linkedListNode[T]{
		data: data,
		next: nil,
	}
	if list.head == nil {
		list.head = newNode
		return
	}
	node := list.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

// Insert data into the selected position of the linked list
func (list *linkedList[T]) Insert(pos int, data T) error {
	if pos == 0 {
		newNode := linkedListNode[T]{
			data: data,
			next: list.head,
		}
		list.head = &newNode
		return nil
	}

	node := list.head
	currentPos := 1
	for node != nil {
		if currentPos == pos {
			newNode := linkedListNode[T]{
				data: data,
				next: node.next,
			}
			node.next = &newNode
			return nil
		}
		currentPos++
		node = node.next
	}
	return errors.New("pos not in list")
}

// Delete selected position from the linked list
func (list *linkedList[T]) Delete(pos int) {
	if list.head == nil {
		return
	}

	if pos == 0 {
		list.head = list.head.next
		return
	}

	currentPos := 1
	node := list.head
	for node != nil {
		if pos == currentPos {
			if node.next != nil {
				node.next = node.next.next
			}
			return
		}
		currentPos++
		node = node.next
	}
}

// Get data from selected position or an error
func (list *linkedList[T]) Get(pos int) (T, error) {
	if list.head == nil {
		var val T
		return val, errors.New("")
	}

	node := list.head
	currentPos := 0
	for currentPos <= pos && node != nil {
		if currentPos == pos {
			return node.data, nil
		}
		currentPos++
		node = node.next
	}

	var val T
	return val, errors.New("")
}

// Get the linked list length
func (list *linkedList[T]) Length() int {
	len := 0
	node := list.head
	for node != nil {
		node = node.next
		len++
	}
	return len
}

// Search linked list.
// Returns the position of the data or an error
func (list *linkedList[T]) Search(data T) (int, error) {
	pos := 0
	node := list.head
	for node != nil {
		if data == node.data {
			return pos, nil
		}
		node = node.next
		pos++
	}
	return 0, errors.New("")
}
