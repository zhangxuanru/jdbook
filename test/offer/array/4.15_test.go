/*
@Time : 2020/4/8 15:21
@Author : zxr
@File : 4.15_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何对数组进行循环移位
func TestRightShift(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 2
	//array.RightShift(arr, k)
	//array.RightShift2(arr, k)
	array.RightShift3(arr, k)
	fmt.Println(arr)
}
