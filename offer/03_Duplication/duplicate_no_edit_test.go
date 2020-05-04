package offer

import "testing"

func TestDuplicateNoEdit(t *testing.T) {
	// 多个重复的数字
	l := []int{2, 3, 5, 4, 3, 2, 6, 7}
	res := 3
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 一个重复的数字
	l = []int{3, 2, 1, 4, 4, 5, 6, 7}
	res = 4
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 重复的数字是数组中最小的数字
	l = []int{1, 2, 3, 4, 5, 6, 7, 1, 8}
	res = 1
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 重复的数字是数组中最大的数字
	l = []int{1, 7, 3, 4, 5, 6, 8, 2, 8}
	res = 8
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 数组中只有两个数字
	l = []int{1, 1}
	res = 1
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 重复的数字位于数组中间
	l = []int{3, 2, 1, 3, 4, 5, 6, 7}
	res = 3
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 一个数字重复三次
	l = []int{1, 2, 2, 6, 4, 5, 2}
	res = 2
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 没有重复的数字
	l = []int{1, 2, 6, 4, 5, 3}
	res = -1
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}

	// 无效输入
	l = nil
	res = -1
	if DuplicateNoEdit(l) != res {
		t.Error("my shit code got: ", DuplicateNoEdit(l))
	}
}
