package list

import "fmt"

// cllNode...new node of circular linked list
type cllNode struct {
	data int
	next *cllNode
}

// cll...a linked list with len, head
type cll struct {
	size int
	head *cllNode
}

func (c *cll) len() int {
	current := c.head
	if current == nil {
		return 0
	}

	count := 1
	for current.next != c.head {
		count++
		current = current.next
	}

	return count
}

func (c *cll) traverse() {
	current := c.head
	for i := 0; i < c.size; i++ {
		fmt.Print(current.data, " ")
		current = current.next
	}
}

func (c *cll) checkIfEmptyAndAdd(data int) bool {
	newNode := &cllNode{data: data}
	if c.size == 0 {
		c.head = newNode
		c.head.next = newNode
		c.size++
		return true
	}

	return false
}

func (c *cll) insertBeginning(data int) {
	if !c.checkIfEmptyAndAdd(data) {
		newNode := &cllNode{data: data}

		current := c.head

		// insert on current
		newNode.next = current

		// update tail next to new node
		for current.next != c.head {
			current = current.next
		}

		current.next = newNode

		// update head
		c.head = newNode
		c.size++
	}
}

func (c *cll) insertEnd(data int) {
	if !c.checkIfEmptyAndAdd(data) {
		newNode := &cllNode{data: data}

		current := c.head
		for current.next != c.head {
			current = current.next
		}

		current.next = newNode
		newNode.next = c.head
		c.size++
	}
}

func (c *cll) insert(data int, position int) {
	if !c.checkIfEmptyAndAdd(data) {
		if position == 1 {
			c.insertBeginning(data)
			return
		}

		newNode := &cllNode{data: data}
		current := c.head
		count := 1

		for {
			if current.next == nil && position - 1 > count {
				break
			}

			if count == position -1 {
				newNode.next = current.next
				current.next = newNode
				c.size++
				break
			}

			current = current.next
			count++
		}
		
	}
}

func (c *cll) checkIfEmpty() bool {
	if c.size == 0 {
		return true
	}

	return false
}

func (c *cll) deleteFirst() int {
	if !c.checkIfEmpty() {
		current := c.head
		delData := current.data
		if c.size == 1 {
			c.head = nil
			c.size--
			return delData
		}
		tmp := c.head

		// update head
		c.head = current.next

		for current.next != tmp {
			current = current.next
		}
		// now current is tail node
		current.next = c.head
		c.size--
		return delData
	}

	return -1
}

func (c *cll) deleteLast() int {
	if !c.checkIfEmpty() {
		current := c.head
		if c.size == 1 {
			return c.deleteFirst()
		}

		for current.next.next != c.head {
			current = current.next
		}

		delData := current.next.data

		current.next = c.head
		c.size--
		return delData
	}

	return -1
}

func (c *cll) delete(position int) int {
	if !c.checkIfEmpty() {
		current := c.head

		if position > c.size {
			return -1
		}

		if position == 1 {
			return c.deleteFirst()
		}

		if position == c.size {
			return c.deleteLast()
		}

		count := 1
		delData := current.data
		for {
			if count + 1 == position {
				delData = current.next.data
				break
			}

			current = current.next
			count++
		}

		current.next = current.next.next
		c.size--
		return delData
	}

	return -1
}

/*
func main() {
	cll := cll{}
	fmt.Println("Size: ", cll.len())
	cll.insertBeginning(100)
	cll.insertBeginning(90)
	cll.insertBeginning(70)
	cll.insertBeginning(60)
	cll.insertBeginning(50)
	cll.insertBeginning(40)
	cll.insertBeginning(30)

	cll.insertBeginning(19) // Update dll

	cll.traverse()
	cll.deleteFirst()
	cll.traverse()

	cll.delete(3)
	cll.traverse()

	cll.insert(66, 2)
	cll.traverse()
	fmt.Println("Size: ", cll.len())
}
*/