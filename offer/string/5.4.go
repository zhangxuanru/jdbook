/*
@Time : 2020/4/15 16:18
@Author : zxr
@File : 5.4
@Software: GoLand
*/
package string

//如何判断两个字符串是否为换位字符串
func ComParc(str1, str2 string) bool {
	strCount := make([]int, 256)
	for _, v := range str1 {
		strCount[v]++
	}
	for _, v := range str2 {
		strCount[v]--
	}
	for _, v := range strCount {
		if v != 0 {
			return false
		}
	}
	return true
}
