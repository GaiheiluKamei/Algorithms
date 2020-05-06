/*
	面试题6：从尾到头打印链表
	题目：输入一个链表的头节点，从尾到头反过来打印出每个节点的值。
 */
package offer

// List represent a linked list node.
type Node struct {
	Val int
	Next *Node
}

// DoubleRangeVersion print linked list in reverse order, in double range way.
func DoubleRangeVersion(head *Node) []int {
	if head == nil {
		return nil
	}

	count := 0
	tmpHead := head
	for head != nil {
		count++
		head = head.Next
	}

	res := make([]int, count)
	for i := 0; tmpHead != nil; i++ {
		res[count-1-i] = tmpHead.Val
		tmpHead = tmpHead.Next
	}

	return res
}

// RecursionVersion use recursion finish the problem.
func RecursionVersion(head *Node) []int {
	for head.Next != nil {
		res := RecursionVersion(head.Next)
		res = append(res, head.Val)
		return res
	}
	
	return []int{head.Val}
}