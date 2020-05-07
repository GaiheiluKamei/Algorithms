/*
	面试题7：重建二叉树
	题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如，输入
	前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建【如书中图】所示的二叉树。

	前序遍历：先访问根节点，再访问左子节点，最后访问右子节点。
	中序遍历：先访问左子节点，再访问根节点，最后访问右子节点。
	后序遍历：先访问左子节点，再访问右子节点，最后访问根节点。
*/
package offer

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func Construct(pre, mid []int) *BinaryTree {
	if len(pre) != len(mid) || len(pre) == 0 {
		return nil
	}

	root := &BinaryTree{Val: pre[0]}
	index := 0
	for i := range mid {
		if mid[i] == root.Val {
			index = i
			break
		}
	}

	if index > 0 {
		root.Left = Construct(pre[1:index+1], mid[:index])
	}

	if len(pre)-index-1 > 0 {
		root.Right = Construct(pre[index+1:], mid[index+1:])
	}
	
	return root
}

// 前序遍历
func preTraverse(root *BinaryTree, res []int) []int {
	if root != nil {
		res = append(res, root.Val)
		res = preTraverse(root.Left, res)
		res = preTraverse(root.Right, res)
	}

	return res
}

// 中序遍历
func midTraverse(root *BinaryTree, res []int) []int {
	if root != nil {
		res = midTraverse(root.Left, res)
		res = append(res, root.Val)
		res = midTraverse(root.Right, res)
	}

	return res
}

