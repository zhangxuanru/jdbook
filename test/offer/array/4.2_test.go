/*
@Time : 2020/4/3 14:53
@Author : zxr
@File : 4.2_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何查找数组中元素的最大值和最小值
func TestGetArrMaxAndMin(t *testing.T) {
	nums := []int{8, 5, 9, 90, 23, 45, 100, 12, -10, 87, 6, 2}
	max, min := array.GetArrMaxAndMin(nums, 0, len(nums)-1)
	fmt.Println("分治法获取数组中的最大值和最小值:", max, "----", min)
	max, min = array.GetArrMaxAndMin2(nums)
	fmt.Println("顺序法获取数组中的最大值和最小值:", max, "----", min)
}
