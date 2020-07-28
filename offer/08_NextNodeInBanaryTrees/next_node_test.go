package offer

import (
	"fmt"
	"testing"
)

func TestNextNode(t *testing.T) {
	test1()
}

// -------------------辅助函数-------------------

func createNode(val int) *Node {
	return &Node{Val: val}
}

func buildTree(father, left, right *Node) {
	if father != nil {
		father.Left = left
		father.Right = right

		if left != nil {
			left.Father = father
		}

		if right != nil {
			right.Father = father
		}
	}
}

func printNode(n *Node) {
	if n != nil {
		fmt.Printf("value of this node is: %d\n", n.Val)

		if n.Left != nil {
			fmt.Printf("value of its left child is: %d\n", n.Left.Val)
		}

		if n.Right != nil {
			fmt.Printf("value of its left child is: %d\n", n.Right.Val)
		}
	}

	fmt.Println()
}

func printTree(root *Node) {
	printNode(root)

	if root != nil {
		if root.Left != nil {
			printTree(root.Left)
		}

		if root.Right != nil {
			printTree(root.Right)
		}
	}
}

func test(testName string, in, expect *Node) {
	if NextNode(in) == expect {
		fmt.Println(testName, "Success!")
		return
	}

	fmt.Println(testName, "Failed!")
}

//            8
//        6      10
//      5   7   9  11
func test1() {
	n8 := createNode(8)
	n6 := createNode(6)
	n10 := createNode(10)
	n5 := createNode(5)
	n7 := createNode(7)
	n9 := createNode(9)
	n11 := createNode(11)

	buildTree(n8, n6, n10)
	buildTree(n6, n5, n7)
	buildTree(n10, n9, n11)

	test("test1_1", n8, n9)
	test("test1_2", n6, n7)
	test("test1_3", n10, n11)
	test("test1_4", n5, n6)
	test("test1_5", n7, n8)
	test("test1_6", n9, n10)
	test("test1_7", n11, nil)
}
