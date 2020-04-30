/*
@Time : 2020/4/28 11:24
@Author : zxr
@File : 8.2_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//插入排序
func TestInsertSort(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.InsertSort(arr)
	fmt.Println("InsertSort:", arr)
}
