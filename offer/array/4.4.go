/*
@Time : 2020/4/3 16:44
@Author : zxr
@File : 4.4
@Software: GoLand
*/
package array

//找出数组中丢失的数
/**
给定一个由n-1个整数组成的未排序的数组序列，
其元素都是1----n中的不同的整数。请写出一个寻找数组序列中缺失整数的算法
*/

//方法一(求和法)
func FindLoseNumByArr(nums []int) int {
	n := len(nums)
	sum := 0
	sum1 := n
	for i := 0; i < n; i++ {
		sum += nums[i]
		sum1 += i
	}
	sum1 += n + 1
	return sum1 - sum
}
