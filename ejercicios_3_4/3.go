package ejercicios34

import (
	"linkedlist"
)

func UnirList(list1, list2 linkedlist.LinkedList[any]) linkedlist.LinkedList[any] {
	listf := list1
	for i := 0; i < list2.Size(); i++ {
		value, _ := list2.Get(i)
		listf.Append(value)
	}

	return listf
}
