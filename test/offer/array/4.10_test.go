/*
@Time : 2020/4/7 16:39
@Author : zxr
@File : 4.10_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何求数组连续最大和
func TestFindContMaximumSum(t *testing.T) {
	nums := []int{1, -2, 4, 8, -4, 7, -1, -5}
	sum := array.FindMaxSubArr(nums)
	fmt.Println("数组连续最大和:", sum)

	sum = array.FindMaxSubArr2(nums)
	fmt.Println("数组连续最大和:", sum)

	sum = array.FindMaxSubArr3(nums)
	fmt.Println("数组连续最大和:", sum)
}
