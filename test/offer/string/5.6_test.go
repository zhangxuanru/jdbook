/*
@Time : 2020/4/15 17:37
@Author : zxr
@File : 5.6_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

//如何对由大小写字母组成的字符数组排序
func TestSortString(t *testing.T) {
	str := "AbcDef"
	s := string2.SortString(str)
	fmt.Println("string:", s)
}
