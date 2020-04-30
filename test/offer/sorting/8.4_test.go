/*
@Time : 2020/4/28 11:53
@Author : zxr
@File : 8.4_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//归并排序
func TestMcrge(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.McrgeSort(arr, 0, len(arr)-1)
	fmt.Println("归并排序：", arr)
}
