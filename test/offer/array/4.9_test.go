/*
@Time : 2020/4/7 16:18
@Author : zxr
@File : 4.9_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何求数组中绝对值最小的数
func TestFindMin(t *testing.T) {
	arr := []int{-10, -5, -2, 7, 15, 50}
	min := array.FindMin(arr)
	fmt.Println("绝对值最小的数为：", min)
}
