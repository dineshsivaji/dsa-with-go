package main

import (
	ll "example.com/dsa/linked_list"
)

func main() {

	x := ll.DoublyLinkedList{}
	x.InsertFirst("12")
	x.InsertFirst("13")
	x.InsertFirst("14")
	x.Display()
	x.InsertLast("26")
	x.InsertLast("27")
	x.InsertLast("28")
	x.Display()
	x.RemoveLast()
	x.Display()
	x.RemoveLast()
	x.Display()
	x.RemoveFirst()
	x.Display()
	x.RemoveFirst()
	x.Display()
	x.RemoveFirst()
	x.Display()
	x.RemoveFirst()
	x.Display()
	x.RemoveLast()
	x.Display()

}
