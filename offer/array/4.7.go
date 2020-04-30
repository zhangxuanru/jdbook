/*
@Time : 2020/4/7 12:02
@Author : zxr
@File : 4.7
@Software: GoLand
*/
package array

import "math"

//如何求数组中两个元素的最小距离

//方法一
func MinDistance(nums []int, num1 int, num2 int) int {
	minDis := math.MaxInt64
	if nums == nil || len(nums) == 0 {
		return minDis
	}
	dist := 0
	for i, v := range nums {
		if v == num1 {
			for j, v := range nums {
				if v == num2 {
					dist = int(math.Abs(float64(j - i)))
					if dist < minDis {
						minDis = dist
					}
				}
			}
		}
	}
	return minDis
}

//方法二
func MinDistance2(arr []int, num1, num2 int) int {
	minDis := math.MaxInt64
	if arr == nil || len(arr) == 0 {
		return minDis
	}
	last1 := -1
	last2 := -1
	for i, v := range arr {
		if v == num1 {
			last1 = i
			if last2 >= 0 {
				minDis = int(math.Min(float64(minDis), float64(last1-last2)))
			}
		}
		if v == num2 {
			last2 = i
			if last1 > 0 {
				minDis = int(math.Min(float64(minDis), float64(last2-last1)))
			}
		}
	}
	return minDis
}
