package list

import (
	"fmt"
)
/*
- 基本操作
  - 遍历链表
  - 插入元素
  - 删除元素
*/

// Node represent a node
type Node struct {
	data interface{}
	next *Node
}

// Linked represent a linked list (head)
type Linked struct {
	head *Node // stores the memory address of the head or the first node of the linked list.
	size int // linked list length
}

func (l *Linked) traverse() error {
	if l.head == nil {
		return fmt.Errorf("traverse: list is empty")
	}

	current := l.head
	for current != nil {
		fmt.Printf("%v ->\n", current.data)
		current = current.next
	}

	return nil
}

/*
To find the length of a linked list, one would normally have to traverse all nodes, and count them. This operation takes O(n) time. In order to have this operation done in constant time, we can add the size field in linked list structure definition.
*/

// Length with extra size field
func (l *Linked) Length() int {
	return l.size
}

func (l *Linked) length() int {
	size := 0
	current := l.head
	for current != nil {
		size++
		current = current.next
	}

	return size
}

/*
Insertion into a singly-linked list has three cases:
- Inserting a new node before the head(at the beginning)
- Inserting a new node after thetail (at the end of the list)
- Inserting a new node at the middle of the list(random location)
*/

func (l *Linked) insertBeginning(data interface{}) {
	node := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.size++
}

func (l *Linked) insertEnd(data interface{}) {
	node := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	l.size++
}

func (l *Linked) insert(position int, data interface{}) error {
	if position < 0 || position > l.size + 1 {
		return fmt.Errorf("index out of bounds")
	}

	node := &Node{
		data: data,
	}

	var prev, current *Node
	current = l.head

	for position > 1 {
		prev = current
		current = current.next
		position--
	}

	if prev != nil {
		prev.next = node
		node.next = current
	} else {
		node.next = current
		l.head = node
	}

	l.size++
	return nil
}

func (l *Linked) deleteFirst() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("deleteFirst: List is empty")
	}

	data := l.head.data
	l.head = l.head.next
	l.size--
	return data, nil
}

func (l *Linked) deleteLast() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("deleteFirst: List is empty")
	}

	var prev *Node
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}

	if prev != nil {
		prev.next = nil
	} else {
		l.head = nil
	}

	l.size--
	return current.data, nil
}

func (l *Linked) delete(position int) (interface{}, error) {
	if position < 0 || position > l.size + 1 {
		return nil, fmt.Errorf("delete: index out of bounds")
	}

	var prev *Node
	current := l.head
	pos := 0

	if position == 1 {
		l.head = l.head.next
	} else {
		for pos != position - 1 {
			prev = current
			current = current.next
			pos++
		}

		if current.next != nil {
			prev.next = current.next
		}
	}

	l.size--
	return current.data, nil
}
