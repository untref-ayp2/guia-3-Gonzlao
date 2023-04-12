package tests

import (
	"stack"
	"testing"
)

func TestPush(t *testing.T) { // se modifico ya que el slice esta encapsulado. Para poder comparar cada elemento se agregaron los Pops correspondientes.
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)
	s.Push(3)

	n3, _ := s.Pop()
	n2, _ := s.Pop()
	n1, _ := s.Pop()

	if n1 != 1 || n2 != 2 || n3 != 3 {
		t.Error("Error en Push")
	}
}

func TestPop(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()

	if err != nil || v != 3 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 2 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 1 {
		t.Error("Error en Pop")
	}

	_, err = s.Pop()

	if err == nil {
		t.Error("Error en Pop")
	}
}
