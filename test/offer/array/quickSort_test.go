/*
@Time : 2020/4/1 11:54
@Author : zxr
@File : quickSort_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{7, 43, 6, 3, 4, 12, 9, 1, 5, 23, 11, 2}
	array.QuickSort(nums)
	fmt.Println(nums)
}
