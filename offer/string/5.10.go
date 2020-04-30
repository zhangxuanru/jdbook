/*
@Time : 2020/4/20 16:59
@Author : zxr
@File : 5.10
@Software: GoLand
*/
package string

//最长回文子串
func LongestPalindrome(str string) string {
	if len(str) < 2 {
		return str
	}
	s := str[0:1]
	for i := 1; i < len(str); i++ {
		for r := 0; r < 2; r++ {
			for p, q := i-1, i+r; p >= 0 && q < len(str) && str[p] == str[q]; {
				if q-p+1 > len(s) {
					s = str[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return s
}

//方法二
func LongestPalindrome2(str string) string {
	n := len(str)
	var maxLen int
	var maxStr string
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			l := isPalindrome(str[i:j])
			if l > maxLen {
				maxLen = l
				maxStr = str[i:j]
			}
		}
	}
	return maxStr
}

func isPalindrome(str string) int {
	var newStr string
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	if newStr == str {
		return len(str)
	}
	return 0
}
