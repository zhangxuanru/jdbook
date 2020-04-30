/*
@Time : 2020/4/15 15:18
@Author : zxr
@File : 5.3_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

//如何对字符串进行反转
func TestReverseStr(t *testing.T) {
	str := "abcdefg"
	reverseStr := string2.ReverseStr(str)
	fmt.Println(reverseStr)
}
