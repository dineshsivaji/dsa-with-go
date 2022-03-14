package linkedlist

import (
	"fmt"
)

type DNode struct {
	data string
	next *DNode
	prev *DNode
}

type DoublyLinkedList struct {
	head *DNode
	size int
}

func (dll DoublyLinkedList) GetSize() int {
	return dll.size
}

func (dll *DoublyLinkedList) InsertFirst(elem string) {
	newNode := &DNode{data: elem}
	if dll.GetSize() == 0 {
		dll.head = newNode
		newNode.next = nil
		newNode.prev = nil
	} else {
		currentNode := dll.head
		newNode.next = currentNode
		newNode.prev = nil
		dll.head = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) InsertLast(elem string) {
	newNode := &DNode{data: elem}
	if dll.GetSize() == 0 {
		dll.head = newNode
		newNode.next = nil
		newNode.prev = nil
	} else {
		currentNode := dll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
		newNode.prev = currentNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) RemoveFirst() {
	if dll.GetSize() == 0 {
		print("Empty list")
		return
	}
	currentNode := dll.head
	if currentNode.next == nil {
		dll.head = nil
	} else {
		secondNode := currentNode.next
		secondNode.prev = nil
		dll.head = secondNode
	}
	dll.size--
}

func (dll *DoublyLinkedList) RemoveLast() {
	if dll.GetSize() == 0 {
		print("Empty list")
		return
	}
	currentNode := dll.head
	var prevNode *DNode
	if currentNode.next == nil {
		dll.head = nil
	} else {
		for currentNode.next != nil {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		prevNode.next = nil
	}
	dll.size--
}

func (dll DoublyLinkedList) Display() {
	currentNode := dll.head
	for currentNode != nil {
		if currentNode.next != nil {
			fmt.Printf("%s->", currentNode.data)
		} else {
			fmt.Println(currentNode.data)
		}
		currentNode = currentNode.next
	}
}
