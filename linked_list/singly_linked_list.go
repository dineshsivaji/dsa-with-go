package linkedlist

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	size int
}

func (sll *SinglyLinkedList) InsertLast(elem string) {
	newNode := Node{data: elem}
	if sll.head == nil {
		sll.head = &newNode
	} else {
		currentNode := sll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
	}
	sll.size++
}

func (sll *SinglyLinkedList) InsertFirst(elem string) {
	newNode := Node{data: elem}
	if sll.head == nil {
		sll.head = &newNode
	} else {
		newNode.next = sll.head
		sll.head = &newNode
	}
	sll.size++
}

func (sll *SinglyLinkedList) RemoveFirst() {
	if sll.size == 0 {
		fmt.Println("Empty linked list")
		return
	}
	currentNode := sll.head
	if currentNode.next != nil {
		sll.head = currentNode.next
	} else {
		sll.head = nil
	}
	sll.size--
}

func (sll *SinglyLinkedList) RemoveLast() {
	if sll.size == 0 {
		fmt.Println("Empty linked list")
		return
	}
	currentNode := sll.head
	prevNode := sll.head
	for currentNode.next != nil {
		prevNode = currentNode
		currentNode = currentNode.next
	}
	prevNode.next = nil
	sll.size--
}

func (sll SinglyLinkedList) Display() {
	currentNode := sll.head
	for currentNode != nil {
		if currentNode.next != nil {
			fmt.Printf("%s->", currentNode.data)
		} else {
			fmt.Printf("%s\n", currentNode.data)
		}
		currentNode = currentNode.next
	}
}
