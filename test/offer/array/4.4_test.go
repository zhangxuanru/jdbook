/*
@Time : 2020/4/3 16:52
@Author : zxr
@File : 4.4_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestFindLoseNumByArr(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 9}
	n := array.FindLoseNumByArr(nums)
	fmt.Println(n)
}
