/*
@Time : 2020/4/8 16:41
@Author : zxr
@File : 4.16_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何在有规律的二维数组中进行高效的数据查找
func TestFindWithBinary(t *testing.T) {
	arr := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	binary := array.FindWithBinary(arr, 11)
	fmt.Println("binary:", binary)
}
