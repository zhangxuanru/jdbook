/*
@Time : 2020/4/3 16:03
@Author : zxr
@File : rotate_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	arr := array.RotateArr(nums, 3)
	fmt.Println(arr)

	nums = []int{1, 2, 3, 4, 5}
	array.RotateArr2(nums, 2)
	fmt.Println("RotateArr2:", nums)
}
