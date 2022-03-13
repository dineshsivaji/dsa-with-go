package main

import (
	linkedlist "example.com/dsa/linked_list"
)

func main() {

	x := linkedlist.SinglyLinkedList{}

	x.InsertLast("A")
	x.InsertLast("B")
	x.InsertFirst("Z")
	x.InsertFirst("Y")
	x.Display()
	x.RemoveFirst()
	x.Display()
	x.RemoveLast()
	x.Display()
}
