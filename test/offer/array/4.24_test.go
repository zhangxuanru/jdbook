/*
@Time : 2020/4/9 15:44
@Author : zxr
@File : 4.24_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何对有大量重复的数字的数组进行排序
func TestInOrder(t *testing.T) {
	arr := []int{6, 8, 4, 3, 6, 7, 87, 43, 5, 6, 8, 8, 83, 3, 4, 5, 2}
	array.InOrder(arr)
	fmt.Println(arr)
}
