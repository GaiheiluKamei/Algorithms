package list

import "fmt"

type dNode struct {
	prev *dNode
	next *dNode
	data interface{}
}

type dl struct {
	head *dNode
	tail *dNode
	size int
}

func (d *dl) checkIfEmptyAndAdd(node *dNode) bool {
	if d.size == 0 {
		d.head = node
		d.tail = node
		d.size++
		return true
	}

	return false
}

func (d *dl) insertBeginning(data interface{}) {
	newNode := &dNode {
		data: data,
	}

	if !d.checkIfEmptyAndAdd(newNode) {
		head := d.head
		
		// update new node
		newNode.prev = nil
		newNode.next = head

		// update head
		head.prev = newNode

		// update d head
		d.head = newNode
		d.size++
		return
	}

	return
}

func (d *dl) insertEnd(data interface{}) {
	newNode := &dNode {
		data: data,
	}

	if !d.checkIfEmptyAndAdd(newNode) {
		current := d.head
		for i := 0; i < d.size; i++ {
			if current.next == nil {
				// update new node
				newNode.prev = current
				newNode.next = nil

				// update head
				current.next = newNode

				// update d and size
				d.tail = newNode
				d.size++
				break
			}
			current = current.next
		}
	}

	return
}

func (d *dl) insert(position int, data interface{}) {
	newNode := &dNode {
		data: data,
	}

	if !d.checkIfEmptyAndAdd(newNode) {
		current := d.head
		for i := 1; i < d.size; i++ {
			if i == position {
				// update new node
				newNode.prev = current.prev
				newNode.next = current

				// update current
				current.prev.next = newNode
				current.prev = newNode

				// update d and size
				d.size++
				return
			}
			current = current.next
		}
	}

	return
}

func (d *dl) checkIfEmpty() bool {
	if d.size == 0 {
		return true
	}

	return false
}

func (d *dl) deleteFirst() interface{} {
	if !d.checkIfEmpty() {
		head := d.head

		if head.prev == nil {
			// update second node
			d.head = head.next
			d.head.prev = nil
			d.size--

			return head.data
		}
	}

	return -1
}

func (d *dl) deleteLast() interface{} {
	if !d.checkIfEmpty() {
		current := d.head
		for i := 0; i < d.size; i++ {
			if current.next == nil {
				// update the one before last
				current.prev.next = nil
				d.tail = current.prev
				d.size--
				return current.data
			}
			current = current.next
		}
	}

	return -1
}

func (d *dl) delete(position int) interface{} {
	if !d.checkIfEmpty() {
		current := d.head
		for i := 1; i <= position; i++ {
			if current.next == nil && position > i {
				return -1
			} else if i == position {
				// update next one and prev one
				current.next.prev = current.prev
				current.prev.next = current.next
				d.size--
				return current.data
			}

			current = current.next
		}
	}

	return -1
}

func (d *dl) traverse() {
	current := d.head

	for i := 0; i < d.size; i++ {
		fmt.Println("-----newNode no------", i+1)
		fmt.Println("data", current.data)
		fmt.Println("prev", current.prev)
		fmt.Println("next", current.next)
		fmt.Println("------------------")
		current = current.next
	}
}