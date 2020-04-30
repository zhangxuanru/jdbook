/*
@Time : 2020/4/20 12:10
@Author : zxr
@File : longCommPre_test
@Software: GoLand
*/
package string

import (
	"fmt"
	"strings"
	"testing"
)

//查找字符串数组中的最长公共前缀
func TestLongCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	prefix := LongestCommonPrefix(strs)
	fmt.Println(prefix)
	prefix2 := LongestCommonPrefix2(strs)
	fmt.Println("prefix2:", prefix2)
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	for i := 0; i < len(s); i++ {
		for _, str := range strs[1:] {
			if len(str) < i || str[i] != s[i] {
				return s[:i]
			}
		}
	}
	return ""
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		index := strings.Index(strs[i], s)
		if index != 0 {
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
