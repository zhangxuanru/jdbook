/*
@Time : 2020/4/28 14:36
@Author : zxr
@File : 8.5_test
@Software: GoLand
*/
package sorting

import (
	"fmt"
	"jdbook/offer/sorting"
	"testing"
)

//快速排序
func TestQS(t *testing.T) {
	arr := []int{30, 4, 23, 90, 57, 45, 8, 9, 67, 3}
	sorting.QS1(arr, 0, len(arr)-1)
	fmt.Println("快速排序:", arr)
}
