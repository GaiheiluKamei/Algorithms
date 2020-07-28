package offer

import (
	"reflect"
	"testing"
)

func TestConstruct(t *testing.T) {
	// 普通二叉树
	//              1
	//           /     \
	//          2       3
	//         /       / \
	//        4       5   6
	//         \         /
	//          7       8
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root := Construct(pre, mid)
	res := make([]int, 0)
	if !reflect.DeepEqual(preTraverse(root, res), pre) || !reflect.DeepEqual(midTraverse(root, res), mid) {
		t.Error("eat shit in public if this print out...", preTraverse(root, res), midTraverse(root, res))
	}

	// 所有结点都没有右子结点
	//            1
	//           /
	//          2
	//         /
	//        3
	//       /
	//      4
	//     /
	//    5
	pre = []int{1, 2, 3, 4, 5}
	mid = []int{5, 4, 3, 2, 1}
	root = Construct(pre, mid)
	res = make([]int, 0)
	if !reflect.DeepEqual(preTraverse(root, res), pre) || !reflect.DeepEqual(midTraverse(root, res), mid) {
		t.Error("eat shit in public if this print out...", preTraverse(root, res), midTraverse(root, res))
	}

	// 所有结点都没有左子结点
	//            1
	//             \
	//              2
	//               \
	//                3
	//                 \
	//                  4
	//                   \
	//                    5
	pre = []int{1, 2, 3, 4, 5}
	mid = []int{1, 2, 3, 4, 5}
	root = Construct(pre, mid)
	res = make([]int, 0)
	if !reflect.DeepEqual(preTraverse(root, res), pre) || !reflect.DeepEqual(midTraverse(root, res), mid) {
		t.Error("eat shit in public if this print out...", preTraverse(root, res), midTraverse(root, res))
	}

	// 树中只有一个结点
	pre = []int{1}
	mid = []int{1}
	root = Construct(pre, mid)
	res = make([]int, 0)
	if !reflect.DeepEqual(preTraverse(root, res), pre) || !reflect.DeepEqual(midTraverse(root, res), mid) {
		t.Error("eat shit in public if this print out...", preTraverse(root, res), midTraverse(root, res))
	}

	// 完全二叉树
	//              1
	//           /     \
	//          2       3
	//         / \     / \
	//        4   5   6   7
	pre = []int{1, 2, 4, 5, 3, 6, 7}
	mid = []int{4, 2, 5, 1, 6, 3, 7}
	root = Construct(pre, mid)
	res = make([]int, 0)
	if !reflect.DeepEqual(preTraverse(root, res), pre) || !reflect.DeepEqual(midTraverse(root, res), mid) {
		t.Error("eat shit in public if this print out...", preTraverse(root, res), midTraverse(root, res))
	}

	// TODO: 鲁棒性
	// 输入的两个序列不匹配
	pre = []int{1, 2, 4, 5, 3, 6, 7}
	mid = []int{4, 2, 8, 1, 6, 3, 7}
	root = Construct(pre, mid)
	res = make([]int, 0)
}
