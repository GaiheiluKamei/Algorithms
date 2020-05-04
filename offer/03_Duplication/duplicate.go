/*
	面试题3（一）：找出数组中重复的数字。
	题目：在一个长度为n的数组里的所有数字都在0～n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不
	知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应
	的输出是重复的数字2或者3。
 */

package offer

// Duplicate find duplication number in an array, return -1 if not found.
func Duplicate(l []int) int {
	if len(l) <= 0 {
		return -1
	}

	for i := range l {
		for l[i] != i {
			if l[i] == l[l[i]] {
				return l[i]
			}

			l[i], l[l[i]] = l[l[i]], l[i]
		}
	}

	return -1
}
