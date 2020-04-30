/*
@Time : 2020/4/28 16:39
@Author : zxr
@File : sortList_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//排序汇总
func TestSelectSort2(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.SelectSort2(arr)
	fmt.Println("TestSelectSort2:", arr)
}

//排序汇总
func TestInsertSort2(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.InsertSort2(arr)
	fmt.Println("TestInsertSort2:", arr)
}

func TestBubSort2(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.BubSort2(arr)
	fmt.Println("BubSort2:", arr)
}

func TestQuickSort2(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.QuickSort2(arr)
	fmt.Println("QuickSort2:", arr)
}

func TestQuickSort3(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.QuickSort3(arr, 0, len(arr)-1)
	fmt.Println("QuickSort3:", arr)
}
