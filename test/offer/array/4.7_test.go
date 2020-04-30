/*
@Time : 2020/4/7 12:03
@Author : zxr
@File : 4.7_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何求数组中两个元素的最小距离
func TestMinDistance(t *testing.T) {
	nums := []int{4, 5, 6, 4, 7, 4, 6, 4, 7, 8, 5, 6, 4, 3, 10, 8}
	num1 := 4
	num2 := 8
	distance := array.MinDistance(nums, num1, num2)
	fmt.Println("最小距离为:", distance)

	distance = array.MinDistance2(nums, num1, num2)
	fmt.Println("最小距离为:", distance)
}
