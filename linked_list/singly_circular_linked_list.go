package linkedlist

import (
	"fmt"
)

type SinglyCircularLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (scll SinglyCircularLinkedList) GetSize() int {
	return scll.size
}

func (scll *SinglyCircularLinkedList) InsertLast(elem string) {
	newNode := &Node{data: elem}
	if scll.GetSize() == 0 {
		scll.head = newNode
		newNode.next = newNode
		scll.tail = newNode
	} else {
		lastNode := scll.tail
		lastNode.next = newNode
		newNode.next = scll.head
		scll.tail = newNode
	}
	scll.size++
}

func (scll *SinglyCircularLinkedList) RemoveFirst() {
	if scll.GetSize() == 0 {
		print("Empty linked list")
		return
	}
	currentNode := scll.head
	tailNode := scll.tail
	if currentNode == tailNode {
		scll.head = nil
		scll.tail = nil
	} else {
		secondNode := currentNode.next
		scll.head = secondNode
		tailNode.next = secondNode
	}
	scll.size--
}

func (scll SinglyCircularLinkedList) Display() {
	currentNode := scll.head
	tailNode := scll.tail
	for currentNode != nil {
		if currentNode == tailNode {
			fmt.Println(currentNode.data)
			break
		} else {
			fmt.Printf("%s->", currentNode.data)
		}
		currentNode = currentNode.next
	}
}
