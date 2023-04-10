package ejercicios34

import "linkedlist"

func MixList(list1, list2 linkedlist.LinkedList[any]) linkedlist.LinkedList[any] {
	var mix_list linkedlist.LinkedList[any]
	n := (list1.Size()+list2.Size())/2 + 1

	for i := 0; i < n; i++ {
		if i < list1.Size() {
			value, _ := list1.Get(i)
			mix_list.Append(value)
		}

		if i < list2.Size() {
			value2, _ := list2.Get(i)
			mix_list.Append(value2)
		}
	}
	return mix_list
}
