package main

import (
	ll "example.com/dsa/linked_list"
)

func main() {

	// x := linkedlist.SinglyLinkedList{}

	// x.InsertLast("A")
	// x.InsertLast("B")
	// x.InsertFirst("Z")
	// x.InsertFirst("Y")
	// x.Display()
	// x.RemoveFirst()
	// x.Display()
	// x.RemoveLast()
	// x.Display()
	// x.InsertAt("Q", 1)
	// x.Display()
	x := ll.SinglyCircularLinkedList{}
	x.InsertLast("12")
	x.InsertLast("14")
	x.Display()
	x.RemoveFirst()
	x.Display()
}
