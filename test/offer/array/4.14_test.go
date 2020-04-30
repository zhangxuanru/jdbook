/*
@Time : 2020/4/8 14:59
@Author : zxr
@File : 4.14_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

//如何求集合的所有子集
func TestGetAllSybsetSub(t *testing.T) {
	str := "abc"
	arr := array.GetAllSybsetSub(str)
	for _, v := range arr {
		fmt.Println(v)
	}
}
