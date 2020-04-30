/*
@Time : 2020/3/31 11:42
@Author : zxr
@File : 2.10_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何从数组中找出a+b=c+d的数对
func TestFindPairs(t *testing.T) {
	arr := []int{3, 4, 7, 10, 20, 9, 8}
	pairs := array.FindPairs(arr)
	fmt.Println(pairs)
}
