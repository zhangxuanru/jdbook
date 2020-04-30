/*
@Time : 2020/4/28 11:06
@Author : zxr
@File : 8.1_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//选择排序
func TestSelectSort(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.SelectSort(arr)
	fmt.Println("SelectSort:", arr)
}
