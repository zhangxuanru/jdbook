/*
@Time : 2020/4/21 18:10
@Author : zxr
@File : 5.22_test
@Software: GoLand
*/
package string

import (
	"bytes"
	"fmt"
	string2 "jdbook/offer/string"
	"strings"
	"testing"
)

//如何求相对路径
func TestGetRelativePath(t *testing.T) {
	a := "/qihoo/app/a/b/c/d/new.c"
	b := "/qihoo/app/1/2/test.c"
	path := string2.GetRelativePath(a, b)
	fmt.Println("relative path:", path)
	s := relativePath(a, b)
	fmt.Println("relative path2:", s)
	path2 := relativePath2(a, b)
	fmt.Println("relative path2:", path2)
}

//方法一
func relativePath(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		return ""
	}
	path1Arr := strings.Split(path1, "/")
	path2Arr := strings.Split(path2, "/")
	dept := 0
	for i := 0; i < len(path1Arr) && i < len(path2Arr); i++ {
		if path1Arr[i] == path2Arr[i] {
			dept++
		} else {
			break
		}
	}
	prefix := ""
	if len(path1Arr)-dept-1 <= 0 {
		prefix = "./"
	} else {
		for i := len(path1Arr) - dept - 1; i > 0; i-- {
			prefix += "../"
		}
	}
	if len(path2Arr)-dept > 0 {
		prefix += strings.Join(path2Arr[dept:], "/")
	}
	fmt.Println(path1Arr)
	fmt.Println(path2Arr)
	return prefix
}

//方法二
func relativePath2(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		return ""
	}
	path1R, path2R := []rune(path1), []rune(path2)
	path1Len, path2Len := len(path1R), len(path2R)
	diff1, diff2 := 0, 0
	var buf bytes.Buffer
	for i, j := 0, 0; i < path1Len && j < path2Len; {
		if path1R[i] == path2R[j] {
			if path1R[i] == '/' {
				diff1, diff2 = i, j
			}
			i++
			j++
		} else {
			diff1++
			for diff1 < path1Len {
				if path1R[diff1] == '/' {
					buf.WriteString("../")
				}
				diff1++
			}
			diff2++
			buf.WriteString(string(path2R[diff2:]))
			break
		}
	}
	return buf.String()
}
