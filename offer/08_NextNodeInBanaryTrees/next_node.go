/*
	面试题8: 二叉树的下一个节点
	题目：给定一棵二叉树和其中一个节点，如何找出中序遍历序列的下一个节点？
	🌲中的节点除了有两个分别指向左右子节点的指针以外，还有一个指向父节点的指针。

	分析：
	1. 如果一个节点有右子树，那么它的下一个节点就是它的右子树中的最左子节点
	2. 如果一个节点没有右子树，且是它父节点的左子节点，那么它的下一个节点就是它的父节点
	3. 如果一个节点没有右子树，且不是它父节点的左子节点，我们可以沿着指向父节点的指针一直
	向上遍历，直到找到一个是它父节点的左子节点的节点，如果这样的节点存在，那么这个节点的
	父节点就是我们要找的下一个节点
*/

package offer

// Node is a binary tree node.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Father *Node
}

// NextNode return a binary tree's next node.
func NextNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		next := n.Right
		for next.Left != nil {
			next = next.Left
		}

		return next
	}

	for n.Father != nil {
		if n.Father.Left == n {
			return n.Father
		}

		n = n.Father
	}

	return nil
}
