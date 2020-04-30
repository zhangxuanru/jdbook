/*
@Time : 2020/4/7 16:32
@Author : zxr
@File : 4.10
@Software: GoLand
*/
package array

import (
	"fmt"
	"math"
)

//如何求数组连续最大和

//方法一
func FindMaxSubArr(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	max := 0
	sum := 0
	start := -1
	end := -1
	for i, v := range arr {
		sum += v
		if v > sum {
			sum = v
			start = i
		}
		if sum > max {
			end = i
			max = sum
		}
	}
	fmt.Println("最大和对应的数组起始和结束坐标为：", start, ",", end)
	return max
}

//方法二
func FindMaxSubArr2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	n := len(arr)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

//方法三
func FindMaxSubArr3(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	nAll := arr[0]
	nEnd := arr[0]
	for _, v := range arr {
		nEnd = int(math.Max(float64(nEnd+v), float64(v)))
		nAll = int(math.Max(float64(nAll), float64(nEnd)))
	}
	return nAll
}
