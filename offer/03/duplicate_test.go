package rn

import "testing"

func TestDuplicate(t *testing.T) {
	// 重复的数字是数组中最小的数字
	l := []int{2, 1, 3, 1, 4}
	res := 1
	if Duplicate(l) != res {
		t.Error("my shit code got: ", Duplicate(l))
	}

	// 重复的数字是数组中最大的数字
	l = []int{2, 4, 3, 1, 4}
	res = 4
	if Duplicate(l) != res {
		t.Error("my shit code got: ", Duplicate(l))
	}

	// 数组中存在多个重复的数字
	l = []int{2, 4, 2, 1, 4}
	res = 2
	if Duplicate(l) != res {
		t.Error("my shit code got: ", Duplicate(l))
	}

	// 没有重复的数字
	l = []int{2, 1, 3, 0, 4}
	res = -1
	if Duplicate(l) != res {
		t.Error("my shit code got: ", Duplicate(l))
	}

	// 无效的输入
	l = []int{}
	res = -1
	if Duplicate(l) != res {
		t.Error("my shit code got: ", Duplicate(l))
	}
}
