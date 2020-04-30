/*
@Time : 2020/4/9 15:43
@Author : zxr
@File : 4.24
@Software: GoLand
*/
package array

import (
	"sort"
)

//如何对有大量重复的数字的数组进行排序

//哈希法
func InOrder(arr []int) {
	dataCount := make(map[int]int)
	serice := make([]int, 0)
	for _, v := range arr {
		dataCount[v]++
		if dataCount[v] == 1 {
			serice = append(serice, v)
		}
	}
	sort.Ints(serice)
	index := 0
	for _, v := range serice {
		val := dataCount[v]
		for i := 0; i < val; i++ {
			arr[index] = v
			index++
		}
	}
}

func quickSortTmp(arr []int) {
	start := 0
	end := len(arr) - 1
	mid := 0
	for start < end {
		if arr[mid] < arr[end] {
			end--
			continue
		}
		if arr[mid] >= arr[start] {
			start++
			continue
		}
		arr[start], arr[end] = arr[end], arr[start]
	}
	if arr[end] < arr[mid] {
		arr[end], arr[mid] = arr[mid], arr[end]
	}
	if len(arr[:end]) > 1 {
		quickSortTmp(arr[:end])
	}
	if len(arr[end+1:]) > 1 {
		quickSortTmp(arr[end+1:])
	}
}
