/*
@Time : 2020/4/28 11:44
@Author : zxr
@File : 8.3_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//冒泡排序
func TestBubSort(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.BubSort(arr)
	fmt.Println("BubSort:", arr)
}
