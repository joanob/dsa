package dsa

import "errors"

type queue[T any] struct {
	list []T
}

func CreateQueue[T any]() *queue[T] {
	return &queue[T]{
		list: make([]T, 0, 8),
	}
}

func (queue *queue[T]) Enqueue(data T) {
	queue.list = append(queue.list, data)
}

// Get first value
func (queue *queue[T]) Dequeue() (T, error) {
	var val T
	if len(queue.list) == 0 {
		return val, errors.New("")
	}
	val = queue.list[0]
	queue.list = queue.list[1:]
	return val, nil
}

// Get first value without deleting it
func (queue *queue[T]) Peek() (T, error) {
	var val T
	if len(queue.list) == 0 {
		return val, errors.New("")
	}
	val = queue.list[0]
	return val, nil
}
