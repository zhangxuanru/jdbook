/*
@Time : 2020/4/16 19:03
@Author : zxr
@File : 5.10_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

func TestLong(t *testing.T) {
	str := "cbbcabacbca"
	palindrome := string2.LongestPalindrome(str)
	fmt.Println("palindrome:", palindrome)
	palindrome2 := string2.LongestPalindrome2(str)
	fmt.Println("palindrome2:", palindrome2)
}

//最长回文子串
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	longest := s[0:1]
	for i := 1; i < len(s); i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			fmt.Println("i=", i, "rightStep:", rightStep)
			for p, q := i-1, i+rightStep; p >= 0 && q < len(s) && s[p] == s[q]; {
				fmt.Println("p=", p, "q=", q)
				if q-p+1 > len(longest) {
					longest = s[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return longest
}
