package slicelist

import (
	"errors"
	"fmt"
)

type SliceList[T comparable] struct {
	list []T
}

// O(1)
func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{}
}

// O(1)
func (l *SliceList[T]) Append(value T) {
	l.list = append(l.list, value)
}

// O(1)
func (l *SliceList[T]) Prepend(value T) {
	l.list = append([]T{value}, l.list...)
}

// O(n)
func (l *SliceList[T]) InsertAt(value T, position int) {
	if position < 0 {
		return
	}
	if position == 0 {
		l.Prepend(value)
		return
	}
	if position == (l.Size() - 1) {
		l.Append(value)
		return
	}

	for i := l.Size() - 1; i >= position; i-- {
		if i == (l.Size() - 1) {
			l.list = append(l.list, l.list[i])
		}
		l.list[i+1] = l.list[i]
	}
	l.list[position] = value
}

// O(n)
func (l *SliceList[T]) Remove(value T) {
	if l.Size() == 0 {
		return // no hay nada que eliminar
	}
	if l.list[0] == value {
		l.list = l.list[1:]
		return
	}
	index_value := l.Search(value) //O(n)
	l.list = append(l.list[:index_value], l.list[index_value+1:]...)
}

// O(1)
func (l *SliceList[T]) String() string {
	return fmt.Sprintf("%v", l.list)
}

// O(n)
func (l *SliceList[T]) Search(value T) int {
	if l.Size() == 0 {
		return -1
	}
	position := 0
	for position < l.Size() {
		if l.list[position] == value {
			return position
		}
		position++
	}
	return -1
}

// O(1)
func (l *SliceList[T]) Get(position int) (T, error) {
	if l.Size() == 0 {
		var t T
		return t, errors.New("Lista vacía")
	}
	if position < 0 || position > l.Size() {
		var t T
		return t, errors.New("Posición inválida")
	}

	return l.list[position], nil
}

// O(1)
func (l *SliceList[T]) Size() int {
	return len(l.list)
}
