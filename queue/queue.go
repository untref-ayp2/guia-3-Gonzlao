package queue

import (
	"errors"
	"linkedlist"
)

type Queue[T comparable] struct {
	list linkedlist.LinkedList[T]
}

func (q *Queue[T]) Enqueue(v T) {
	q.list.Append(v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("la cola esta vacia")
	}
	head, _ := q.Front()
	q.list.Remove(head)
	return head, nil
}

func (q *Queue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("la cola esta vacia")
	}
	head, _ := q.list.Get(0)
	return head, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Size() == 0
}
