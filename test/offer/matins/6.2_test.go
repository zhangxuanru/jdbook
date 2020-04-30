/*
@Time : 2020/4/24 10:59
@Author : zxr
@File : 6.2_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"testing"
)

//如何判断一个数是否是2的N次方
func TestIsPow2N(t *testing.T) {
	n := matins.IsPow2N(8)
	fmt.Println("8 是否是2的N次方:", n)
	n2 := matins.IsPow2N2(16)
	fmt.Println("16 是否是2的N次方:", n2)
}
