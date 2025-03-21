package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) Display() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}

	fmt.Println("nil")
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	// If deleting head
	if l.head.value == value {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil // if list becomes empty
		}
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.value == value {
			if current.next == l.tail {
				l.tail = current
			}
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func main() {
	ll := LinkedList{}

	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)
	ll.Insert(50)

	ll.Display()

	println("Original LinkedList")

	ll.Delete(20)

	ll.Display()

	println("After LinkedList")
}
