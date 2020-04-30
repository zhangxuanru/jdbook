/*
@Time : 2020/4/23 18:38
@Author : zxr
@File : 6.1_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"testing"
)

//判断一个自然数是否是某个数的平方
func TestIsPow(t *testing.T) {
	pow := matins.IsPow(16)
	fmt.Println("isPow:", pow)

	pow2 := matins.IsPow2(16)
	fmt.Println("IsPow2:", pow2)
}
