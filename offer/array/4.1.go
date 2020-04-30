/*
@Time : 2020/4/3 11:30
@Author : zxr
@File : 4.1
@Software: GoLand
*/
package array

import (
	"fmt"
	"math"
)

//如何找出连续数组中唯一重复的元素

//方法一 (hash)
func FindDupByHash(arr []int) int {
	if arr == nil {
		return math.MinInt8
	}
	numHash := make(map[int]int, len(arr))
	for _, v := range arr {
		if _, ok := numHash[v]; ok {
			return v
		} else {
			numHash[v] = 1
		}
	}
	return math.MinInt8
}

//方法二（累加求和法）
func FindDupBySum(arr []int) int {
	sum := 0
	sum1 := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
		sum1 += i
	}
	return sum - sum1
}

//方法三（异或法）
func FindDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}
	r := 0
	n := len(arr)
	for _, v := range arr {
		r = r ^ v
	}
	for i := 1; i < n; i++ {
		r = r ^ i
	}
	return r
}

//方法四（数据映射法）
func FindDupByMap(arr []int) int {
	if arr == nil {
		return -1
	}
	n := len(arr)
	i := 0
	index := 0
	for true {
		if arr[i] >= n {
			return -1
		}
		if arr[index] < 0 {
			break
		}
		arr[index] *= -1
		index = arr[index] * -1
		if index > n {
			fmt.Println("error:")
			return -1
		}
	}
	return index
}
