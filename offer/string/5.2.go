/*
@Time : 2020/4/13 17:11
@Author : zxr
@File : 5.2
@Software: GoLand
*/
package string

//如何求两个字符串的最长公共子串
func MaxSubLen(str1, str2 string) string {
	chs1 := len(str1)
	chs2 := len(str2)

	maxLength := 0 //记录最大长度
	end := 0       //记录最大长度的结尾位置

	rows := 0
	cols := chs2 - 1
	for rows < chs1 {
		i, j := rows, cols
		length := 0 //记录长度
		for i < chs1 && j < chs2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}
	return str1[(end - maxLength + 1):(end + 1)]

}

func MaxSubStr2(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	if n1 == 0 || n2 == 0 {
		return ""
	}
	maxLength := 0
	end := 0

	left := 0
	right := n2 - 1
	for left < n1 {
		i, j := left, right
		length := 0
		for i < n1 && j < n2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if right > 0 {
			right--
		} else {
			left++
		}
	}
	start := end - maxLength + 1
	return str1[start : end+1]
}
