/*
@Time : 2020/4/22 19:28
@Author : zxr
@File : relative_path
@Software: GoLand
*/
package t

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

//如何求相对路径
func TestGetRelativePath(t *testing.T) {
	a := "/qihoo/app/a/b/c/d/new.c"
	b := "/qihoo/app/1/2/test.c"
	path := getRelativePath(a, b)
	fmt.Println("path:", path)

	path2 := getRelativePath2(a, b)
	fmt.Println("path:", path2)
}

func getRelativePath(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		return ""
	}
	arr1 := []rune(path1)
	arr2 := []rune(path2)
	dept1 := 0
	dept2 := 0
	var buf bytes.Buffer
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		if arr1[i] == arr2[j] {
			if arr1[i] == '/' {
				dept1 = i
				dept2 = j
			}
			i++
			j++
		} else {
			dept1++
			for dept1 < len(arr1) {
				if arr1[dept1] == '/' {
					buf.WriteString("../")
				}
				dept1++
			}
			dept2++
			buf.WriteString(string(arr2[dept2:]))
			break
		}
	}
	return buf.String()
}

func getRelativePath2(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		return ""
	}
	arr1 := strings.Split(path1, "/")
	arr2 := strings.Split(path2, "/")
	dept := 0
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		if arr1[i] == arr2[i] {
			dept++
		} else {
			break
		}
	}
	prefix := ""
	if len(arr1)-dept-1 <= 0 {
		prefix = "./"
	} else {
		for i := 0; i < len(arr1)-dept-1; i++ {
			prefix += "../"
		}
	}
	if len(arr2)-dept > 0 {
		prefix += strings.Join(arr2[dept:], "/")
	}
	return prefix
}
