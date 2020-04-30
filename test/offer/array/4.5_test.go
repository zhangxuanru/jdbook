/*
@Time : 2020/4/3 17:29
@Author : zxr
@File : 4.5_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestFindOddNums(t *testing.T) {
	nums := []int{3, 5, 6, 6, 5, 7, 2, 2}
	oddNums := array.FindOddNums(nums)
	fmt.Println(oddNums)
}
