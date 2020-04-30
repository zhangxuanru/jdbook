/*
@Time : 2020/4/21 18:00
@Author : zxr
@File : 5.22
@Software: GoLand
*/
package string

import (
	"bytes"
)

//如何求相对路径
func GetRelativePath(a, b string) string {
	if a == "" || b == "" {
		return ""
	}
	arr1, arr2 := []rune(a), []rune(b)
	diff1, diff2 := 0, 0
	len1, len2 := len(arr1), len(arr2)
	var path bytes.Buffer

	for i, j := 0, 0; i < len1 && j < len2; {
		if arr1[i] == arr2[j] {
			if arr1[i] == '/' {
				diff1, diff2 = i, j
			}
			i++
			j++
		} else {
			diff1++
			for diff1 < len1 {
				if arr1[diff1] == '/' {
					path.WriteString("../")
				}
				diff1++
			}
			diff2++
			path.WriteString(string(arr2[diff2:]))
			break
		}
	}
	return path.String()
}
