/*
@Time : 2020/4/15 16:20
@Author : zxr
@File : 5.4_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

func TestComParc(t *testing.T) {
	str1 := "aaaabbc"
	str2 := "abcbaaa"
	parc := string2.ComParc(str1, str2)
	fmt.Println("是否为换位字符串:", parc)
}
