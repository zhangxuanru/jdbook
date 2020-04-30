/*
@Time : 2020/4/24 11:00
@Author : zxr
@File : 6.2
@Software: GoLand
*/
package matins

//如何判断一个数是否是2的N次方
func IsPow2N(n int) bool {
	if n < 1 {
		return false
	}
	for i := 1; i <= n; {
		if i == n {
			return true
		}
		i <<= 1
	}
	return false
}

func IsPow2N2(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}
