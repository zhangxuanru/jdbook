/*
@Time : 2020/4/3 17:38
@Author : zxr
@File : 4.6_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestFindSmallKNumber(t *testing.T) {
	nums := []int{100, 9, 3, 6, 2, 1, 7, 90, 34, 23, 8, 95}
	k := 4
	number := array.FindSmallKNumber(nums, 4)
	fmt.Println("第", k, "小的数是：", number)
}
