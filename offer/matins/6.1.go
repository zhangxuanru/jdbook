/*
@Time : 2020/4/23 18:38
@Author : zxr
@File : 6.1
@Software: GoLand
*/
package matins

import "fmt"

//判断一个自然数是否是某个数的平方
func IsPow(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 1; i < n/2; i++ {
		m := i * i
		if m == n {
			return true
		}
		if m > n {
			return false
		}
	}
	return false
}

//判断一个自然数是否是某个数的平方
func IsPow2(n int) bool {
	if n <= 0 {
		return false
	}
	low := 0
	high := n
	for low <= high {
		mid := (low + high) / 2
		m := mid * mid
		fmt.Println("low=", low, "high=", high, "mid=", mid, "m=", m)
		if m > n {
			high = mid - 1
		} else if m < n {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
