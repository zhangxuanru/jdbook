/*
@Time : 2020/4/7 16:17
@Author : zxr
@File : 4.9
@Software: GoLand
*/
package array

import "math"

//如何求数组中绝对值最小的数
func FindMin(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	end := len(arr) - 1
	if arr[0] > 0 {
		return arr[0]
	}
	if arr[end] <= 0 {
		return arr[end]
	}
	minNum := math.MaxInt64
	for _, v := range arr {
		if math.Abs(float64(v)) < math.Abs(float64(minNum)) {
			minNum = v
		}
	}
	return minNum
}
