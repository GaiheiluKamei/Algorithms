/*
	面试题3（二）：不修改数组找出重复的数字。
	题目：在一个长度为n+1的数组里的所有数字都在1～n的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重
	复的数字，但不能修改输入的数组。例如，如果输入长度为8的数组{2,3,5,4,3,2,6,7}，那么对应的输出是重复的数字2或者3。
*/

package offer

// DuplicateNoEdit find duplication number in an array, without edit it.
func DuplicateNoEdit(l []int) int {
	if l == nil || len(l) == 0 {
		return -1
	}

	start := 1
	end := len(l) - 1
	// n := 0
	for end >= start {
		mid := ((end - start) >> 1) + start
		count := countRange(l, start, mid)
		// n++
		// fmt.Println("一次：", n, start, end, mid, count)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func countRange(l []int, start, end int) int {
	if l == nil {
		return 0
	}

	count := 0
	for i := range l {
		if l[i] >= start && l[i] <= end {
			count++
		}
	}
	return count
}


