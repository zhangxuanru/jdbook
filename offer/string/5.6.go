/*
@Time : 2020/4/15 17:36
@Author : zxr
@File : 5.6
@Software: GoLand
*/
package string

//如何对由大小写字母组成的字符数组排序
func SortString(str string) string {
	r := []rune(str)
	left := 0
	end := len(r) - 1
	for left < end {
		for r[left] >= 'a' && r[left] <= 'z' && end > left {
			left++
		}
		for r[end] >= 'A' && r[end] <= 'Z' && end > left {
			end--
		}
		r[left], r[end] = r[end], r[left]
	}
	return string(r)
}
