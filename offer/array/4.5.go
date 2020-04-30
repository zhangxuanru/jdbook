/*
@Time : 2020/4/3 17:08
@Author : zxr
@File : 4.5
@Software: GoLand
*/
package array

//如何找出数组中出现奇数次的数

/**
数组中有N+2个数，其中，N个数出现了偶数次，2个数出现了奇数次（这2个数不相等，找出这2个数）
**/

//方法一（HASH方法）
func FindOddNums(nums []int) []int {
	numMap := make(map[int]int)
	var ret []int
	var n int
	for _, v := range nums {
		if i, ok := numMap[v]; ok {
			if i == 0 {
				n = 1
			} else {
				n = 0
			}
			numMap[v] = n
		} else {
			numMap[v] = 1
		}
	}
	for v, n := range numMap {
		if n == 1 {
			ret = append(ret, v)
		}
	}
	return ret
}
