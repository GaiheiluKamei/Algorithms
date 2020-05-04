/*
	面试题5：替换空格
	题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如，输入"We are happy."，则
	输出"we%20are%20happy."。
 */
package offer

// ReplaceBlank replace blank to %20 in a string.
func ReplaceBlank(str []byte, length int) {
	if str == nil {
		return
	}

	count := 0
	// 统计字符出现的数目，计算替换后字符串的长度
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}

	for index, newIndex := length-1, length + count*2 - 1; index >=0 && newIndex >= 0; {
		if str[index] == ' ' {
			str[newIndex] = '0'
			newIndex--
			str[newIndex] = '2'
			newIndex--
			str[newIndex] = '%'
			newIndex--
			index--
		} else {
			str[newIndex] = str[index]
			newIndex--
			index--
		}
	}
}

