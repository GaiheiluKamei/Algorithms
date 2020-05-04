/*
	面试题4：二维数组中的查找
	题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照都上到下递增的
	顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 */
package offer

// Find determines whether the specified number exists in the matrix.
func Find(l [][]int, goal int) bool {
	 if l == nil || len(l) == 0 || len(l[0]) == 0 {
	 	return false
	 }

	 row := 0
	 column := len(l[row]) - 1
	 for row < len(l) && column >= 0 {
	 	if l[row][column] == goal {
	 		return true
		}

		if l[row][column] > goal {
			column--
		} else {
			row++
		}
	 }

	 return false
}


