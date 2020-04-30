/*
@Time : 2020/4/24 15:24
@Author : zxr
@File : 6.7
@Software: GoLand
*/
package matins

import "math"

//如何比较a,b两个数的大小， 不能用大于，小于 语句

//方法一（绝对值法）
func Max1(a, b int) int {
	if math.Abs(float64(a-b)) == float64(a-b) {
		return a
	} else {
		return b
	}
}
