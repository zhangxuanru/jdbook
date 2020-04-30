/*
@Time : 2020/4/8 18:39
@Author : zxr
@File : 4.22_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何从三个有序数组中找出它们的公共元素
func TestFindCommon(t *testing.T) {
	arr1 := []int{2, 5, 12, 20, 45, 85}
	arr2 := []int{16, 19, 20, 85, 200}
	arr3 := []int{3, 4, 15, 20, 39, 72, 85, 190}
	common := array.FindCommon(arr1, arr2, arr3)
	fmt.Println("common:", common)
}
