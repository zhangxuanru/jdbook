/*
@Time : 2020/4/9 17:47
@Author : zxr
@File : 5.1_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

//如何求一个字符串的所有排列
func TestPermutationStr(t *testing.T) {
	str := "abc"
	//	str = "DiLNSr"
	//str = "qwe"
	//permutationStr := string2.PermutationStr(str)
	//fmt.Println(permutationStr)

	//success
	permutation2 := string2.Permutation2(str)
	fmt.Println(permutation2)

	string2.Permutation3([]rune(str), 0)
}
