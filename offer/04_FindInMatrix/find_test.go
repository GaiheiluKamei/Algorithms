package offer

import "testing"

func TestFind(t *testing.T) {
	// 要查找的数在数组中
	l := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	goal := 7
	if !Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 要查找的数不在数组中
	goal = 5
	if Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 要查找的数是数组中最小的数字
	goal = 1
	if !Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 要查找的数是数组中最大的数字
	goal = 15
	if !Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 要查找的数比数组中最小的数字还小
	goal = 0
	if Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 要查找的数比数组中最大的数字还大
	goal = 16
	if Find(l, goal) {
		t.Error("your code is amazing!")
	}

	// 鲁棒性测试，输入nil
	l = nil
	if Find(l, goal) {
		t.Error("your code is amazing!")
	}
}
