package offer

import (
	"reflect"
	"testing"
)

var (
	n1 = &Node{5, n2}
	n2 = &Node{4, n3}
	n3 = &Node{3, n4}
	n4 = &Node{2, n5}
	n5 = &Node{1, nil}
)

func TestDoubleRangeVersion(t *testing.T) {
	// 多个节点的链表
	if !reflect.DeepEqual(DoubleRangeVersion(n1), []int{1, 2, 3, 4, 5}) {
		t.Error("【double】sorry for my shit code: ", DoubleRangeVersion(n1))
	}

	if !reflect.DeepEqual(RecursionVersion(n1), []int{1, 2, 3, 4, 5}) {
		t.Error("【recursion】sorry for my shit code: ", RecursionVersion(n1))
	}

	// 只有一个节点的链表
	if !reflect.DeepEqual(DoubleRangeVersion(n5), []int{1}) {
		t.Error("【double】sorry for my shit code: ", DoubleRangeVersion(n1))
	}

	if !reflect.DeepEqual(RecursionVersion(n5), []int{1}) {
		t.Error("【recursion】sorry for my shit code: ", RecursionVersion(n1))
	}
}
