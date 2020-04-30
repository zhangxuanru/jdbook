/*
@Time : 2020/4/7 18:29
@Author : zxr
@File : 4.11
@Software: GoLand
*/
package array

import (
	"fmt"
)

//如何找出数组中出现一次的数[一个数组里，除了3个数是出现一次的，其它都是出现偶数次，找出三个数中任意一个]
func FindOneNum(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	intMap := make(map[int]int)
	n := len(arr)
	mid := n / 2
	end := n - 1
	for i := 0; i <= mid; i++ {
		intMap[arr[i]]++
		if end != mid {
			intMap[arr[end]]++
			end--
		}
	}
	fmt.Printf("%#v\n\n", intMap)
	for k, v := range intMap {
		if v == 1 {
			return k
		}
	}
	return -1
}
