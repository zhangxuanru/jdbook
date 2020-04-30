/*
@Time : 2020/4/15 15:16
@Author : zxr
@File : 5.3
@Software: GoLand
*/
package string

//如何对字符串进行反转
func ReverseStr(str string) string {
	r := []rune(str)
	left, right := 0, len(str)-1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}
