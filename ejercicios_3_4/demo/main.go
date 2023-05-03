package main

import (
	"ejercicios_3_4/ejercicios"
	"fmt"
	"linkedlist"
)

func main() {
	list1 := linkedlist.NewLinkedList[any]()
	list2 := linkedlist.NewLinkedList[any]()

	list1.Append(1)
	list1.Append(2)
	list1.Append(3)

	list2.Append(6)
	list2.Append(9)
	list2.Append(5)
	list2.Append(7)

	listf := ejercicios.UnirList(*list1, *list2)

	fmt.Println(listf.String())

	mix_list := ejercicios.MixList(*list1, *list2)

	fmt.Println(mix_list.String())
}
