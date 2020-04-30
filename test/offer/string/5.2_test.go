/*
@Time : 2020/4/13 17:12
@Author : zxr
@File : 5.2_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

//如何求两个字符串的最长公共子串
func TestMaxSubLen(t *testing.T) {
	str1 := "abccade"
	str2 := "dgcadde"
	subLen := string2.MaxSubLen(str1, str2)
	fmt.Println(subLen)
	fmt.Println("---------------------------------")
	subStr2 := string2.MaxSubStr2(str1, str2)
	fmt.Println(subStr2)
}

func TestA(t *testing.T) {
	str1 := "abccade"
	str2 := "dgcadde"
	s := getMaxSub(str1, str2)
	fmt.Println("getmaxsub:", s)
}

func getMaxSub(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	if n1 == 0 || n2 == 0 {
		return ""
	}
	maxLength := 0
	end := 0

	left := 0
	right := n2 - 1

	for left < n1 {
		i, j := left, right
		length := 0
		for i < n1 && j < n2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				maxLength = length
				end = i
			}
			i++
			j++
		}
		if right > 0 {
			right--
		} else {
			left++
		}
	}
	start := end - maxLength + 1
	return str1[start : end+1]
}
